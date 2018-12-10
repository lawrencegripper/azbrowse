[![Go Report Card](https://goreportcard.com/badge/github.com/yangl900/armclient-go)](https://goreportcard.com/report/github.com/yangl900/armclient-go) [![Build Status](https://travis-ci.org/yangl900/armclient-go.svg?branch=master)](https://travis-ci.org/yangl900/armclient-go)
# armclient
A simple command line tool to invoke the Azure Resource Manager API from any OS. Inspired by original windows version ARMClient (https://github.com/projectkudu/ARMClient).

# Why we need this
I always loved the windows version ARMClient. It is super useful when exploring Azure Resource Manager APIs. You just work with REST API directly and with `--verbose` flag you can see the raw request / response with headers.

While I started working on non-windows platform, there isn't such thing available. Existing ARMClient code is based on full .Net framework and winform. Porting to .Net Core requires siginificant changes. You can do curl but it's too much work and you need handle the Azure AD login manually. So I decided to implement one in Golang and release it in Windows, Linux and MacOS.

## Highlights
* Integrated with Azure Cloud Shell. When running armclient in Cloud Shell, sign-in will be taken care automatically. No sign in needed, just go! (you do need install)

[![Launch Cloud Shell](https://shell.azure.com/images/launchcloudshell.png "Launch Cloud Shell")](https://shell.azure.com)

# Installation
armclient is just one binary, you can copy it to whereever you want and use it.

For Linux:
```bash
curl -sL https://github.com/yangl900/armclient-go/releases/download/v0.2.2/armclient-go_0.2.2_linux_64-bit.tar.gz | tar xz
```

For Windows (In PowerShell):
```powershell
curl https://github.com/yangl900/armclient-go/releases/download/v0.2.2/armclient-go_0.2.2_windows_64-
bit.tar.gz -OutFile armclient.tar.gz
```
And unzip the tar.gz file, only binary needed is armclient.exe

For MacOS:
```bash
curl -sL https://github.com/yangl900/armclient-go/releases/download/v0.2.2/armclient-go_0.2.2_macOS_64-bit.tar.gz | tar xz
```

# How to use it
It's very straight-forward. Syntax is exactly same as original ARMClient. To *GET* your subscriptions, simply do

```
armclient get /subscriptions?api-version=2018-01-01
```

Output is simply JSON returned from Azure Resource Manager endpoint, e.g.
```json
{
  "value": [
    {
      "id": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
      "subscriptionId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
      "displayName": "Visual Studio Ultimate with MSDN",
      "state": "Enabled",
      "subscriptionPolicies": {
        "locationPlacementId": "Public_2014-09-01",
        "quotaId": "MSDN_2014-09-01",
        "spendingLimit": "On"
      }
    }
  ]
}
```
If more details of the request needed, add `--verbose` flag and here you go
```
---------- Request -----------------------

GET https://management.azure.com/subscriptions?api-version=2015-01-01
Host: management.azure.com
Authorization: Bearer eyJ0eXAi...
User-Agent: github.com/yangl900/armclient-go
Accept: application/json
x-ms-client-request-id: 9e6cceb1-8a4e-40eb-9701-11d341150220

---------- Response (215ms) ------------

HTTP/1.1: 200 OK
cache-control: no-cache
pragma: no-cache
expires: -1
x-ms-request-id: 64e0fc41-98a3-42c4-808a-ef2fcb7e688c
x-ms-correlation-request-id: 64e0fc41-98a3-42c4-808a-ef2fcb7e688c
x-ms-routing-request-id: WESTUS:20180207T075009Z:64e0fc41-98a3-42c4-808a-ef2fcb7e688c
date: Wed, 07 Feb 2018 07:50:08 GMT
content-type: application/json; charset=utf-8
strict-transport-security: max-age=31536000; includeSubDomains
vary: Accept-Encoding
x-ms-ratelimit-remaining-tenant-reads: 14998

{
  "value": [
    {
      "id": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
      "subscriptionId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
      "displayName": "Visual Studio Ultimate with MSDN",
      "state": "Enabled",
      "subscriptionPolicies": {
        "locationPlacementId": "Public_2014-09-01",
        "quotaId": "MSDN_2014-09-01",
        "spendingLimit": "On"
      }
    }
  ]
}
```

To print out current tenant access token claims, do
```
armclient token
```

Output looks like following, token will also be copied to clipboard automatically (if available). Linux environment requires `xclip` installed for the clipboard copy.
```json
{
  "aud": "https://management.core.windows.net/",
  "iss": "https://sts.windows.net/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx/",
  "iat": 1518072605,
  "nbf": 1518072605,
  "exp": 1518076505,
  "acr": "1",
  "appid": "04b07795-8ddb-461a-bbee-02f9e1bf7b46",
  "appidacr": "0",
  "idp": "live.com",
  "name": "Anders Liu",
  "scp": "user_impersonation",
  "ver": "1.0"
}
```

To print out the raw JWT token, do
```
armclient token -r
```

To print the access token of a different tenant:
```
armclient token --tenant {tenantId or name}
```

## Input for request body
2 ways to specify an input for request body, take create resource group as an example. You can do one of following 2 ways:

1. inline the request body in command line
```
armclient put /subscriptions/{subscription}/resourceGroups/{resourceGroup}?api-version=2018-01-01 "{'location':'westus'}"
```
2. save the request body in a JSON file and use @<file-path> as parameter
```
armclient put /subscriptions/{subscription}/resourceGroups/{resourceGroup}?api-version=2018-01-01 @./resourceGroup.json
```

## Add additional request headers
Use flag `--header` or `-H` for additional request headers. For example:

```bash
armclient get /subscriptions?api-version=2018-01-01 -H Custom-Header=my-header-value-123 --verbose
```

## Target ARM endpoint in specific region
Absolute Uri is accepted, so just specify the complete Uri:

```
armclient get https://westus.management.azure.com/subscriptions?api-version=2018-01-01
```

## Working with multiple Azure AD Directories (tenants)
To list all tenants you have access to:
```bash
armclient tenant list
```

To set a tenant as active tenant(default is the first tenant):
```bash
armclient tenant set {tenantID}
```

To show current active tenant:
```bash
armclient tenant show
```

# Exploring Azure APIs
More REST APIs please see Azure REST API document. The original [ARMClient wiki](https://github.com/projectkudu/ARMClient/wiki) also has pretty good documentation.

# Contribution
Build the project
```
make
```

Add dependency
```
dep ensure
```
