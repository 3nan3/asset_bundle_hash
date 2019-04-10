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

	hashBases := make([]string, len(assets), len(assets))
	for i, assetPath := range assets {
		file, err := os.Open(assetPath)
		if err != nil {
			return err
		}
		defer file.Close()

		hash, err := sha256Sum(file)
		if err != nil {
			return err
		}
		hashBases[i] = hash + assetPath
	}

	assetBundleHash, err := sha256Sum(strings.NewReader(strings.Join(hashBases, "")))
	if err != nil {
		return
	}
	fmt.Println(assetBundleHash)
	return
}
