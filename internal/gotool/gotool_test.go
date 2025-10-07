package gotool

import (
	"path/filepath"
	"testing"

	"github.com/Arthur1/go-tools-migrator/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	const wantNewGoModContent = `module github.com/Arthur1/go-tools-migrator/internal/gotool/testdata

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
`

	t.Run("dryrun", func(t *testing.T) {
		tmpDir := t.TempDir()
		toolsGoPath := filepath.Join(tmpDir, "tools.go")
		goModPath := filepath.Join(tmpDir, "go.mod")
		testutil.CopyFile(t, "./testdata/tools.go", toolsGoPath)
		testutil.CopyFile(t, "./testdata/go.mod", goModPath)

		newGoModContent, err := Migrate(toolsGoPath, goModPath, true)
		assert.NoError(t, err)
		assert.Equal(t, wantNewGoModContent, newGoModContent)

		originalGoModBytes := testutil.ReadFile(t, "./testdata/go.mod")
		tmpGoModBytes := testutil.ReadFile(t, goModPath)
		assert.Equal(t, originalGoModBytes, tmpGoModBytes)

		assert.FileExists(t, toolsGoPath)
	})

	t.Run("non-dryrun", func(t *testing.T) {
		tmpDir := t.TempDir()
		toolsGoPath := filepath.Join(tmpDir, "tools.go")
		goModPath := filepath.Join(tmpDir, "go.mod")
		testutil.CopyFile(t, "./testdata/tools.go", toolsGoPath)
		testutil.CopyFile(t, "./testdata/go.mod", goModPath)

		newGoModContent, err := Migrate(toolsGoPath, goModPath, false)
		assert.NoError(t, err)
		assert.Equal(t, wantNewGoModContent, newGoModContent)

		tmpGoModBytes := testutil.ReadFile(t, goModPath)
		assert.Equal(t, wantNewGoModContent, string(tmpGoModBytes))

		assert.NoFileExists(t, toolsGoPath)
	})
}
