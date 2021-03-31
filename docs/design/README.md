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
  - [Status and Notifications](#status-and-notifications)
  - [Automation, Error Handling and Autocomplete](#automation-error-handling-and-autocomplete)

## Overview

At it's core `azbrowse` is a `tree` which the user walks. The `root` node uses the `az cli` to discover all 
the subscriptions a user has access too. From this point the user can walk down the tree to `resources`, `subresources` and `actions`. 

[![](https://mermaid.ink/img/eyJjb2RlIjoiZ3JhcGggVERcbiAgICBBW1Jvb3RdIC0tPiBCKFN1YnNjcmlwdGlvbnMpXG4gICAgQiAtLT4gQyhSZXNvdXJjZSBHcm91cHMpXG4gICAgQyAtLT58UmVzb3VyY2UgRXhwYW5kZXJ8IERbU3RvcmFnZSBBY2NvdW50IFhZXVxuICAgIEMgLS0-fFJlc291cmNlIEV4cGFuZGVyfCBGW1NlcnZpY2VCdXMgQWNjb3VudCBaWF1cbiAgICBDIC0tPnxBZGRpdGlvbmFsIEV4cGFuZGVyfCBHW01ldHJpY3MgWlhdXG4gICAgRCAtLT58QWN0aW9ufCBIW2xpc3RLZXlzXVxuICAgIEQgLS0-fEFjdGlvbnwgWVtyZWdlbmVyYXRlS2V5c10iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiZ3JhcGggVERcbiAgICBBW1Jvb3RdIC0tPiBCKFN1YnNjcmlwdGlvbnMpXG4gICAgQiAtLT4gQyhSZXNvdXJjZSBHcm91cHMpXG4gICAgQyAtLT58UmVzb3VyY2UgRXhwYW5kZXJ8IERbU3RvcmFnZSBBY2NvdW50IFhZXVxuICAgIEMgLS0-fFJlc291cmNlIEV4cGFuZGVyfCBGW1NlcnZpY2VCdXMgQWNjb3VudCBaWF1cbiAgICBDIC0tPnxBZGRpdGlvbmFsIEV4cGFuZGVyfCBHW01ldHJpY3MgWlhdXG4gICAgRCAtLT58QWN0aW9ufCBIW2xpc3RLZXlzXVxuICAgIEQgLS0-fEFjdGlvbnwgWVtyZWdlbmVyYXRlS2V5c10iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)

Everything (or nearly everything) is represented as `expander`'s which return `ExpanderResult`'s. 

An `ExpanderResult` contains the content to display in the `itemView` and also `treeNode`'s each of which represents and item in the list of the left panel.

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

## Expanders, ExpanderResults and TreeNodes

WIP Sections

### Expanders

The hierarchical drill-down from Subscription -> Resource Group -> Resource -> ... is driven by Expanders. These are registered in `registerExpanders.go` and when the list widget expands a node it calls each expander asking if they have any nodes to provide. Multiple expanders can return nodes for any given parent node, but only one expander should mark the response as the primary response.

Each node has an ID and IDs should be unique (to support the `--navigate` command), and typically are the resource ID for the resource in Azure (this allows the `open in portal` action to function)

### APISets

The `SwaggerResourceExpander` is used to drill down within resources. It works against `SwaggerAPISet`s which provide the swagger metadata as well as encapsulating access to the the endpoints identified in the metadata.

The default API Set is `SwaggerAPISetARMResources` which is based on code generated at build time via `make swagger-codegen`. The swagger codegen process loads all of the manamgement plane swagger documents published on GitHub and builds a hierarchy based on the URLs. This is then distilled down into a slightly simpler format based around the `ResourceType` struct. Access to the endpoints in `SwaggerAPISetARMResources` is performed by the `armclient` which piggy-backs on the authentication from the Azure CLI.

Other API Sets can be registered and currently containerService and search are two examples. The Azure Search API Set also uses a `ResourceType` hierarchy generated at build time, but it is dynamically registered with the `SwaggerResourceExpander` when the user expands the "Search Service" node (added by the `AzureSearchServiceExpander`). The API Set instance that is registered at that point has the credentials for authenticating to that specific instance of the Azure Search Service.

The pattern for the container Service API Set is similar: a Kubernetes API node is added by the `AzureKubernetesServiceExpander` and when that is expanded the credentials to the Kubernetes cluster are retrieved and passed to an instance of the API Set. One difference is that the `ResourceType`s for the container service API Set are generated at runtime by querying the Kubernetes API (this allows the node expansion to accurately represent the cluster version as well as any other endpoints that are specific to the cluster)

Issuing `PUT`/`DELETE` requests requires the same authentication as `GET` requests so the `SwaggerResourceExpander` also forwards these to the relevant API Set. (The metadata for the node contains the name of the API Set that returned it)

## Key bindings

Key bindings are initialised in the `setupViewsAndKeybindings` function in `main.go`. Each binding is registered via the `keybindings.AddHandler` function and subsequently bound through the `keybindings.Bind` function.

Handlers implement the `KeyHandler` interface which specifies an ID, an implementation to invoke, the widget that the binding is scoped to, and the default key.

The `ID` and `DefaultKey` functions are both provided by `KeyHandlerBase`. `ID` simply returns the `id` property, and `DefaultKey` performs a lookup in `DefaultKeys` using the ID.

## Views and GoCUI

WIP Section covering views 

## Status and Notifications

WIP Section

## Automation, Error Handling and Autocomplete

WIP Section