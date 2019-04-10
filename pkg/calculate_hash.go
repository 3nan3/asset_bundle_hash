package pkg

import (
	"fmt"
	"os"
	"strings"
)

func CalculateHash(filepath string) (err error) {
	manifestFile := ManifestFile{filepath}
	assets, err := manifestFile.ReadAssets()
	if err != nil {
		return
	}

	hashes := make([]string, len(assets), len(assets))
	for i, assetFile := range assets {
		file, err := os.Open(assetFile)
		if err != nil {
			return err
		}
		defer file.Close()

		hash, err := sha256Sum(file)
		if err != nil {
			return err
		}
		hashes[i] = hash
	}

	assetBundleHash, err := sha256Sum(strings.NewReader(strings.Join(hashes, "")))
	if err != nil {
		return
	}
	fmt.Println(assetBundleHash)
	return
}
