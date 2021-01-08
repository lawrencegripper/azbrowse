package tfprovider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-logr/logr"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/hashicorp/terraform-exec/tfinstall"
	"github.com/hashicorp/terraform/lang"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/providers"

	"github.com/zclconf/go-cty/cty"
)

const (
	providerNameEnv   = "TF_PROVIDER_NAME"
	providerVerionEnv = "TF_PROVIDER_VERSION"
	providerPathEnv   = "TF_PROVIDER_PATH"
	tfVersion         = "0.12.29"
)

// TerraformProvider represents a terraform provider install
type TerraformProvider struct {
	Plugin   *plugin.GRPCProvider
	Metadata *TerraformProviderMetadata
}

// TerraformProviderMetadata defines metadata about the terraform provider
type TerraformProviderMetadata struct {
	Name           string
	Version        string
	ChecksumSHA256 string
}

// SetupProvider will return an instance of the TF provider.
// If the TF_PROVIDER_PATH env var is set and a Terraform provider exists at
// that location, this provider will be used. This is the recommended approach of production.
// If the TF_PROVIDER_PATH env var is not set, we will download and initialize the
// provider using Hashicorp's Terraform Registry. This relies on internet access.
func SetupProvider(log logr.Logger) (*TerraformProvider, error) {
	providerName := os.Getenv(providerNameEnv)
	if providerName == "" {
		return nil, fmt.Errorf("Env %q not set and is required", providerNameEnv)
	}

	var err error
	var providerVersion string
	providerPath := os.Getenv(providerPathEnv)
	if providerPath == "" {
		log.Info("Downloading provider binary")
		providerVersion = os.Getenv(providerVerionEnv)
		if providerVersion == "" {
			return nil, fmt.Errorf("Env %q not set and is required when path to provider binary isn't set with %q", providerVerionEnv, providerPathEnv)
		}
		providerPath, err = installProvider(providerName, providerVersion)
		if err != nil {
			return nil, fmt.Errorf("Failed to setup provider as provider install failed: %w", err)
		}
	}
	providerInstance, err := getInstanceOfProvider(providerName, providerPath, providerVersion)
	if err != nil {
		return nil, fmt.Errorf("failed getting provider instance %w", err)
	}
	err = configureProvider(log, providerInstance.Plugin)
	if err != nil {
		return nil, err
	}

	return providerInstance, nil
}

func installProvider(name string, version string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "tfinstall")
	if err != nil {
		return "", fmt.Errorf("Failed to create temp dir. %w", err)
	}
	defer os.RemoveAll(tmpDir) //nolint: errcheck

	execPath, err := tfinstall.Find(tfinstall.ExactVersion(tfVersion, tmpDir))
	if err != nil {
		return "", fmt.Errorf("Failed to install Terraform %w", err)
	}

	workingDir, err := ioutil.TempDir("", "tfproviders")
	if err != nil {
		return "", fmt.Errorf("Failed create tfprovider dir %w", err)
	}

	providerFileContent := fmt.Sprintf(`
	provider "%s" {
		version = "%s"
	}
	`, name, version)

	err = ioutil.WriteFile(path.Join(workingDir, "provider.tf"), []byte(providerFileContent), 0644)
	if err != nil {
		return "", fmt.Errorf("Failed to create provider.tf file %w", err)
	}
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return "", fmt.Errorf("Failed to create TF context %w", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true), tfexec.LockTimeout("60s"))
	if err != nil {
		return "", fmt.Errorf("Failed to init TF %w", err)
	}

	return path.Join(workingDir, "/.terraform/plugins/linux_amd64/"), nil
}

func getSHA256Checksum(path string) (string, error) {
	hash := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close() //nolint: errcheck
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	checksum := hex.EncodeToString(hash.Sum(nil))
	return checksum, nil
}

