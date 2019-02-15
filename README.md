# SynapseFI Go Library
Go-based API wrapper for Synapse REST API. This library handles the user authentication process. As long as the user's fingerprint is registered, further authentication is not necessary in the development flow.

## Documentation

[Synapse Docs](https://docs.synapsefi.com)

## Installation
```bash
$ go get github.com/synapsefi/synapsego
```

**main.go**
```go
import github.com/synapsefi/synapsego
```

## Examples

Refer to [samples.md](samples/samples.md) and our [API documentation](https://docs.synapsefi.com/v3.1) for examples.

## Testing

Functions that exist as part of the Synapse API:

```bash
cd synapse/
go test -v --tags mock
```

Other functions including (limited) API requests:

```bash
cd synapse/
go test -v
```
