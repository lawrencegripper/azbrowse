## azbrowse completion

Generates shell completion scripts

### Synopsis

To load completion run
	
	. <(azbrowse completion SHELL)
	Valid values for SHELL are : bash, fish, powershell, zsh
	
	To configure your bash shell to load completions for each session add to your bashrc:
	
	# ~/.bashrc or ~/.profile
	source <(azbrowse completion bash)

	To configure completion for zsh run the following command:

	$ azbrowse completion zsh > "${fpath[1]}/_azbrowse"
	
	Ensure you have 'autoload -Uz compinit && compinit' present in your '.zshrc' file to load these completions

	

```
azbrowse completion SHELL [flags]
```

### Options

```
  -h, --help   help for completion
```

### SEE ALSO

* [azbrowse](azbrowse.md)	 - An interactive CLI for browsing Azure