func getInstanceOfProvider(name, path, version string) (*TerraformProvider, error) {
	pluginMeta := discovery.FindPlugins(plugin.ProviderPluginName, []string{path}).WithName(name)

	if pluginMeta.Count() < 1 {
		return nil, fmt.Errorf("Provide:%q not found at path:%q", name, path)
	}
	versionedPlugin := pluginMeta.Newest()
	clientConfig := plugin.ClientConfig(versionedPlugin)

	// Don't log provider details unless provider log is enabled by env
	if _, exists := os.LookupEnv("ENABLE_PROVIDER_LOG"); !exists {
		clientConfig.Logger = hclog.NewNullLogger()
	}
	pluginClient := goplugin.NewClient(clientConfig)

	rpcClient, err := pluginClient.Client()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize plugin: %w", err)
	}

	// create a new resource provisioner.
	raw, err := rpcClient.Dispense(plugin.ProviderPluginName)
	if err != nil {
		panic(fmt.Errorf("Failed to dispense plugin: %s", err))
	}

	ver, err := versionedPlugin.Version.Parse()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse provider version %w", err)
	}
	checksumSHA256, err := getSHA256Checksum(versionedPlugin.Path)
	if err != nil {
		return nil, err
	}
	return &TerraformProvider{
		Plugin: raw.(*plugin.GRPCProvider),
		Metadata: &TerraformProviderMetadata{
			Name:           versionedPlugin.Name,
			Version:        ver.String(),
			ChecksumSHA256: checksumSHA256,
		},
	}, nil
}

func createEmptyProviderConfWithDefaults(provider *plugin.GRPCProvider, configBody string) (*cty.Value, error) {
	if configBody == "" {
		configBody = os.Getenv("PROVIDER_CONFIG_HCL")
	}

	providerConfigBlock := provider.GetSchema().Provider.Block

	// Parse the content of the provider block given to us into a body.
	// Note: The file name is required but isn't important in this context so we provide a nonexistent dummy filename.
	file, diagParse := hclsyntax.ParseConfig([]byte(configBody), "dummy.tf", hcl.Pos{})
	if diagParse.HasErrors() {
		return nil, fmt.Errorf("Failed parsing provider config block: %s", diagParse.Error())
	}

	scope := lang.Scope{}
	expandedConf, diags := scope.ExpandBlock(file.Body, providerConfigBlock)
	if diags.Err() != nil {
		return nil, fmt.Errorf("Failed expanding provider config block: %w", diags.Err())
	}
	configFull, diags := scope.EvalBlock(expandedConf, providerConfigBlock)
	if diags.Err() != nil {
		return nil, fmt.Errorf("Failed evaluating provider config block: %w", diags.Err())
	}

	// Call the `PrepareProviderConfig` with the config object. This returns a version of that config with the
	// required default setup as `PreparedConfig` under the response object.
	// Warning: Diagnostics houses errors, the typical go err pattern isn't followed - must check `resp.Diagnostics.Err()`
	prepConfigResp := provider.PrepareProviderConfig(providers.PrepareProviderConfigRequest{
		Config: configFull,
	})
	if err := prepConfigResp.Diagnostics.Err(); err != nil {
		return nil, fmt.Errorf(`Failed to set configure provider from config: %w.`+
			`Hint: See startup docs on using "PROVIDER_CONFIG_HCL" or the providers env vars to set required fields`, err)
	}

	return &configFull, nil
}

func configureProvider(log logr.Logger, provider *plugin.GRPCProvider) error {
	configWithDefaults, err := createEmptyProviderConfWithDefaults(provider, "")
	if err != nil {
		return err
	}
	// Now we have a prepared config we can configure the provider.
	// Warning (again): Diagnostics houses errors, the typical go err pattern isn't followed - must check `resp.Diagnostics.Err()`
	configureProviderResp := provider.Configure(providers.ConfigureRequest{
		Config: *configWithDefaults,
	})
	if err := configureProviderResp.Diagnostics.Err(); err != nil {
		log.Error(err, fmt.Sprintf("Failed to configure provider: %s", err))
		return err
	}

	return nil
}
