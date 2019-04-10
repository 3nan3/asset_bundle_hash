package pkg

import (
	"io"
	"crypto/sha256"
	"encoding/hex"
)

func sha256Sum(stream io.Reader) (result string, err error) {
	s := sha256.New()
	_, err = io.Copy(s, stream)
	if err != nil {
		return 
	}

	result = hex.EncodeToString(s.Sum(nil))
	return
}
