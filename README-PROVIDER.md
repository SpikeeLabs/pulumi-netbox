# Netbox Resource Provider

The Netbox Resource Provider lets you manage [Netbox](https://netbox.dev) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @SpikeeLabs/pulumi-netbox
```

or `yarn`:

```bash
yarn add @SpikeeLabs/pulumi-netbox
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_netbox
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/SpikeeLabs/pulumi-netbox/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package SpikeeLabs.Netbox
```

## Configuration

The following configuration points are available for the `foo` provider:

- `netbox:apiToken` (environment: `FOO_API_KEY`) - the API key for `foo`
- `netbox:hostname` (environment: `FOO_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/netbox/api-docs/).
