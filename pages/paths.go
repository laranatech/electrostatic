package pages

import "strings"

func FormatFilepathToRoute(root, path string) string {
	path = strings.Replace(path, root, "", 1)
	path = strings.Replace(path, "/index.md", "", 1)
	path = strings.Replace(path, ".md", "", 1)

	if path == "" {
		return "/"
	}

	return path
}
