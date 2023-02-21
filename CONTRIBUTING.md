# Contributing

## Design

[Before you get started we strongly recomment reading this doc on](docs/design/README.md) the design of `azbrowse`.

## Developing

### Environment Setup

The repository has a pre-configured `devcontainer` which can be used with VSCode to setup and run a fully configured build environment. 

1. Install [Remote Extensions Pack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
2. Clone the Repo and open in VSCode 
4. `CTRL`+`Shift`+`P` -> `Reopen in container` (or `cmd` on mac / however you get to your command panel in VSCode)

The first time you launch with the container a build will occur which can take some time, subsequent launches will use the cache so start immediately. 

### Without VSCode

Using [`devcontainer cli`](https://github.com/devcontainers/cli/) you can get the full development environment without VSCode.

```bash
	Run 'npm install -g @devcontainers/cli' to install the CLI
	devcontainer up --workspace-folder ${PWD}
	devcontainer run-user-commands --workspace-folder ${PWD}
	devcontainer exec --workspace-folder ${PWD} '/workspaces/azbrowse/scripts/local-ci.sh'
```

## Running full CI locally

### Inside VSCode

1. Start the repo and build the devcontainer using [Environment Setup](#environment-setup)
1. Open a terminal
1. Run `make local-ci`

### Without VSCode

Prereqs:
- Docker
- [`devcontainer cli`](https://github.com/devcontainers/cli/)


```
make devcontainer-local-ci
```

This will spin up a docker container and execute the CI within it.

## Building

With your Go development environment set up, use `make` to build `azbrowse`.

### What Targets are there?

Run `make help` to see what `targets` exist and what they do or take a look at the `Makefile` yourself. 

Below are the simples ones to get you started.

### Run Checks, Tests and Build

```bash
# Runs linting and formatting checks
make checks
```

``` bash
make build
```

Running integration tests (requires a full terminal)

``` bash
make integration
```

### Install Local Development Build

``` bash
make install
```

## Debugging

Running `azbrowse --debug` will start an in-memory collector for the `opentracing` and a GUI to browse this at http://localhost:8700. You can use this to look at tracing information output by `azbrowse` as it runs.

![tracing ui](docs/images/trace.png)

## Automated builds

The `AzureDevOps` build runs `golang` build, unit tests and linting then runs integration tests under `XTerm`.

To check `golang` build, unit tests and linting pass run `make build` before pushing.

### Running integration tests/release locally

The CI process re-uses the `devcontainer` used for developing the solution locally so you can reproduce build errors locally.

Integration tests in docker:

``` bash
make devcontainer-integration
```

To run the full CI locally, you need to have the `BUILD_NUMBER` environment variable defined, so running it as follows may be easier:

```bash
export BUILD_NUMBER=99999
make devcontainer-release
```

If you would like to do a publish step of the assets set `IS_CI=1`, `BUILD_NUMBER`, `DOCKER_USERNAME`, `DOCKER_PASSWORD`, `BRANCH=main` and `GITHUB_TOKEN` then run `make devcontainer-release`

## Patterns

The aim of this section of the doc is to capture some of the patterns for working with `azbrowse` to add new providers or other features.

If using [Visual Studio Code](https://code.visualstudio.com) and you have the Remote Development extension (+Docker) installed then you can take advantage of the devcontainer support to have a development environment set up ready for you to use.

## Testing

### Mocked Tests 

Each `expander` implements a set of `testcases` which mock out the `armclient` responses and assert that the `expander` behaves correct. 

For a good example see the [`default` expanders](https://github.com/lawrencegripper/azbrowse/blob/8b307d3/internal/pkg/expanders/default.go#L84) test set. 

Below is an annotated version of the code for example. 

```golang
    const testPath = "subscriptions/thing"

    // Build an example item which your `expander` should handle
	itemToExpand := &TreeNode{
		ExpandURL: "https://management.azure.com/" + testPath,
    }
    
    // Provide the body of the response you want the ARM client to give when 
    // it's called. 
	const testResponseFile = "./testdata/armsamples/resource/failingResource.json"

	return true, &[]expanderTestCase{
		{
            // This case asserts that the `expander` returns correctly with the `statusIndicator` set correctly
			name:         "Default->Resource",
			nodeToExpand: itemToExpand,
			urlPath:      testPath,
			responseFile: testResponseFile,
            statusCode:   200, // Define Status code the ARM client mock should return when called
            // This function is used to assert the expander result is what you expected \/ 
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				st.Expect(t, r.Err, nil)
				st.Expect(t, len(r.Nodes), 0)

				dat, err := os.ReadFile(testResponseFile)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				st.Expect(t, strings.TrimSpace(r.Response.Response), string(dat))
				st.Expect(t, itemToExpand.StatusIndicator, "⛈")
			},
		},
		{
            // This test case asserts that the `expander` returns an error when the ARM client Mock returns a 500 status code. 
			name:         "Default->500StatusCode",
			nodeToExpand: itemToExpand,
			urlPath:      testPath,
			responseFile: testResponseFile,
			statusCode:   500,
			treeNodeCheckerFunc: func(t *testing.T, r ExpanderResult) {
				if r.Err == nil {
					t.Error("Failed expanding resource. Should have errored and didn't", r)
				}
			},
		},
	}
```

### Integration testing/Fuzzing

> WARNING: The fuzzer is EXPERIMENTAL and will talk all subscriptions you have access to and call nodes. Please ensure you use with caution so as not to effect production environments.

In many circumstances the ARM endpoints may return results which can't be predicted from the documentation. 

The only way to test these cases is against a live subscription. `azbrowse` features a [`fuzzer`](https://github.com/lawrencegripper/azbrowse/blob/8b307d3/internal/pkg/automation/fuzzer.go#L30) which will do an automated walk of a subscription. 

You can use the [`shouldSkip` func](https://github.com/lawrencegripper/azbrowse/blob/8b307d3/internal/pkg/automation/fuzzer.go#L51) to limit the number of child nodes walked under a resource or skip a resource. This is useful on items which have unbounded numbers of children like `Activity Log` or `Deployments`. 

You should add assertions about how your items will be handled to [`testFunc`](https://github.com/lawrencegripper/azbrowse/blob/8b307d3/internal/pkg/automation/fuzzer.go#L44) so during the fuzzing checks are made on the results. 

To launch the fuzzer use `azbrowse -fuzzer 1` where `1` is the number of mins to walk the node tree from. 

When working on a particular area you can use the `fuzzer` with the `-navigate` command and it will jump to a node specified then start walking the node tree.

The `make` file provides shortcuts `make fuzz` and `make fuzz-from node_id=/subscriptions/SOMESUB/resourceGroups/lk-scratch/providers/Microsoft.Web/sites/SOMESITE` to build and then fuzz easily. 

In future the intention is to have a test subscription and run the fuzzer during PR builds against a known set of resources defined in the subscription. 