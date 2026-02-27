package export

import (
	"log"
	"os"
	"strings"

	"larana.tech/go/electrostatic/pages"
)

func Export(root, dist string) error {
	err := os.Mkdir(dist, 0755)

	if err != nil {
		if os.IsExist(err) {
			os.RemoveAll(dist)
			return Export(root, dist)
		}
		return err
	}

	err = exportStatic(root, dist)

	if err != nil {
		return err
	}

	err = exportPages(root, dist)

	if err != nil {
		return err
	}

	return nil
}

func exportStatic(root, dist string) error {
	log.Println("Exporting static files...")

	log.Println("Destination directory: ", dist)

	publicFs := os.DirFS(root + "/public")
	err := os.CopyFS(dist+"/", publicFs)

	if err != nil {
		log.Println("Error while exporting `/public` directory:", err.Error())
	}

	err = exportPagesList(root, dist)

	if err != nil {
		return err
	}

	return nil
}

func exportPages(root, dist string) error {
	log.Println("Exporting pages...")

	paths, err := pages.ScanAllFilepaths(root)

	if err != nil {
		return err
	}

	tmp, err := pages.ReadTemplateFile(root)

	if err != nil {
		return err
	}

	for _, v := range paths {
		page, err := pages.ReadPageFile(root, v)

		if err != nil {
			return err
		}

		formatted := pages.FormatTemplate(tmp, page)

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

func exportPagesList(root, dist string) error {
	log.Println("Exporting pages list...")

	result, err := pages.FormatPageList(root)

	if err != nil {
		return err
	}

	os.Mkdir(dist+"/articles", 0755)

	return os.WriteFile(dist+"/articles/index.html", []byte(result), 0666)
}
