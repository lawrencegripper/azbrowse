## azbrowse

An interactive CLI for browsing Azure

```
azbrowse [flags]
```

### Options

```
      --debug                 run in debug mode
      --demo                  run in demo mode to filter sensitive output
      --fuzzer int            run fuzzer (optionally specify the duration in minutes) (default -1)
  -h, --help                  help for azbrowse
  -m, --mouse                 (optional) enable mouse support. Note this disables normal text selection in the terminal
  -n, --navigate string       (optional) navigate to resource by resource ID
  -r, --resume                (optional) resume navigating from your last session
  -s, --subscription string   (optional) specify a subscription to load
      --tenant-id string      (optional) specify the tenant id to get an access token for (see az account list -o json)
```

### SEE ALSO

* [azbrowse azfs](azbrowse_azfs.md)	 - Mount the Azure ARM API as a fuse filesystem
* [azbrowse completion](azbrowse_completion.md)	 - Generates shell completion scripts
* [azbrowse version](azbrowse_version.md)	 - Print version information

