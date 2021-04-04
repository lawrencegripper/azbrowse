# Design Docs

This section of the docs is designed to help people new to the code understand 
how `azbrowse` is written, where stuff is and what we're proud of vs regret!

- [Design Docs](#design-docs)
  - [Overview](#overview)
  - [Talking to Azure](#talking-to-azure)
  - [Expanders, ExpanderResults and TreeNodes](#expanders-expanderresults-and-treenodes)
    - [Expanders](#expanders)
    - [APISets](#apisets)
  - [Key bindings](#key-bindings)
  - [Views and GoCUI](#views-and-gocui)
    - [`itemView`](#itemview)
    - [`commandPanel`](#commandpanel)
    - [`notifications`](#notifications)
    - [`list`](#list)
    - [`statusbar`](#statusbar)
    - [`help`](#help)
  - [Status and Notifications](#status-and-notifications)
  - [Automation, Error Handling and Autocomplete](#automation-error-handling-and-autocomplete)

## Overview

At it's core `azbrowse` is a `tree` which the user walks. The `root` node uses the `az cli` to discover all 
the subscriptions a user has access too. From this point the user can walk down the tree to `resources`, `subresources` and `actions`. 

[![](https://mermaid.ink/img/eyJjb2RlIjoiZ3JhcGggVERcbiAgICBBW1Jvb3RdIC0tPiBCKFN1YnNjcmlwdGlvbnMpXG4gICAgQiAtLT4gQyhSZXNvdXJjZSBHcm91cHMpXG4gICAgQyAtLT58UmVzb3VyY2UgRXhwYW5kZXJ8IERbU3RvcmFnZSBBY2NvdW50IFhZXVxuICAgIEMgLS0-fFJlc291cmNlIEV4cGFuZGVyfCBGW1NlcnZpY2VCdXMgQWNjb3VudCBaWF1cbiAgICBDIC0tPnxBZGRpdGlvbmFsIEV4cGFuZGVyfCBHW01ldHJpY3MgWlhdXG4gICAgRCAtLT58QWN0aW9ufCBIW2xpc3RLZXlzXVxuICAgIEQgLS0-fEFjdGlvbnwgWVtyZWdlbmVyYXRlS2V5c10iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiZ3JhcGggVERcbiAgICBBW1Jvb3RdIC0tPiBCKFN1YnNjcmlwdGlvbnMpXG4gICAgQiAtLT4gQyhSZXNvdXJjZSBHcm91cHMpXG4gICAgQyAtLT58UmVzb3VyY2UgRXhwYW5kZXJ8IERbU3RvcmFnZSBBY2NvdW50IFhZXVxuICAgIEMgLS0-fFJlc291cmNlIEV4cGFuZGVyfCBGW1NlcnZpY2VCdXMgQWNjb3VudCBaWF1cbiAgICBDIC0tPnxBZGRpdGlvbmFsIEV4cGFuZGVyfCBHW01ldHJpY3MgWlhdXG4gICAgRCAtLT58QWN0aW9ufCBIW2xpc3RLZXlzXVxuICAgIEQgLS0-fEFjdGlvbnwgWVtyZWdlbmVyYXRlS2V5c10iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)

Everything (or nearly everything) is represented as `expander`'s which return `ExpanderResult`'s. 

An `ExpanderResult` contains the content to display in the `itemView` and also `treeNode`'s each of which represents and item in the list of the left panel.

Everything gets started in `cmd/azbrowse/cmd.go`

## Style/Qwerks

1. Lots of strings are composed in the code. The `+` concat style is preferred throughout the codebase over `fmt` and `%s` as it was my (Lawrence) preference as I believe it's more readable when things started out with `go 1.13`.
1. The `Status` notifications are an overly complex `pub/sub` system that I'd not use if I had a do-over. I'll cover this later in the doc.

... Probably more 

## Talking to Azure

When talking to Azure in `azbrowse` you'll use the `armclient` helpers.

Early on we made a choice **not** to use the Golang SDK for Azure. 

Instead we talk directly to the API's of 
the individual [`resourceProviders`](https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/resource-providers-and-types) using the code in [`armclient`](../../pkg/armclient).

Using the REST methods directly via the helpers in `armclient` allows us to do the following:
- Meta-programming: We use the [Azure API Specs](https://github.com/Azure/azure-rest-api-specs) to generate the `tree` the user is going to walk
- Flexibility: We're able to use endpoints which may not be represented correctly or at all in the Go SDK for Azure

The `armclient` provides some helpers to make this easier. 

1. Requests the latest `api-versions` for each `resourceProvider` and appends this to requests
2. Handles throttling to ensure we don't make too many requests to the ARM API and get the user throttled
3. Allows async requests using Golang `channels`
4. Helpers for querying the [`azureResourceGraph`](https://azure.microsoft.com/en-us/features/resource-graph/) this is used to get status from `resources` in Azure
5. Handle mocking and testing allowing authors of `expanders` to create mock responses to test their code

Sometimes calls to Azure are expensive (take a while and return lots of data), when working with these calls we cache results with a `TTL` using ``internal/pkg/storage/store.go`. 

`PopulateResourceAPILookup` in `armclient` is an example of this caching approach.

## Expanders, ExpanderResults and TreeNodes

### TreeNodes

These represent items in the list panel, mapping to Azure Resources, Subresources or Actions.

The items have has evolved over a year, it has some duplication and fields which are used differently by different expanders. 

```go
// TreeNode is an item in the ListWidget
type TreeNode struct {
	Parentid               string                // The ID of the parent resource
	Parent                 *TreeNode             // Reference to the parent node
	ID                     string                // The ID of the resource in ARM
	Name                   string                // The name of the object returned by the API
	Display                string                // The Text used to draw the object in the list
	ExpandURL              string                // The URL to call to expand the item
	ItemType               string                // The type of item either subscription, resourcegroup, resource, deployment or action
	ExpandReturnType       string                // The type of the items returned by the expandURL
	DeleteURL              string                // The URL to call to delete the current resource
	Namespace              string                // The ARM Namespace of the item eg StorageAccount
	ArmType                string                // The ARM type of the item eg Microsoft.Storage/StorageAccount
	Metadata               map[string]string     // Metadata is used to pass arbritray data between `Expander`'s
	SubscriptionID         string                // The SubId of this item
	StatusIndicator        string                // Displays the resources status
	SwaggerResourceType    *swagger.ResourceType // caches the swagger ResourceType to avoid repeated lookups
	Expander               Expander              // The Expander that created the node (set automatically by the list)
	SuppressSwaggerExpand  bool                  // Prevent the SwaggerResourceExpander attempting to expand the node
	SuppressGenericExpand  bool                  // Prevent the DefaultExpander (aka GenericExpander) attempting to expand the node
	TimeoutOverrideSeconds *int                  // Override the default expand timeout for a node
	ExpandInPlace          bool                  // Indicates that the node is a "More..." node. Must be the last in the list and will be removed and replaced with the expanded nodes
}
```

Simple expanders may use a little as the following to construct a `treeNode` instance:

```go
&TreeNode{
			Display:        sub.DisplayName,
			Name:           sub.DisplayName,
			ID:             sub.ID,
			ExpandURL:      sub.ID + "/resourceGroups?api-version=2018-05-01",
			ItemType:       SubscriptionType,
			SubscriptionID: sub.SubscriptionID,
		}
```

A more complex example would be the `metrics` expander. This constructs a complex `ExpandURL` while also suporessing other expanders for `Swagger` and `DefaultExpander`. 

Another interesting case here is the use of `Metadata` map, this allows expanders to store arbritrary data which they, or another expander, can reuse when expanding the item (see `internal/pkg/expanders/metrics.go` for full code). 

```go
&TreeNode{
			Name:     metric.Name.Value,
			Display:  metric.Name.Value + "\n  " + style.Subtle("Unit: "+metric.Unit),
			ID:       currentItem.Metadata["ResourceID"] + "/providers/microsoft.Insights/metrics",
			Parentid: currentItem.ID,
			ExpandURL: currentItem.Metadata["ResourceID"] + "/providers/microsoft.Insights/metrics?timespan=" +
				time.Now().UTC().Add(-4*time.Hour).Format("2006-01-02T15:04:05.000Z") + "/" +
				time.Now().UTC().Format("2006-01-02T15:04:05.000Z") + "&interval=PT1M&metricnames=" +
				url.QueryEscape(metric.Name.Value) + "&aggregation=" +
				url.QueryEscape(metric.PrimaryAggregationType) +
				"&metricNamespace=" + url.QueryEscape(metric.Namespace) +
				"&autoadjusttimegrain=true&validatedimensions=false&api-version=2018-01-01",
			ItemType:              "metrics.graph",
			SubscriptionID:        currentItem.SubscriptionID,
			SuppressSwaggerExpand: true,
			SuppressGenericExpand: true,
			Metadata: map[string]string{
				"AggregationType": strings.ToLower(metric.PrimaryAggregationType),
				"Units":           strings.ToLower(metric.Unit),
			},
```

In this case the metadata is used to create a caption when rendering the graph for the item when expanded:

```go
caption := style.Title(currentItem.Name) +
		style.Subtle(" (Aggregate: '"+currentItem.Metadata["AggregationType"]+"' Unit: '"+
			currentItem.Metadata["Units"]+"')")
```

### Expanders

The hierarchical drill-down from Subscription -> Resource Group -> Resource -> ... is driven by Expanders. These are registered in `registerExpanders.go` and when the list widget expands a node it calls each expander asking if they have any nodes to provide. Multiple expanders can return nodes for any given parent node, but only one expander should mark the response as the primary response.

> Note: Remember to update `InitializeExpanders` in `internal/pkg/expanders/registerExpanders.go` with your new expander if you add one otherwise it won't get called!

Each node has an ID and IDs should be unique (to support the `--navigate` command), and typically are the resource ID for the resource in Azure (this allows the `open in portal` action to function)

Expanders meet an interface defined in `internal/pkg/expanders/types.go` which is roughly as follows:

```go
// Expander is used to open/expand items in the left list panel
// a single item can be expanded by 1 or more expanders
// each Expander provides two methods.
// `DoesExpand` should return true if this expander can expand the resource
// `Expand` returns the list of sub items from the resource
type Expander interface {
	DoesExpand(ctx context.Context, currentNode *TreeNode) (bool, error)
	Expand(ctx context.Context, currentNode *TreeNode) ExpanderResult
	Name() string // Returns the name of this expander for logging (ie. TenantExpander)
	Delete(context context.Context, item *TreeNode) (bool, error)

	HasActions(ctx context.Context, currentNode *TreeNode) (bool, error)
	ListActions(ctx context.Context, currentNode *TreeNode) ListActionsResult
	ExecuteAction(ctx context.Context, currentNode *TreeNode) ExpanderResult

	CanUpdate(context context.Context, item *TreeNode) (bool, error)
	Update(context context.Context, item *TreeNode, updatedContent string) error

	// Used for testing the expanders
	testCases() (bool, *[]expanderTestCase)
	setClient(c *armclient.Client)
}

```

A good example of a simple expander is the `TenantExpander` at `internal/pkg/expanders/tentant.go`. It uses `armclient` to request
a list of subscriptions for to show the user via the `/subscriptions` Azure REST call. The response is deserialised into a 
go struct for ease and then the list items to show (one for each subscription) are created as `TreeNode`'s. The response
content is set to the API response as JSON or the error received from the call.

### APISets

The `SwaggerResourceExpander` is used to drill down within resources. It works against `SwaggerAPISet`s which provide the swagger metadata as well as encapsulating access to the the endpoints identified in the metadata.

The default API Set is `SwaggerAPISetARMResources` which is based on code generated at build time via `make swagger-codegen`. The swagger codegen process loads all of the manamgement plane swagger documents published on GitHub and builds a hierarchy based on the URLs. This is then distilled down into a slightly simpler format based around the `ResourceType` struct. Access to the endpoints in `SwaggerAPISetARMResources` is performed by the `armclient` which piggy-backs on the authentication from the Azure CLI.

Other API Sets can be registered and currently containerService and search are two examples. The Azure Search API Set also uses a `ResourceType` hierarchy generated at build time, but it is dynamically registered with the `SwaggerResourceExpander` when the user expands the "Search Service" node (added by the `AzureSearchServiceExpander`). The API Set instance that is registered at that point has the credentials for authenticating to that specific instance of the Azure Search Service.

The pattern for the container Service API Set is similar: a Kubernetes API node is added by the `AzureKubernetesServiceExpander` and when that is expanded the credentials to the Kubernetes cluster are retrieved and passed to an instance of the API Set. One difference is that the `ResourceType`s for the container service API Set are generated at runtime by querying the Kubernetes API (this allows the node expansion to accurately represent the cluster version as well as any other endpoints that are specific to the cluster)

Issuing `PUT`/`DELETE` requests requires the same authentication as `GET` requests so the `SwaggerResourceExpander` also forwards these to the relevant API Set. (The metadata for the node contains the name of the API Set that returned it)

### How are expanders called?

`ExpandItemAllowDefaultExpander` in `internal/pkg/expanders/expander.go` handles calling the expanders. 

It does the following:
1. Use `getRegisteredExpanders` to find all expanders
1. Check which expanders are relevant for the `TreeNode` provided using `DoesExpand` on each expander
1. Starts a go routine to call `Expand` asyncronously on each expander which indicated it could expand the item. Results are sent to a go Channel on completion as an `ExpanderResult` see `internal/pkg/expanders/types.go`
1. A timeout is also started to ensure we don't block on a single expander while others have responded
1. Responses are collected from all expanders, by pulling from the go Channel and returned
1. If `AllowDefaultExpander: true` the `DefaultExpander` is used which attempts to make a `GET` request to the `TreeNodes`'s ID. See `internal/pkg/expanders/default.go`

## Key bindings

Key bindings are initialised in the `setupViewsAndKeybindings` function in `main.go`. Each binding is registered via the `keybindings.AddHandler` function and subsequently bound through the `keybindings.Bind` function.

Handlers implement the `KeyHandler` interface which specifies an ID, an implementation to invoke, the widget that the binding is scoped to, and the default key.

The `ID` and `DefaultKey` functions are both provided by `KeyHandlerBase`. `ID` simply returns the `id` property, and `DefaultKey` performs a lookup in `DefaultKeys` using the ID.

## Views and GoCUI

GoCUI draws the terminal interface using `views` (or `widgets`). All the views live in the [internal/pkg/views](../../internal/pkg/views).

### `list`

This view provides the left-hand list view. 

It is responsible for displaying items which can be opened (`TreeNodes` which opened by `expanders`).

This is one of the more complex views in the system, it:

- Handles navigation back/forward in the tree. 
- Keeps a `navStack` which is a first-in first-out history used to go back to previous pages without reloading all items (it tracks the `content` and `TreeNode[]` of the previous items)
- Allows `more ...` style behaviour to incrementally load more items to the list. This is used the `storage` data-plane for when a `container` holds lots of items
- Handles `filtering` items in the list
- Keeps track of the currently selected item and adds the `>` indicator to that item as `up/down` are pressed.

### `itemView`

This view provides the main right-hand output panel displaying the `json`/`yaml`/`xml`/`hcl` content.

It has methods for `GetContent` and `SetContent` for example. 

One of it's responsibilities is to add formatting and highlighting to the the content. 

The content is provided to the views from the [`ExpanderResult`](../../internal/pkg/expanders/types.go) which is produced by the `expanders` we discussed in the sections above.

It also keeps track of a pointer to the `TreeNode` which generated the content. This enables the `GetNode` method
to return a reference to the currently displayed `TreeNode` or `CurrentItem`.

### `commandPanel`

This view is the overlayed command panel (inspired by the CTRL+P panel in VSCode) which allows for more complex
interactions with the UI. 

For example, typing `/` opens the panel and then the user can type text and the `listView` will filter to only
show items that contain the text. 

Alternatively a user can press `CTRL+P` and a small list will show possible options/actions the user can then 
select to invoke.

### `notifications`

This view handles the optional top-right notifications. 

It is used to display pending delete's (resources queued for delete but not yet actioned) alongside 
`toast` style notifications, for example, deletes that have been actioned and we're tracking their
async completion.

The delete functionaility has taken over this view, the aim was to be generalised but it's more specific currently. Methods
for `AddPendingDelete` and `ConfirmDelete` handling logic for deleting items.

Generalised `toast` notifications are driven by events sent via `internal/pkg/eventing/eventing.go` package `IsToast: true`. 
Currently I'm not aware of this being used anywhere so it may not function correctly.

### `statusbar`

This view shows status messages along the bottom of the view. 

It listens to a `pub/sub` style bus of messages published by the `internal/pkg/eventing/eventing.go` package. 

Each message has a `ttl`, `status` etc. It translates these to displayed text color, icons and ensures
they're displayed for the correct amount of time. 

### `help`

This is the simplist view it shows the help for which keys do what.

## Status and Notifications (and a bit of recovery/automation package)

There is an overly complex `pub/sub` system for handling `StatusEvents` and other events in the system. 

```go
// StatusEvent is used to show status information
// in the statusbar
type StatusEvent struct {
	Message    string
	Timeout    time.Duration
	createdAt  time.Time
	InProgress bool
	IsToast    bool
	Failure    bool
	id         uuid.UUID
}
```

This is used by the `statusbar` view, `notifications` view. These events are published, updated and completed as follows:

```go
	statusEvent, done := eventing.SendStatusEvent(&eventing.StatusEvent{
		Message:    "Opening: thing",
		InProgress: true,
	})
	defer done() // This will mark the StatusEvent's InProgress field to false

  // ... Do things here

  event.Failure = true
	event.InProgress = false
	event.Message = "Failed to delete doing thing with error:" + err.Error()
	event.Update()
```

Not great right? There are some helpers hanging off `eventing` which could be updated to make this nicer or the system could be refactored to simplify it.

The same bus is used to push messages when `navigation` occurs, this is used to drive a number of features such as the `--navigate` CLI arg, `make fuzz` fuzzer and `--resume` cli command.

You'll see calls to publish to the bus in the `internal/pkg/views/list.go` ListView like this:

```go
eventing.Publish("list.prenavigate", "GOBACK")
eventing.Publish("list.navigated", ListNavigatedEventState{Success: false})
```

These events power functionality of `--navigate` via the code in `internal/pkg/automation/navigateTo.go` which listens to these events and drives to UI until it reaches the expected item. 

The events are also used in `RegisterGuiAndStartHistoryTracking` part of the `recovery` package `internal/pkg/errorhandling/recovery.go`. In this package the events are used to track the path taken up until a crash/panic occurs to allow it to be easily reproduced.

## Error Handling

All go routines started in the solution have the following first line.

```go
// recover from panic, if one occurrs, and leave terminal usable
	defer errorhandling.RecoveryWithCleanup()
```

This does 2 things. 

1. Ensure we don't leave the terminal session unusable, closing out `gocui` and tidying up
2. Outputting useful information for the user to report the crash like the path they took before the crash, stack etc

## Autocomplete

Inside `cmd/azbrowse/cmd.go` the command structure of the CLI is setup using `cobra`. 

As part of this autocomplete functions are added which do the following:
1. `subscriptionAutocompletion` allows `--subscription stuff<TAB>` --> `--subscription stuffAndThingsSub`
1. `navigateAutocompletion` allows `--navigate /subscription/GUID/resourceg<TAB>` to autocomplete to a resource

To ensure that autocomplete is quick results are cached using the `internal/pkg/storage/store.go` key vault store which also provides a rough `TTL` based expirey (it's a hack - go look ... it works but I'm not proud of it).

## Release and Updating

AzBrowse is built using the `scripts/ci_release.sh` script which runs locally or as part of the Github Action `.github/workflows/build.yaml`.

Under the covers this uses `goreleaser` to build for multiple platforms and packaging systems (brew, docker, deb etc). It is configured in `.goreleaser.yml` and pushes output to Github releases.

The `cmd/azbrowse/selfupdate.go` handles checking for new releases when the CLI started and offering to the user to update. It uses the [`selfupdate` package from rhysd](github.com/rhysd/go-github-selfupdate/selfupdate).