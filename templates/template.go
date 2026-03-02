package templates

import (
	"bytes"
	"html/template"
	"os"

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

func FormatPage(
	t *template.Template,
	p *types.Page,
	cfg *config.Config,
) (string, error) {
	content := p.Content

	content = mdparcer.MdToHTML(content, cfg)

	m := map[string]template.HTML{
		"content": template.HTML(content),
		"meta":    template.HTML(meta.Format(p.Meta, cfg)),
	}

	var buf bytes.Buffer

	err := t.Execute(&buf, m)

	return buf.String(), err
}
