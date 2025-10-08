# go-tools-migrator

go-tools-migrator is a CLI tool that replaces tools management via tools.go with go.mod tool directive (introduced in Go 1.24).

## install

### go install

```
go install github.com/Arthur1/go-tools-migrator/cmd/go-tools-migrator@latest
```

## run

```console
$ tree --noreport
.
├── go.mod
├── go.sum
└── tools.go
$ cat tools.go
//go:build tools

package tools

import (
	_ "golang.org/x/tools/cmd/deadcode"
	_ "golang.org/x/tools/cmd/goimports"
)

$ # dry run
$ go-tools-migrator --dryrun
module sample

go 1.25.0

tool (
        golang.org/x/tools/cmd/deadcode
        golang.org/x/tools/cmd/goimports
)

require golang.org/x/tools v0.37.0

require (
        golang.org/x/mod v0.28.0 // indirect
        golang.org/x/sync v0.17.0 // indirect
        golang.org/x/sys v0.36.0 // indirect
        golang.org/x/telemetry v0.0.0-20250908211612-aef8a434d053 // indirect
)
$ # migrate
$ go-tools-migrator
✅ Succeeded to migrate.
$ tree --noreport
.
├── go.mod
└── go.sum
$ cat go.mod
module sample

go 1.25.0

tool (
        golang.org/x/tools/cmd/deadcode
        golang.org/x/tools/cmd/goimports
)

require golang.org/x/tools v0.37.0

require (
        golang.org/x/mod v0.28.0 // indirect
        golang.org/x/sync v0.17.0 // indirect
        golang.org/x/sys v0.36.0 // indirect
        golang.org/x/telemetry v0.0.0-20250908211612-aef8a434d053 // indirect
)
```
