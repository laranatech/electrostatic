package export

import (
	"log"
	"os"
	"path"
	"strings"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/pages"
)

func Export(root, dist string, cfg *config.Config) error {
	err := os.Mkdir(dist, 0755)

	if err != nil {
		if os.IsExist(err) {
			os.RemoveAll(dist)
			return Export(root, dist, cfg)
		}
		return err
	}

	err = exportStatic(root, dist, cfg)

	if err != nil {
		return err
	}

	err = exportPages(root, dist, cfg)

	if err != nil {
		return err
	}

	return nil
}

func exportStatic(root, dist string, cfg *config.Config) error {
	log.Println("Exporting static files...")

	log.Println("Destination directory: ", dist)

	publicFs := os.DirFS(root + "/public")
	err := os.CopyFS(dist+"/", publicFs)

	if err != nil {
		log.Println("Error while exporting `/public` directory:", err.Error())
	}

	err = exportPagesList(root, dist, cfg)

	if err != nil {
		return err
	}

	return nil
}

func exportPages(root, dist string, cfg *config.Config) error {
	log.Println("Exporting pages...")

	paths, err := pages.ScanAllFilepaths(root)

	if err != nil {
		return err
	}

	tmp, err := pages.ReadTemplateFile(path.Join(root, cfg.DefaultTemplate))

	if err != nil {
		return err
	}

	for _, v := range paths {
		page, err := pages.ReadPageFile(root, v, cfg)

		if err != nil {
			return err
		}

		formatted := pages.FormatTemplate(tmp, page, cfg)

		newPath := strings.Replace(page.Filepath, root, dist, 1)
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

func exportPagesList(root, dist string, cfg *config.Config) error {
	log.Println("Exporting pages list...")

	for _, entry := range cfg.Catalogs.Entries {
		result, err := pages.FormatPageList(root, &entry, cfg)

		if err != nil {
			return err
		}

		distDir := path.Join(dist, entry.Path)

		os.Mkdir(distDir, 0755)

		distFile := path.Join(distDir, "/index.html")

		err = os.WriteFile(distFile, []byte(result), 0666)

		if err != nil {
			return err
		}
	}

	return nil
}
