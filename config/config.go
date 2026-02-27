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
	Path         string `toml:"path"`
	Title        string `tomt:"path"`
	Directory    string `toml:"directory"`
	CardTemplate string `toml:"card_template"`
}

type Catalogs struct {
	DefaultCardTemplate string         `toml:"default_card_template"`
	Entries             []CatalogEntry `toml:"entries"`
}

type Meta struct {
	TitleTemplate       string `toml:"title_template"`
	DescriptionTemplate string `toml:"description_template"`
	KeywordsTemplate    string `toml:"keywords_template"`
	FallbackTitle       string `toml:"fallback_title"`
	FallbackDescription string `toml:"fallback_description"`
	FallbackKeywords    string `toml:"fallback_keywords"`
}

type Config struct {
	DefaultTemplate string `toml:"default_template"`
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
