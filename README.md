# SynapseFI Go Library
![status](https://img.shields.io/badge/status-beta-yellow.svg)

Go-based API wrapper for Synapse REST API. This library handles the user authentication process. As long as the user's fingerprint is registered, further authentication is not necessary in the development flow.

## Documentation

[Main Docs](https://docs.synapsefi.com/)
[API Reference](https://docs.synapsefi.com/reference)

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

### To run test-mock or test-api: 
  1. Add your credentials to "client_credentials.sample.json" and rename it "client_credentials.json"

### To run test-api:
  1. Open the file "request_test.go" and change the value of "var userID string" to a user that was generated on your platform

Functions that mock the Synapse API:

```bash
make test-mock
```

Other functions including (limited) API requests:

```bash
make test-api
```
