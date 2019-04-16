package pkg

import (
	"os"
	"strings"
)

func CalculateHash(filepath string, projectRoot string) (assetBundleHash string, err error) {
	manifestFile := manifestFile{filepath}
	assets, err := manifestFile.readAssets()
	if err != nil {
		return
	}

	hashBases := make([]string, len(assets), len(assets))
	for i, assetPath := range assets {
		hash, err := fileHash(projectRoot + "/" + assetPath)
		if err != nil {
			return "", err
		}

		metaHash, err := fileHash(projectRoot + "/" + assetPath + ".meta")
		if err != nil {
			return "", err
		}

		hashBases[i] = hash + assetPath + metaHash
	}

	result, err := sha256Sum(strings.NewReader(strings.Join(hashBases, "")))
	if err != nil {
		return
	}

	assetBundleHash = result
	return
}

func fileHash(filepath string) (hash string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	hash, err = sha256Sum(file)
	return
}
