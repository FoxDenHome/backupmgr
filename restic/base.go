package restic

import (
	"encoding/json"
	"os"
)

type Config struct {
	Args  map[string][]string `json:"args"`
	Repos map[string]*Repo    `json:"repos"`
}

type Repo struct {
	URI      string `json:"uri"`
	Password string `json:"password"`

	subArgs map[string][]string
}

func LoadConfig(path string) (*Config, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fh.Close()
	}()

	cfg := &Config{}
	decoder := json.NewDecoder(fh)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Args["default"] == nil {
		cfg.Args["default"] = []string{}
	}

	for _, repo := range cfg.Repos {
		repo.setup(cfg)
	}

	return cfg, nil
}
