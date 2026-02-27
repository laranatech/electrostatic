package config

import (
	"log"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

type CatalogEntry struct {
	Path      string `toml:"path"`
	Title     string `tomt:"path"`
	Directory string `toml:"directory"`
}

type Config struct {
	Catalogs struct {
		Entries []CatalogEntry `toml:"entries"`
	}
	Meta struct {
		TitleTemplate       string `toml:"title_template"`
		DescriptionTemplate string `toml:"description_template"`
		KeywordsTemplate    string `toml:"keywords_template"`
		FallbackTitle       string `toml:"fallback_title"`
		FallbackDescription string `toml:"fallback_description"`
		FallbackKeywords    string `toml:"fallback_keywords"`
	}
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
