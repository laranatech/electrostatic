package pages

import (
	"os"
	"path"
	"strings"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/mdparcer"
)

func ReadTemplateFile(filepath string) (string, error) {
	f, err := os.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(f), nil
}

func FormatTemplate(tmplt string, f Page, cfg *config.Config) string {
	text := tmplt

	for key, value := range f.Meta {
		text = strings.Replace(text, "%"+key+"%", value, 1)
	}

	html := mdparcer.MdToHTML(f.Content, cfg)
	text = strings.Replace(text, "%CONTENT%", string(html), 1)

	return text
}

func FormatCardTemplate(
	tmplt string,
	page *Page,
	entry *config.CatalogEntry,
	cfg *config.Config,
) string {
	route := path.Join(entry.Directory, page.Route)
	tmplt = strings.Replace(tmplt, "%TITLE%", page.RawMeta["title"], 1)
	tmplt = strings.Replace(tmplt, "%LINK%", route, 1)
	tmplt = strings.Replace(tmplt, "%DESCRIPTION%", page.RawMeta["description"], 1)
	tmplt = strings.Replace(tmplt, "%DATE%", page.RawMeta["date"], 1)
	tmplt = strings.Replace(tmplt, "%KEYWORDS%", page.RawMeta["keywords"], 1)

	return tmplt
}
