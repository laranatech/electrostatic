package pages

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func IsSkipped(name string) bool {
	skiplist := []string{"404.md", "403.md", "500.md"}

	return slices.Contains(skiplist, name)
}

func PreparePagesList(root string) ([]Page, error) {
	paths, err := ScanAllFilepaths(root)

	if err != nil {
		return []Page{}, err
	}

	pages := make([]Page, 0, len(paths))

	for _, path := range paths {
		p, err := ReadPageFile(root, path)

		if err != nil {
			return []Page{}, err
		}

		pages = append(pages, p)
	}

	return pages, nil
}

func FilterUtilityPages(pages []Page) []Page {
	filtered := make([]Page, 0, len(pages))
	for _, v := range pages {
		s := strings.Split(v.Filepath, "/")

		if IsSkipped(s[len(s)-1]) {
			continue
		}
		filtered = append(filtered, v)
	}

	return filtered
}

func ReadPageFile(root, path string) (Page, error) {
	f, err := os.ReadFile(path)

	if err != nil {
		return Page{}, err
	}

	page, err := ParsePageInfo(root, f)

	if err != nil {
		return page, err
	}

	page.Filepath = path
	page.Route = FormatFilepathToRoute(root, path)

	return page, nil
}

func ScanAllFilepaths(root string) ([]string, error) {
	paths := []string{}

	entries, err := os.ReadDir(root)

	if err != nil {
		return []string{}, err
	}

	for _, entry := range entries {
		if entry.Name()[0] == '.' {
			continue
		}
		if entry.IsDir() {
			var p, err = ScanAllFilepaths(root + "/" + entry.Name())

			if err != nil {
				return []string{}, err
			}

			paths = append(paths, p...)
			continue
		}

		s := strings.Split(entry.Name(), ".")

		if s[len(s)-1] != "md" {
			continue
		}

		paths = append(paths, root+"/"+entry.Name())
	}

	return paths, nil
}

func FormatPageList(root string) (string, error) {
	pages, err := PreparePagesList(root)

	if err != nil {
		return "", err
	}

	links := []string{}

	for _, v := range FilterUtilityPages(pages) {
		links = append(links, fmt.Sprintf("<a href='%s'>%s</a>", v.Route, v.Meta["title"]))
	}

	meta, err := NewMetaMap(root, map[string]string{
		"title": "Список статей",
	})

	if err != nil {
		return "", err
	}

	f := Page{
		Content: []byte(strings.Join(links, "<br>")),
		Meta:    meta,
	}

	tmp, err := ReadTemplateFile(root)

	if err != nil {
		return "", err
	}

	result := FormatTemplate(tmp, f)

	return result, nil
}
