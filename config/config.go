package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type (
	Config struct {
		Test string
	}
)

func Get(filePath string) (cfg *Config, err error) {
	configs, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = errors.Wrapf(err, "[Config] fail to read file from %s", filePath)
		return
	}

	err = json.Unmarshal(configs, &cfg)
	if err != nil {
		err = errors.Wrapf(err, "[Config] fail to unmarshal file from %s", filePath)
		return
	}

	return
}
