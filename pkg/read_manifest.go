package pkg

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type ManifestFile struct {
	filepath string
}

type ManifestFormat struct {
	Assets []string `yaml:"Assets"`
}

func (manifestFile ManifestFile) ReadAssets() (result []string, err error) {
	buf, err := ioutil.ReadFile(manifestFile.filepath)
	if err != nil {
		return
	}

	var manifest ManifestFormat
	err = yaml.Unmarshal(buf, &manifest)
	if err != nil {
		return
	}

	result = manifest.Assets
	return
}
