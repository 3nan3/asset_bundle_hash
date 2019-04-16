package cmd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	manifestFile := "../pkg/testdata/sample.manifest"
	projectRoot := "../"

	rootCmd.SetArgs([]string{"-d", projectRoot, manifestFile})
	stdout, stderr, err := captureStdio(func() {
		Execute()
	})
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, stdout)
	assert.Empty(t, stderr)
}
