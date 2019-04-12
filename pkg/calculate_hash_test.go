package pkg

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculateHashSuccess(t *testing.T) {
	manifestFile := "testdata/sample.manifest"
	projectRoot := "../"

	// asset1="pkg/testdata/sample_1.txt"
	// asset2="pkg/testdata/sample_2.txt"
	// hash1=`cat $asset1 | openssl dgst -sha256`
	// hash2=`cat $asset2 | openssl dgst -sha256`
	// echo -n "${hash1}${asset1}${hash2}${asset2}" | openssl dgst -sha256
	expected := "43e81c963f4777b75b3fd09938a1d075a91904d1bd93e0d4fdaedbd870263b7f"

	hash, err := CalculateHash(manifestFile, projectRoot)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, hash, expected, "Expected %s, but returns %s", expected, hash)
}
