package pkg

import (
	"io/ioutil"
	"errors"
	"gopkg.in/yaml.v2"
)

type manifestFile struct {
	filepath string
}

type manifestFormat struct {
	Assets []string `yaml:"Assets"`
}

func (manifestFile manifestFile) readAssets() (result []string, err error) {
	buf, err := ioutil.ReadFile(manifestFile.filepath)
	if err != nil {
		return
	}

	var manifest manifestFormat
	err = yaml.Unmarshal(buf, &manifest)
	if err != nil {
		return
	}
	if len(manifest.Assets) == 0 {
		err = errors.New("Assets not found")
		return
	}

	result = manifest.Assets
	return
}
