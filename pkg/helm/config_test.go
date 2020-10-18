package helm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadGeneratorConfig(t *testing.T) {
	f, err := os.Open(filepath.Join(currDir, "../../example/invalid-requirements-lock/chartref.yaml"))
	require.NoError(t, err)
	defer f.Close()
	cfg, err := ReadGeneratorConfig(f)
	require.NoError(t, err)
	require.NotNil(t, cfg, "result")
}

func TestReadGeneratorConfigUnsupportedField(t *testing.T) {
	f, err := os.Open(filepath.Join(currDir, "../../example/unsupported-field/chartref.yaml"))
	require.NoError(t, err)
	defer f.Close()
	_, err = ReadGeneratorConfig(f)
	require.Error(t, err)
}