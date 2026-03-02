package app

import "strings"

func FormatFilepathToRoute(root, pth string) string {
	pth = strings.Replace(pth, root, "", 1)
	pth = strings.Replace(pth, "/index.md", "", 1)
	pth = strings.Replace(pth, ".md", "", 1)

	if pth == "" {
		return "/"
	}

	if pth[0] != '/' {
		pth = "/" + pth
	}

	return pth
}
