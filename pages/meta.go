package pages

import (
	"maps"
	"strings"

	"larana.tech/go/electrostatic/config"
)

func NewMetaMap(root string, params map[string]string, cfg *config.Config) (map[string]string, error) {
	meta := map[string]string{
		"title": strings.Replace(
			cfg.Meta.TitleTemplate,
			"%title%",
			cfg.Meta.FallbackTitle,
			1,
		),
		"description": strings.Replace(
			cfg.Meta.DescriptionTemplate,
			"%description%",
			cfg.Meta.FallbackDescription,
			1,
		),
		"keywords": strings.Replace(
			cfg.Meta.KeywordsTemplate,
			"%keywords%",
			cfg.Meta.FallbackKeywords,
			1,
		),
		"date": "",
	}

	maps.Copy(meta, params)

	for key, value := range params {
		switch key {
		case "title":
			meta["title"] = strings.Replace(
				cfg.Meta.TitleTemplate,
				"%title%",
				value,
				1,
			)
		case "description":
			meta["description"] = strings.Replace(
				cfg.Meta.DescriptionTemplate,
				"%description%",
				value,
				1,
			)
		case "keywords":
			meta["keywords"] = strings.Replace(
				cfg.Meta.KeywordsTemplate,
				"%keywords%",
				value,
				1,
			)
		}
	}

	return meta, nil
}
