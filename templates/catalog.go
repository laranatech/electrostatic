package templates

import (
	"bytes"
	"html/template"
	"log"
	"maps"
	"path"
	"strings"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/meta"
	"larana.tech/go/electrostatic/types"
)

func FormatCatalogPage(
	root string,
	entry *config.CatalogEntry,
	pages []*types.Page,
	cfg *config.Config,
) (string, error) {
	links, err := FormatCatalogList(root, entry, pages, cfg)

	if err != nil {
		return "", err
	}

	m := map[string]string{
		"title":       entry.Title,
		"description": entry.Title,
		"keywords":    entry.Title,
	}

	var tmpltPath string

	if entry.CatalogTemplate != "" {
		tmpltPath = path.Join(root, entry.CatalogTemplate)
	} else {
		tmpltPath = path.Join(root, cfg.DefaultTemplate)
	}

	f, err := Read(tmpltPath)

	if err != nil {
		return "", err
	}

	t := template.Must(template.New("catalog").Parse(f))

	var result bytes.Buffer

	err = t.Execute(&result, map[string]template.HTML{
		"content": template.HTML(strings.Join(links, "")),
		"meta":    template.HTML(meta.Format(m, cfg)),
	})

	return result.String(), err
}

func FormatCatalogList(
	root string,
	entry *config.CatalogEntry,
	pages []*types.Page,
	cfg *config.Config,
) ([]string, error) {
	var tmpltPath string

	if entry.CardTemplate != "" {
		tmpltPath = path.Join(root, entry.CardTemplate)
	} else {
		tmpltPath = path.Join(root, cfg.Catalogs.DefaultCardTemplate)
	}

	f, err := Read(tmpltPath)

	if err != nil {
		return nil, err
	}

	t, err := template.New("card").Parse(f)

	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 100)

	for _, p := range pages {
		l, err := FormatCardTemplate(t, p, cfg)

		if err != nil {
			return nil, err
		}

		list = append(list, l)
	}

	return list, nil
}

func FormatCardTemplate(
	t *template.Template,
	page *types.Page,
	cfg *config.Config,
) (string, error) {
	m := map[string]string{
		"route": page.Route,
	}

	if cfg.Debug {
		log.Println("[DEBUG] ", page.Route)
	}

	maps.Copy(m, page.Meta)

	var buf bytes.Buffer

	err := t.Execute(&buf, m)

	return buf.String(), err
}
