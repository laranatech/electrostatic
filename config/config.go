package config

import (
	"log"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

type Laziness struct {
	Images     bool `toml:"lazy_images"`
	FirstImage bool `toml:"lazy_first_image"`
}

type CatalogEntry struct {
	Path            string `toml:"path"`
	Title           string `tomt:"path"`
	Directory       string `toml:"directory"`
	CatalogTemplate string `toml:"catalog_template"`
	CardTemplate    string `toml:"card_template"`
}

type Catalogs struct {
	DefaultCardTemplate    string         `toml:"default_card_template"`
	DefaultCatalogTemplate string         `toml:"default_catalog_template"`
	Entries                []CatalogEntry `toml:"entries"`
}

type MetaTag struct {
	Key      string `toml:"key"`
	Template string `toml:"template"`
	Fallback string `toml:"fallback"`
	Tag      string `toml:"tag"`
}

type Meta struct {
	Tags []MetaTag `toml:"tags"`
}

type Config struct {
	DefaultTemplate string `toml:"default_template"`
	Hotreload       bool   `toml:"hotreload"`
	Debug           bool   `toml:"debug"`
	Catalogs        Catalogs
	Meta            Meta
	Laziness        Laziness
}

func Read(root string) (*Config, error) {
	configPath := path.Join(root, "/config.toml")

	log.Println("Reading config: ", configPath)

	data, err := os.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
