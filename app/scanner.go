package app

import (
	"os"
	"path"
	"slices"
	"strings"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/types"
)

func IsSkipped(name string) bool {
	skiplist := []string{
		"404.md", "403.md", "500.md",
		"README.md", "CHANGELOG.md",
	}

	return slices.Contains(skiplist, name)
}

func (a *App) PreparePagesList() ([]*types.Page, error) {
	paths, err := ScanAllFilepaths(a.Root)

	if err != nil {
		return []*types.Page{}, err
	}

	pages := make([]*types.Page, 0, len(paths))

	for _, path := range paths {
		p, err := ReadPageFile(a.Root, path, a.Cfg)

		if err != nil {
			return []*types.Page{}, err
		}

		pages = append(pages, p)
	}

	return pages, nil
}

func FilterUtilityPages(pages []types.Page) []types.Page {
	filtered := make([]types.Page, 0, len(pages))
	for _, v := range pages {
		s := strings.Split(v.Filepath, "/")

		if IsSkipped(s[len(s)-1]) {
			continue
		}
		filtered = append(filtered, v)
	}

	return filtered
}

func ReadPageFile(root, path string, cfg *config.Config) (*types.Page, error) {
	f, err := os.ReadFile(path)

	if err != nil {
		return &types.Page{}, err
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
		pth := path.Join(root, entry.Name())
		if entry.IsDir() {
			var p, err = ScanAllFilepaths(pth)

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

		paths = append(paths, pth)
	}

	return paths, nil
}
