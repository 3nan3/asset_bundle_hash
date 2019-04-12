package pkg

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadAssetsSuccess(t *testing.T) {
	filePath := "testdata/sample.manifest"
	expected := []string{"pkg/testdata/sample_1.txt", "pkg/testdata/sample_2.txt"}

	manifestFile := manifestFile{filePath}
	assets, err := manifestFile.readAssets()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, assets, expected, "Expected %s, but returns %s", expected, assets)
}

func TestReadAssetsInvalidFormat(t *testing.T) {
	filePath := "testdata/fatal_sample.manifest"

	manifestFile := manifestFile{filePath}
	_, err := manifestFile.readAssets()
	assert.Error(t, err, "Expected to raise error, but not raised")
}
