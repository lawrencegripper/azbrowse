# AzBrowse

An interactive CLI for browsing azure resources, inspired by [resources.azure.com](https://resources.azure.com)

[![Build Status](https://travis-ci.com/lawrencegripper/azbrowse.svg?branch=master)](https://travis-ci.com/lawrencegripper/azbrowse)

# Quick Start

Simply download the binary suitable for your machine, [from the release page](https://github.com/lawrencegripper/azbrowse/releases), and execute it.

### Status

It's an MVP to prove out the use case. Basic navigation and operations with a boltdb based cache for expensive (slow) API calls.

Currently I'm using it every day **but it is experimental so use with caution on a production environment!!**

![Demo](./docs/quickdemo-azbrowse.gif)

### Install

Grab the binaries from the release page or for MacOS and Linux run this script

```
curl -sSL https://raw.githubusercontent.com/lawrencegripper/azbrowse/master/scripts/install_azbrowse.sh | sudo sh
```

You may need to reload your terminal to pick up `azbrowse` after the script completes.

### Usage

Below is a table containing the default key bindings. If you'd like to customise the key bindings to be more suitable for your setup, please refer to the section on [custom key bindings](#custom-key-bindings).

## Navigation

| Key       | Does                 |
| --------- | -------------------- |
| ↑/↓       | Select resource      |
| Backspace | Go back              |
| ENTER     | Expand/View resource |

## Operations

| Key                 | Does                      |                                                                                    |
| ------------------- | ------------------------- | ---------------------------------------------------------------------------------- |
| CTRL+E              | Toggle Browse JSON        | For longer responses you can move the cursor to scroll the doc                     |
| CTRL+o (o for open) | Open Portal               | Opens the portal at the currently selected resource                                |
| DEL:                | Delete resource           | The currently selected resource will be deleted (Requires double press to confirm) |
| CTLT+F:             | Toggle Fullscreen         | Gives a fullscreen view of the JSON for smaller terminals                          |
| CTLT+S:             | Save JSON to clipboard    | Saves the last JSON response to the clipboard for export                           |
| CTLT+A:             | View Actions for resource | This allows things like ListKeys on storage or Restart on VMs                      |




## Debugging

Running `azbrowse --debug` will start an in-memory collector for the `opentracing` and a GUI to browse this at http://localhost:8700. You can use this to look at tracing information output by `azbrowse` as it runs.

![tracing ui](docs/trace.png)

## Developing

### Environment Setup

> Note: Golang 1.12 is recommended. 

First, clone this repository. `azbrowse` is written in [Go][golang] and so you will want to set up your Go development environment first. If this is your first time, the [offical install guide][installguide] is probably a good place to start. Make sure you have `GOPATH/bin` in your `PATH`, using the instructions [here][gopath] as guidance on doing that.

In addition to installing [Go][golang], there are a couple of tool dependencies you'll need. These are:

- [Go Meta Linter][gometalinter]
- [Dep; Go dependency management tool][golang]

You can install these yourself following the instructions on their github pages, or you can run...

 ``` bash
 make setup
 ```

 This runs the script `scripts/install_dev_tools.sh`, which will install these tools for you.

### Building

With your Go development environment set up, use `make` to build `azbrowse`.

Take a look at the `Makefile` yourself, but the main rules are:

#### Run Tests and Build

``` bash
make build
```

#### Install Local Development Build


``` bash
make install
```

#### Run Travis-CI build locally

``` bash
make ci-docker
```

To run the full Travis-CI locally, you need to have the `TRAVIS_BUILD_NUMBER` environment variable defined, so running it as follows may be easier:

```bash
TRAVIS_BUILD_NUMBER=0.1 make ci-docker
```

## Custom Key Bindings

If you wish to override the default key bindings, create a `bindings.json` file in the same directory as your `azbrowse` binary.

The file should be formated like so:
```json
{
    ...
    "Copy": "F8",
    "Help": "Ctrl+H",
    ...
}
```

In the file you can override the keys for actions using keys from the lists below.

### Actions

| Actions:       | Does                                  |
| ------------------------ | --------------------------------------|
| Quit                     | Terminates the program                |
| Copy                     | Copies the resource JSON to clipboard |
| ListDelete               | Deletes a resources                   |
| Fullscreen               | Toggles fullscreen                    |
| Help                     | Toggles help view                     |
| ItemBack                 | Go back from an item to a list        |
| ItemLeft                 | Switch from the item json to the menu |
| ListActions              | List available actions on a resource  |
| ListBack                 | Go back on a list                     |
| ListBackLegacy           | Go back on a list (legacy terminals)  |
| ListDown                 | Navigate down a list                  |
| ListUp                   | Navigate up a list                    |
| ListRight                | Switch from the list to an item view  |
| ListEdit                 | Toggle edit mode on a resource        |
| ListExpand               | Expand a selected resource            |
| ListOpen                 | Open a resource in the Azure portal   |
| ListRefresh              | Refresh a list                        |

### Keys

- Up
- Down
- Left
- Right
- Backspace
- Backspace2
- Delete
- Home
- End
- PageUp
- PageDown
- Insert
- Tab
- Space
- Esc
- Enter
- Ctrl+2
- Ctrl+3
- Ctrl+4
- Ctrl+5
- Ctrl+6
- Ctrl+7
- Ctrl+8
- Ctrl+[
- Ctrl+]
- Ctrl+Space
- Ctrl+_
- Ctrl+~
- Ctrl+A
- Ctrl+B
- Ctrl+C
- Ctrl+D
- Ctrl+E
- Ctrl+F
- Ctrl+G
- Ctrl+H
- Ctrl+I
- Ctrl+J
- Ctrl+K
- Ctrl+L
- Ctrl+M
- Ctrl+N
- Ctrl+O
- Ctrl+P
- Ctrl+Q
- Ctrl+R
- Ctrl+S
- Ctrl+T
- Ctrl+U
- Ctrl+V
- Ctrl+W
- Ctrl+X
- Ctrl+Y
- Ctrl+Z
- F1
- F2
- F3
- F4
- F5
- F6
- F7
- F8
- F9
- F10
- F11
- F12

> For compatibility reasons you may notice some keys will have multiple mappings.

## Plans

[Issues on the repository track plans](https://github.com/lawrencegripper/azbrowse/issues), I'd love help so feel free to comment on an issue you'd like to work on and we'll go from there.

[golang]: https://golang.org/
[installguide]: https://golang.org/doc/install
[gometalinter]: https://github.com/alecthomas/gometalinter
[golangdep]: https://github.com/golang/dep
[gopath]: https://golang.org/doc/code.html#GOPATH
