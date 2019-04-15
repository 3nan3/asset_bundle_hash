package pkg

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestSha256Sum(t *testing.T) {
	testData := "test string"
	expected := "d5579c46dfcc7f18207013e65b44e4cb4e2c2298f4ac457ba8f82743f31e930b" // echo -n "test string" | openssl dgst -sha256

	result, err := sha256Sum(strings.NewReader(testData))
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, result, expected, "Expected %s, but returns %s", expected, result)
}
