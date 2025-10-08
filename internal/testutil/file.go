package testutil

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func CopyFile(t testing.TB, src, dst string) {
	t.Helper()
	sourceFile, err := os.Open(src)
	require.NoError(t, err)
	defer sourceFile.Close() //nolint:errcheck

	destFile, err := os.Create(dst)
	require.NoError(t, err)
	defer destFile.Close() //nolint:errcheck

	_, err = io.Copy(destFile, sourceFile)
	require.NoError(t, err)
}

func ReadFile(t testing.TB, src string) []byte {
	t.Helper()
	b, err := os.ReadFile(src)
	require.NoError(t, err)
	return b
}
