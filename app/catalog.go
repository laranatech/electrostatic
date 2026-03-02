package app

import (
	"os"
	"path"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/types"
)

func (a *App) ScanCatalogPages(entry *config.CatalogEntry) ([]*types.Page, error) {
	pth := path.Join(a.Root, entry.Directory)

	return scan(a, pth)
}

func scan(a *App, pth string) ([]*types.Page, error) {
	pages := make([]*types.Page, 0, 100)

	files, err := os.ReadDir(pth)

	if err != nil {
		return pages, err
	}

	for _, f := range files {
		if f.Name()[0] == '.' {
			continue
		}
		fpth := path.Join(pth, f.Name())
		if f.IsDir() {
			dPages, err := scan(a, fpth)

			if err != nil {
				return pages, err
			}

			pages = append(pages, dPages...)
			continue
		}

		p, err := ReadPageFile(a.Root, fpth, a.Cfg)

		if err != nil {
			return pages, err
		}

		pages = append(pages, p)
	}

	return pages, nil
}
