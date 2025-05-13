![Go](https://img.shields.io/badge/shared%20library-Go-blue) ![License](https://img.shields.io/github/license/aaronlmathis/gosight-shared)

# GoSight Shared

This repository contains the shared types, proto files, and utilities used by both `gosight-agent` and `gosight-server`.

## What's Included

- `proto/` – Protobuf definitions (`metric.proto`, `log.proto`, `meta.proto`)
- `model/` – Internal Go structs for metrics, logs, metadata, events
- `utils/` – Common utility functions for time, tags, logging, etc.

## Used by

- [`gosight-agent`](https://github.com/aaronlmathis/gosight-agent)
- [`gosight-server`](https://github.com/aaronlmathis/gosight-server)

## Usage

Import in Go modules:

```go
import "github.com/aaronlmathis/gosight-shared/model"
```

## License

GPL-3.0-or-later
