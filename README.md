# AzBrowse

An interactive CLI for browsing azure resources, inspired by [resources.azure.com](https://resources.azure.com)

[![Build Status](https://travis-ci.com/lawrencegripper/azbrowse.svg?branch=master)](https://travis-ci.com/lawrencegripper/azbrowse)

# Quick Start

Simply download the binary suitable for your machine, [from the release page](https://github.com/lawrencegripper/azbrowse/releases), and execute it. 

### Status

Basic MVP to prove out the use case. Basic navigation and operations with a boltdb based cache for expensive (slow) API calls.

### Usage

## Navigation 

↑/↓:     Select resource   
Backspace: Go back
ENTER:   Expand/View resource

## Operations

CTRL+o (o for open):    Open Portal at this resource           
DEL:                    Delete resource the currently selected resource (Requires double press to confirm)

![Demo](./docs/quickdemo-azbrowse.gif) 

## Plans

On the TODO list:
 
 - Add ability to bring up pre-filled `az cli` commands for resources
 - Add support for operations on resource (like listing storage account keys)
 - Add editor to update a resources `json` and post back to `rest` endpoint
 - Full cache read-through cache to improve speed (First load from cache then from api) and allow offline use browsing cached data 