package finiki

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/naoina/toml"
)

const localConfigFile = ".finiki"
const siteConfigFile = "site"

type localConfig struct {
	DataLocation string
}

type SiteConfig struct {
	RecentPages []string
}

var defaultLocal = localConfig{
	DataLocation: "sample",
}

func ReadLocalCfg() localConfig {
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

func ReadSiteCfg(dataLocation string) SiteConfig {
	var cfg SiteConfig

	f, err := os.Open(filepath.Join(dataLocation, siteConfigFile))
	if err == nil {
		dec := json.NewDecoder(f)
		dec.Decode(&cfg)
	}

	return cfg
}
