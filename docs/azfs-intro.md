# Getting started

Once you have [azbrowse installed](../README.md#install) and [fuse installed](https://en.wikipedia.org/wiki/Filesystem_in_Userspace) you should be able to run `azbrowse azfs -mount /mnt/azfs -accept-risk -sub subscriptionname` from your terminal/command line.

Additional command line options:
- `-edit` Enables edit mode so files can be updated or deleted
- `-demo` Similar to the `azbrowse -demo` attempts to strip sensitive data

## Risks

This is alpha quality at the moment, to minimize risk please use with caution and consider using `-sub` to filter to a single non-production subscription.

The following risks are present:

1. This feature may cause loss of damage to configuration/resources in to your subscriptions. 
1. Once mounted with `-edit` option enabled any updates/deletes will be submitted to the Azure API removing or editing your resource definitions. There is no "Are you sure?" step. Any tooling which can interact with the filesystem can delete Resource Groups or Resources in this mode. 
1. In `-edit` mode it's possible to trample on another persons change as the filesystem reads the current state, changes it and then applies it. If during this time other changes are made to the object by another user they will be lost (caching of the file content makes this time window large //Todo: Investigate periodically invalidating the cache).
1. Tools which walk the file system (VSCode is one) will cause a large number of calls to the Azure APIs. `azfs` attempts to limit requests to a max of 10 per second. Be aware that it's possible for a search indexer, grep or VSCode to make lots of API calls and result in throttling on your subscription or provider under it. 

# Known errors

If you see `mount helper error: fusermount: failed to access mountpoint /mnt/azfs: Transport endpoint is not connected` this indicates the last instance didn't clean up after itself on exit. Exit any applications of terminals currently viewing the mounted filesystem then use `fusermount -u /mnt/azfs` (replace /mnt/azfs with your mount location) to unmount it. Once completed restart `azfs`.

## Can I help?

Yes please, PRs, Issues and thoughts welcome. If you can help mitigate one of the above risks or improve other aspects of `azfs` I'd love to hear from you. 

## What does this look like?

![azfsdemo](https://user-images.githubusercontent.com/1939288/76356245-88db0080-630d-11ea-9137-f3b69f070676.gif)

