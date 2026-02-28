package app

import (
	"log"
	"os"
	"path"
	"strings"

	"larana.tech/go/electrostatic/templates"
	"larana.tech/go/electrostatic/types"
)

func (a *App) Export() error {
	err := os.Mkdir(a.Dist, 0755)

	if err != nil {
		if os.IsExist(err) {
			os.RemoveAll(a.Dist)
			return a.Export()
		}
		return err
	}

	err = a.ExportStatic()

	if err != nil {
		return err
	}

	err = a.ExportPages()

	if err != nil {
		return err
	}

	return nil
}

func (a *App) ExportStatic() error {
	log.Println("Exporting static files...")

	log.Println("Destination directory: ", a.Dist)

	publicFs := os.DirFS(path.Join(a.Root, "/public"))
	err := os.CopyFS(a.Dist+"/", publicFs)

	if err != nil {
		log.Println("Error while exporting `/public` directory:", err.Error())
	}

	err = a.ExportPagesList()

	if err != nil {
		return err
	}

	return nil
}

func (a *App) ExportPages() error {
	log.Println("Exporting pages...")

	paths, err := ScanAllFilepaths(a.Root)

	if err != nil {
		return err
	}

	tmp, err := templates.Read(path.Join(a.Root, a.Cfg.DefaultTemplate))

	if err != nil {
		return err
	}

	for _, v := range paths {
		page, err := ReadPageFile(a.Root, v, a.Cfg)

		if err != nil {
			return err
		}

		formatted := templates.FormatTemplate(tmp, page, a.Cfg)

		newPath := strings.Replace(page.Filepath, a.Root, a.Dist, 1)
		newPath = strings.Replace(newPath, ".md", ".html", 1)

		s := strings.Split(newPath, "/")

		dirPath := strings.Join(s[0:len(s)-1], "/")

		os.MkdirAll(dirPath, 0777)

		err = os.WriteFile(newPath, []byte(formatted), 0666)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) ExportPagesList() error {
	log.Println("Exporting pages list...")

	for _, entry := range a.Cfg.Catalogs.Entries {
		// TODO
		pages := make([]*types.Page, 0, 100)

		result, err := templates.FormatCatalogPage(a.Root, &entry, pages, a.Cfg)

		if err != nil {
			return err
		}

		distDir := path.Join(a.Dist, entry.Path)

		os.Mkdir(distDir, 0755)

		distFile := path.Join(distDir, "/index.html")

		err = os.WriteFile(distFile, []byte(result), 0666)

		if err != nil {
			return err
		}
	}

	return nil
}
