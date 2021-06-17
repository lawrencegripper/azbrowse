# Custom Views with Azure Resource Graph Queries

This feature lets you write an Azure Resource Graphy Query as a view to be shown.

This allows you to group, filter and sort resource groups from multiple subscriptions in one view. 

For each query you'd like create a file in `./azurebrowse/queries/` which ends in `.kql`.

## Writing your query

Head into the [portal to author your query, there is a good guide here.](https://docs.microsoft.com/en-us/azure/governance/resource-graph/first-query-portal)

Your query **must** return a list Resource Groups. To fit this the items should have:

`type == 'microsoft.resources/subscriptions/resourcegroups`

The easiest way to achieve this is to include the following `where` clause in your query:

`| where type == 'microsoft.resources/subscriptions/resourcegroups'`

## Simple Example

For a file `./azurebrowse/queries/rgsWithStableInName.kql` with the content

```kusto
// Query finds resources groups beginning with 'stable'
resourcecontainers 
| where type == 'microsoft.resources/subscriptions/resourcegroups' 
| where name contains("stable")
```

When you start up azbrowse you'll see

![](./images/kqlQueries.png)

You can the open the query and you'll see the resource groups returned by the query

![](./images/kqlResults.png)

## Complex Example

// Todo Contributions welcome

