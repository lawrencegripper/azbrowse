## azbrowse completion

Generates shell completion scripts

### Synopsis

To load completion run
	
	. <(azbrowse completion SHELL)
	Valid values for SHELL are : bash, fish, powershell, zsh
	
	For example, to configure your bash shell to load completions for each session add to your bashrc
	
	# ~/.bashrc or ~/.profile
	source <(azbrowse completion bash)
	

```
azbrowse completion SHELL [flags]
```

### Options

```
  -h, --help   help for completion
```

### SEE ALSO

* [azbrowse](azbrowse.md)	 - An interactive CLI for browsing Azure

