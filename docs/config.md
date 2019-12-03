# Configuration

By default azbrowse looks for configuration in `/root/.azbrowse-settings.json` or `~/.azbrowse-settings.json` (where `~` is your users home directory). You can also set the `AZBROWSE_SETTINGS_PATH` environment variable to point to another location where you want to store your config (e.g. `~/clouddrive/.azbrowse-settings.json`)

Below is a table containing some of the default key bindings. If you'd like to customise the key bindings to be more suitable for your setup, please refer to the section on [custom key bindings](#custom-key-bindings).


## Navigation

| Key       | Does                                       |
| --------- | ------------------------------------------ |
| ↑/↓       | Select resource                            |
| PgDn/PgUp | Move up or down a page of resources        |
| Home/End  | Move to the top or bottom of the resources |
| Backspace | Go back                                    |
| ENTER     | Expand/View resource                       |

## Operations

| Key                 | Does                      |                                                                                    |
| ------------------- | ------------------------- | ---------------------------------------------------------------------------------- |
| CTRL+E              | Toggle Browse JSON        | For longer responses you can move the cursor to scroll the doc                     |
| CTRL+o (o for open) | Open Portal               | Opens the portal at the currently selected resource                                |
| DEL:                | Delete resource           | The currently selected resource will be deleted (Requires double press to confirm) |
| CTLT+F:             | Toggle Fullscreen         | Gives a fullscreen view of the JSON for smaller terminals                          |
| CTLT+S:             | Save JSON to clipboard    | Saves the last JSON response to the clipboard for export                           |
| CTLT+A:             | View Actions for resource | This allows things like ListKeys on storage or Restart on VMs                      |

## Custom Key Bindings

If you wish to override the default key bindings, create a `~/.azbrowse-settings.json` file (where `~` is your users home directory).

The file should be formated like so:

```json
{
    "keyBindings": {
        ...
        "Copy": "F8",
        "Help": "Ctrl+H",
        ...
    }
}
```

You can also specify multiple key bindings for a command:

```json
{
    "keyBindings": {
        ...
        "Copy": "F8",
        "Help": "Ctrl+H",
        "ListUp": ["k", "Up"],
        "ListDown": ["j", "Down"]
        ...
    }
}
```

In the file you can override the keys for actions using keys from the lists below.

## Actions

| Actions:                 | Does                                          |
| ------------------------ | ----------------------------------------------|
| Quit                     | Terminates the program                        |
| Copy                     | Copies the resource JSON to clipboard         |
| ListDelete               | Deletes a resources                           |
| Fullscreen               | Toggles fullscreen                            |
| Help                     | Toggles help view                             |
| ItemBack                 | Go back from an item to a list                |
| ItemLeft                 | Switch from the item json to the menu         |
| ListActions              | List available actions on a resource          |
| ListBack                 | Go back on a list                             |
| ListBackLegacy           | Go back on a list (legacy terminals)          |
| ListDown                 | Navigate down a list                          |
| ListUp                   | Navigate up a list                            |
| ListRight                | Switch from the list to an item view          |
| ListEdit                 | Toggle edit mode on a resource                |
| ListExpand               | Expand a selected resource                    |
| ListOpen                 | Open a resource in the Azure portal           |
| ListRefresh              | Refresh a list                                |
| ListUpdate               | Open JSON editor to allow updating a resource |

## Keys

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

## Editing Content

For items in the tree that are editable (i.e. have a `PUT` endpoint), the `ListUpdate` action will open an editor for you to make changes and then issue the `PUT` request to update the item once you have closed the file. By default this is configured to use [Visual Studio Code](https://code.visualstudio.com).

If you wish to override the default editor, create a `~/.azbrowse-settings.json` file (where `~` is your users home directory).

The file should be formated like so:

```json
{
    "editor": {
        "command": {
            "executable": "code",
            "args" : [ "--wait" ]
        },
        "translateFilePathForWSL" : true,
        "tempDir" : "~/tmp",
        "revertToStandardBuffer" : false
    }
}
```

The `command` has two parts, `executable` and `args`. The filename to edit is automatically appended to the `args`. In the example above you can see the `--wait` argument specified which instructs the VS Code executable to not exit until the file is closed. This is important as it is how azbrowse determines that you have finished editing the file and it should perform the `PUT` request.

The `translateFilePathForWSL` property controls whether the path should be converted from a Linux path to a Windows path using the `wslpath` command. This is useful when running azbrowse under [WSL](https://docs.microsoft.com/en-us/windows/wsl/about) and a Windows editor as it converts the temp file path to one that the Windows application can understand.

The `tempDir` property lets you control where the temporary JSON files are written should you have a requirement to not put them in the OS temp location.

The `revertToStandardBuffer` property controls whether the terminal is reverted to the standard buffer (azbrowse uses the alternate buffer for display) when editing files. Set this to `true` when configuring terminal-based editors.

The default configuration uses VS Code as the editor as shown above, and dynamically determines whether to perform the WSL file path translation.

Another example:

### Examples

#### vim

```json
{
    "editor": {
        "command": {
            "executable": "vim"
        },
        "revertToStandardBuffer" : true
    }
}
```

#### notepad (on Windows)

```json
{
    "editor": {
        "command": {
            "executable": "notepad.exe"
        }
    }
}
```

#### notepad (under WSL)

```json
{
    "editor": {
        "command": {
            "executable": "notepad.exe"
        },
        "translateFilePathForWSL" : true,
    }
}
```
