# AzBrowse

An interactive CLI for browsing azure resources, inspired by [resources.azure.com](https://resources.azure.com)

[![Go Report Card](https://goreportcard.com/badge/github.com/lawrencegripper/azbrowse?style=flat-square)](https://goreportcard.com/report/github.com/lawrencegripper/azbrowse)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/lawrencegripper/azbrowse)
[![Release](https://img.shields.io/github/release/lawrencegripper/azbrowse.svg?style=flat-square)](https://github.com/lawrencegripper/azbrowse/releases/latest)
[![release](https://github.com/lawrencegripper/azbrowse/workflows/release/badge.svg)](https://github.com/lawrencegripper/azbrowse/actions?query=workflow%3Arelease+branch%3Amain)

## Status

This is a pet project which has matured thanks to support from awesome contributions.

[![asciicast](https://asciinema.org/a/325237.svg)](https://asciinema.org/a/325237)

> Warning: Please familiarize yourself with the code and the how-to's before using it in a production environment.

## Cool what else can it do?

Lots [check out the guided tour here](docs/getting-started.md).

- Edit/Update resource
- Multi-resource delete
- Actions on resources such as restart and list-keys
- ASCII Graphs for resource metrics
- Interactive command panel for filtering and more
- Build custom views from [Azure Resource Graph Queries](./docs/azure-resource-graph.md)

For advanced [config review the settings page here](docs/config.md). For [command line arguments and docs see this page](./docs/commandline/azbrowse.md).

## Contributing

Take a look at the guide here [for a guide on the design of azbrowse](docs/design/README.md) and then a look here [for how to get started on deving](CONTRIBUTING.md)

## Install/Run

> Pre-req: Ensure you have the [`az` command from Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest) setup on your machine and are logged-in otherwise `azbrowse` won't work!

<details>
  <summary>Mac/OSX via Homebrew</summary>
<br />
    
Install [HomeBrew](https://brew.sh/)

```shell
brew install lawrencegripper/tap/azbrowse
```
</details>
<details>
  <summary>Windows via Scoop</summary>
<br />

[Install Scoop]([Scoop](https://scoop.sh/))

```shell
iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
```

Install AzBrowse using Scoop

```shell
scoop bucket add azbrowse https://github.com/lawrencegripper/scoop-bucket.git
scoop install azbrowse
```
</details>
<details>
    <summary>Run via Docker</summary>
<br />

You can then start `azbrowse` in docker by mounting in your `$HOME` directory so `azbrowse` can access the login details from your machine inside the docker container.

```shell
docker run -it --rm -v $HOME:/root/ -v /etc/localtime:/etc/localtime:ro ghcr.io/lawrencegripper/azbrowse/azbrowse
```
</details>
<details>
    <summary>Linux via Snap Store</summary> 
<br />

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/azbrowse)

</details>

<details>
    <summary>Linux via Releases tar.gz</summary> 
<br />

Grab the URL to the `.tar.gz` for the latest release for your platform/architecture. E.g. `https://github.com/lawrencegripper/azbrowse/releases/download/v1.1.193/azbrowse_linux_amd64.tar.gz`

Download the release (either via the browser or `wget https://github.com/lawrencegripper/azbrowse/releases/download/v1.1.193/azbrowse_linux_amd64.tar.gz`).

Extract the binary from the archive to a suitable location (here we're using `/usr/bin` for convenience): `tar -C /usr/bin -zxvf azbrowse_linux_amd64.tar.gz azbrowse`

> Note: If you have a location on `$PATH` which is writable by the current user like `/home/USERNAMEHERE/go/bin` it's best to use this as it'll allow azbrowse to update itself in place without requiring `sudo` 

Make the binary executable: `chmod +x /usr/bin/azbrowse`

</details>
<details>
    <summary>Install via azure-cli extention</summary>
<br />

This is experimental and Non-functional on Windows. Only tested on Unix based systems.

Want to run `az browse` and have the `azure-cli` install and run `azbrowse`?

[This extension from Noel Bundick lets you do just that](https://github.com/noelbundick/azure-cli-extension-noelbundick/blob/master/README.md#browse)

</details>
<details>
    <summary>DIY</summary>
<br />

Simply download the archive/package suitable for your machine, [from the release page](https://github.com/lawrencegripper/azbrowse/releases), and execute it.

Bonus: Add it to your `$PATH` so you can run `azbrowse` anywhere. 
</details>

## Shell completion

Azbrowse can generate shell completions for a number of different shells using the `azbrowse completion` command. 

For example, `azbrowse -s thing<TAB>` → `azbrowse -s thingSubscription` and jump straight to that Azure subscription.

To configure completion in bash add the following to `~/.bashrc` or `~/.profile`

```bash
source <(azbrowse completion bash)
```

To configure completion for `zsh` run the following command

```bash
azbrowse completion zsh > "${fpath[1]}/_azbrowse"
```

> Ensure you have `autoload -Uz compinit && compinit` present in your `.zshrc` file to load these completions

## Docs

See the [docs](docs/README.md) for getting started guides, configuration docs, ...

## Plans

[Issues on the repository track plans](https://github.com/lawrencegripper/azbrowse/issues), I'd love help so feel free to comment on an issue you'd like to work.

[golang]: https://golang.org/
[installguide]: https://golang.org/doc/install
[golangcilinter]: https://github.com/golangci/golangci-lint
[golangdep]: https://github.com/golang/dep
[gopath]: https://golang.org/doc/code.html#GOPATH

## Updating Snap login for goreleaser

```
sudo snap install snapcraft
snapcraft login
snapcraft export-login .snap.login
cat .snap.login | base64
```

Update the secret for actions on the repo with the base64 output.