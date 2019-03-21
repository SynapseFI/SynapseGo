# SynapseFI Go Library (BETA)
Go-based API wrapper for Synapse REST API. This library handles the user authentication process. As long as the user's fingerprint is registered, further authentication is not necessary in the development flow.

## Documentation

[Synapse Docs](https://docs.synapsefi.com/)

## Installation
```bash
$ go get github.com/SynapseFI/SynapseGo
```

**main.go**
```go
import github.com/SynapseFI/SynapseGo
```

## Examples

Refer to [samples](samples/samples.md) and our [API documentation](https://docs.synapsefi.com/) for examples.

## Testing

Functions that mock the Synapse API:

```bash
go test -v --tags mock
```

Other functions including (limited) API requests:

```bash
go test -v
```
