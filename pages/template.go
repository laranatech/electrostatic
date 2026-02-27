package pages

import (
	"os"
	"strings"

	"larana.tech/go/electrostatic/mdparcer"
)

func ReadTemplateFile(root string) (string, error) {
	f, err := os.ReadFile(root + "/template.html")

	if err != nil {
		return "", err
	}

	return string(f), nil
}

func FormatTemplate(tmplt string, f Page) string {
	text := tmplt

	for key, value := range f.Meta {
		text = strings.Replace(text, "%"+key+"%", value, 1)
	}

	html := mdparcer.MdToHTML(f.Content)
	text = strings.Replace(text, "%CONTENT%", string(html), 1)

	return text
}
