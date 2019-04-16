package pkg

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculateHashSuccess(t *testing.T) {
	manifestFile := "testdata/sample.manifest"
	projectRoot := "../"

	// asset1="pkg/testdata/sample_1.txt"
	// asset2="pkg/testdata/samples/sample_2.txt"
	// hash1=`cat $asset1 | openssl dgst -sha256`
	// hash2=`cat $asset2 | openssl dgst -sha256`
	// meta_hash1=`cat $asset1.meta | openssl dgst -sha256`
	// meta_hash2=`cat $asset2.meta | openssl dgst -sha256`
	// echo -n "${hash1}${asset1}${meta_hash1}${hash2}${asset2}${meta_hash2}" | openssl dgst -sha256
	expected := "1a92152f49e137e4fd15d9e04c53a27e9770a3b6c5268fdf86721a199ab59a41"

	hash, err := CalculateHash(manifestFile, projectRoot)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, hash, expected, "Expected %s, but returns %s", expected, hash)
}
