# SynapseFI Go Library
![status](https://img.shields.io/badge/status-beta-yellow.svg)

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

Refer to [examples](examples/examples.md) and our [API documentation](https://docs.synapsefi.com/) for examples.

## Testing

Functions that mock the Synapse API:

```bash
make test-mock
```

Other functions including (limited) API requests:

```bash
make test-api
```
