package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/naoina/toml"
)

const localConfigFile = ".finiki"

type localConfig struct {
	DataLocation string
}

var defaultLocal = localConfig{
	DataLocation: "sample",
}

func readLocalCfg() localConfig {
	var config = defaultLocal

	user, err := user.Current()
	if err == nil {
		cfgFile := filepath.Join(user.HomeDir, localConfigFile)

		f, err := os.Open(cfgFile)
		defer f.Close()

		if err == nil {
			dec := toml.NewDecoder(f)
			dec.Decode(&config)
		}
	}

	return config
}
