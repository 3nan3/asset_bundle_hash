package pkg

import (
	"fmt"
	"os"
	"strings"
)

func CalculateHash(filepath string, projectRoot string) (err error) {
	manifestFile := manifestFile{filepath}
	assets, err := manifestFile.readAssets()
	if err != nil {
		return
	}

	hashBases := make([]string, len(assets), len(assets))
	for i, assetPath := range assets {
		file, err := os.Open(projectRoot + "/" + assetPath)
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
