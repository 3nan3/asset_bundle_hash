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
	// echo -n "${hash1}${asset1}${hash2}${asset2}" | openssl dgst -sha256
	expected := "db61003eb2b1e97ef0162a33b7772d9a6eca7795c4c894770643f946a351961c"

	hash, err := CalculateHash(manifestFile, projectRoot)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, hash, expected, "Expected %s, but returns %s", expected, hash)
}
