package templates

import (
	"path"
	"strings"

	"larana.tech/go/electrostatic/config"
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

	page := &types.Page{
		Content: []byte("%LINKS%"),
		Meta:    m,
	}

	var tmpltPath string

	if entry.CatalogTemplate != "" {
		tmpltPath = path.Join(root, entry.CatalogTemplate)
	} else {
		tmpltPath = path.Join(root, cfg.DefaultTemplate)
	}

	tmp, err := Read(tmpltPath)

	if err != nil {
		return "", err
	}

	result := FormatTemplate(tmp, page, cfg)

	result = strings.Replace(result, "%LINKS%", strings.Join(links, ""), 1)

	return result, nil
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

	tmplt, err := Read(tmpltPath)

	if err != nil {
		return nil, err
	}

	list := make([]string, 0, 100)

	for _, p := range pages {
		list = append(list, FormatCardTemplate(tmplt, p, entry, cfg))
	}

	return list, nil
}

func FormatCardTemplate(
	tmplt string,
	page *types.Page,
	entry *config.CatalogEntry,
	cfg *config.Config,
) string {
	route := path.Join(entry.Directory, page.Route)

	text := tmplt

	for k, v := range page.Meta {
		text = strings.Replace(text, "%"+k+"%", v, 1)
	}

	text = strings.Replace(text, "%link%", route, 1)

	return tmplt
}
