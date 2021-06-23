package tfprovider

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

const (
	providerInstallPath = "../../../hack/"
)

type providerTestDef struct {
	name         string
	requiredEnvs map[string]string
	configBody   string
	version      string
}

// Ensure you run ./hack/init.sh to
// install the providers before running these tests
var testedProviders = []providerTestDef{
	{
		name:         "aws",
		requiredEnvs: map[string]string{"AWS_REGION": "us-east-1"},
		version:      "2.70.0",
	},
	{
		name:         "azurerm",
		requiredEnvs: map[string]string{},
		configBody:   `features {}`,
		version:      "2.22.0",
	},
	{
		name:         "helm",
		requiredEnvs: map[string]string{},
		version:      "1.2.3",
	},
}

func Test_PrepareProviderConfigWithDefaults_expectNoError(t *testing.T) {
	for _, tt := range testedProviders {
		t.Run(tt.name, func(t *testing.T) {
			// Clear the HCL config env if already set by previous test
			err := os.Setenv("PROVIDER_CONFIG_HCL", "")
			if err != nil {
				panic(err)
			}

			// Set required envs
			for name, value := range tt.requiredEnvs {
				err = os.Setenv(name, value)
				if err != nil {
					panic(err)
				}
			}
			provider, err := getInstanceOfProvider(context.Background(), tt.name, providerInstallPath, tt.version)
			if err != nil {
				t.Errorf("failed to get instance of provider. error = %v", err)
			}

			_, err = createEmptyProviderConfWithDefaults(provider.Plugin, tt.configBody)
			if err != nil {
				t.Errorf("failed to configure provider with defaults. error = %v", err)
				return
			}
		})
	}
}

func Test_installProvider(t *testing.T) {
	for _, tt := range testedProviders {
		t.Run(tt.name, func(t *testing.T) {

			workingDir, err := ioutil.TempDir("", "tfinstall-test")
			if err != nil {
				t.Errorf("installProvider() error = %v, failed to create temp dir", err)
			}

			err = installProvider(context.Background(), tt.name, tt.version, workingDir)
			if err != nil {
				t.Errorf("installProvider() error = %v, no error expected", err)
			}
			installedToPath := path.Join(workingDir, "/.terraform/plugins/linux_amd64/")

			expectedProviderFilename := fmt.Sprintf("terraform-provider-%s_v%s", tt.name, tt.version)

			files, err := ioutil.ReadDir(installedToPath)
			if err != nil {
				t.Errorf("failed to read dir. error = %v", err)
			}

			foundProvider := false
			for _, file := range files {
				if strings.HasPrefix(file.Name(), expectedProviderFilename) {
					// Found the provider binary in the expected location
					foundProvider = true
				}
			}

			if !foundProvider {
				t.Errorf("missing provider for %q in %q", tt.name, installedToPath)
			}
		})
	}
}

func TestSetupProvider(t *testing.T) {
	tests := []struct {
		purpose  string
		config   TerraformProviderConfig
		wantErr  bool
		skipInCI bool // some tests require az credentials - re-enable once added to CI build
	}{
		{
			purpose: "Error_When_OnlyNameSet",
			config: TerraformProviderConfig{
				ProviderName:    "azurerm",
				ProviderVersion: "2.22.0",
				ProviderPath:    "",
			},
			wantErr: true,
		},
		{
			purpose: "Error_When_PathIsMissingProviderBinary",
			config: TerraformProviderConfig{
				ProviderName:    "azurerm",
				ProviderVersion: "2.22.0",
				ProviderPath:    "/tmp",
			},
			wantErr: true,
		},
		{
			purpose: "Succeed_When_ValidProviderPathAndNameSet",
			config: TerraformProviderConfig{
				ProviderName:      "azurerm",
				ProviderVersion:   "2.22.0",
				ProviderPath:      providerInstallPath,
				ProviderConfigHCL: "features {}",
			},
			wantErr:  false,
			skipInCI: true,
		},
		{
			purpose: "Succeed_When_ValidProviderNameAndVersionSet",
			config: TerraformProviderConfig{
				ProviderName:      "azurerm",
				ProviderVersion:   "2.22.0",
				ProviderPath:      providerInstallPath,
				ProviderConfigHCL: "features {}",
			},
			wantErr:  false,
			skipInCI: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			if tt.skipInCI && os.Getenv("IS_CI") != "" {
				t.Skip("Skipping test in CI")
			}

			got, err := SetupProvider(context.Background(), NewNullLogger(), tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetupProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.wantErr == false {
				// If we're expecting a valid provider try and use it to check it's working
				schemaResult := got.Plugin.GetSchema()
				if len(schemaResult.Diagnostics) > 0 {
					t.Errorf("failed to get schema from provider")
				}
			}
		})
	}
}
