package templates

import (
	"os"
	"strings"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/mdparcer"
	"larana.tech/go/electrostatic/meta"
	"larana.tech/go/electrostatic/types"
)

// Read template file
func Read(filepath string) (string, error) {
	f, err := os.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(f), nil
}

func FormatTemplate(tmplt string, f *types.Page, cfg *config.Config) string {
	text := tmplt

	mt := meta.Format(f.Meta, cfg)

	text = strings.Replace(text, "%META%", mt, 1)

	html := mdparcer.MdToHTML(f.Content, cfg)
	text = strings.Replace(text, "%CONTENT%", string(html), 1)

	return text
}
