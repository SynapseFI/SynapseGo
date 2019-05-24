# Changelog

All notable changes to this project will be documented in this file.

## 2019-05-23

### Added
- [User] `GetSubnets` to get a list of user subnets
- Examples
  + [Client]
    * `GetTradeMarkeData`
    * `GetCryptoMarketData`
    * `GetCryptoQuotes`
    * `GetInstitutions`
    * `LocateATMs`
    * `VerifyRoutingNumber`
  + [User]
    * `GetRefreshToken`
    * `Select2FA`
    * `SubmitMFA`
    * `VerifyPIN`
    * `GetNodeStatements`
    * `GetStatements`
    * `CreateNodeStatements`
    * `GetSubnets`
    * `CreateDummyTransactions`

### Changed
- Renamed `samples/` to `examples/`
- Renamed `samples.md` to `examples.md`

---

## 2019-05-21

### Added
- [Client] `GetTradeMarketData` to view market data on a specific stock
- [Client] `VerifyRoutingNumber` to verify the bank details of a routing number
- [User] `CreateNodeStatements` to generate ad-hoc statements for a user node

### Changed
- New user sessions now require a one-time authentication using `Authenticate`
- `Authenticate` requires `fingerprint` and `ipAddress`
- `GetUser` no longer requires `fingerprint` and `ipAddress`
- Updated `README` to reflect new changes