# gocommon

Golang common libraries.

## Installation

`go get -u github.com/quangdangfit/gocommon`

Note that zap only supports the two most recent minor versions of Go.

## Quick Start

#### Logger:
```go
logger.Initialize("production")
logger.Info("This info log")
```

## Development Status: Stable

All APIs are finalized, and no breaking changes will be made in the 1.x series
of releases. Users of semver-aware dependency management systems should pin
zap to `^1`.

## Contributing

If you want to contribute to this boilerplate, clone the repository and just start making pull requests.