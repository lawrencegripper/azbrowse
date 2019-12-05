# Command Line Arguments

Once you have [installed](../README.md#install) azbrowse you should simply be able to run `azbrowse` from your terminal/command line but there are some arguments that you can pass to customise the behaviour.

You can run `azbrowse --help` to get a summary of the arguments and this document gives a bit more detail.

## Demo mode

If you're running azbrowse as part of a demo, then the `--demo` argument will mask (some of) the potentially sensitive values that can be displayed in azbrowse

## Controlling the Azure tenant that is loaded

Passing the `--tenant-id` argument allows you to control the Azure Active Directory tenant that azbrowse uses to load the subscriptions.

Running `az account list --query "[].{name:name, tenantId:tenantId}" -o table` will give you a list of subscriptions and their associated tenant. Then you can pass the tenant to azbrowse, e.g. `azbrowse --tenant-id 00000000-0000-0000-0000-000000000000`

## Navigating to resources

The `--navigate` argument allows you to pass the ID of a resource to navigate to. See [Getting Started](./getting-started.md) for more info on this.

## Debug and Fuzzer

The `--debug` argument changes the behaviour to aid debugging (e.g. extending timeouts)

The `--fuzzer` argument runs the fuzzer to automatically navigate through the UI, e.g. `azbrowse --fuzzer 10` to run it for 10 minutes.
