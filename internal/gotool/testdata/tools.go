//go:build tools

package testdata

import (
	_ "golang.org/x/tools/cmd/deadcode"
	_ "golang.org/x/tools/cmd/goimports"
)
