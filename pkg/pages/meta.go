package pages

import (
	"encoding/json"
	"maps"
	"os"
	"strings"
)

func ReadMetaConfig(root string) (MetaConfig, error) {
	f, err := os.ReadFile(root + "/meta.json")

	if err != nil {
		return MetaConfig{}, err
	}

	config := MetaConfig{}

	err = json.Unmarshal(f, &config)

	if err != nil {
		return MetaConfig{}, err
	}

	return config, nil
}

func NewMetaMap(root string, params map[string]string) (map[string]string, error) {
	config, err := ReadMetaConfig(root)

	if err != nil {
		return map[string]string{}, err
	}

	meta := map[string]string{
		"title":       strings.Replace(config.TitleTemplate, "%title%", config.FallbackTitle, 1),
		"description": strings.Replace(config.DescriptionTemplate, "%description%", config.FallbackDescription, 1),
		"keywords":    strings.Replace(config.KeywordsTemplate, "%keywords%", config.FallbackKeywords, 1),
		"date":        "",
	}

	maps.Copy(meta, params)

	for key, value := range params {
		switch key {
		case "title":
			meta["title"] = strings.Replace(config.TitleTemplate, "%title%", value, 1)
		case "description":
			meta["description"] = strings.Replace(config.DescriptionTemplate, "%description%", value, 1)
		case "keywords":
			meta["keywords"] = strings.Replace(config.KeywordsTemplate, "%keywords%", value, 1)
		}
	}

	return meta, nil
}
