package pages

import (
	"errors"
	"strings"
)

func ParsePageInfo(root string, f []byte) (Page, error) {
	parsedMeta := map[string]string{}

	text := string(f)

	lines := []string{}

	closed := false

	for i, v := range strings.Split(text, "\n") {
		if i == 0 && v != "---" {
			meta, err := NewMetaMap(root, nil)

			if err != nil {
				return Page{}, nil
			}

			return Page{
				Content: f,
				Meta:    meta,
			}, nil
		}

		if i == 0 {
			continue
		}

		if i != 0 && v == "---" {
			closed = true
			continue
		}

		if closed {
			lines = append(lines, v)
			continue
		}

		s := strings.Split(v, ":")

		if len(s) < 2 {
			return Page{}, errors.New("meta string should have key and value")
		}

		if len(s) == 2 {
			parsedMeta[s[0]] = strings.Trim(s[1], " ")
			continue
		}

		parsedMeta[s[0]] = strings.Trim(strings.Join(s[1:], ":"), " ")
	}

	meta, err := NewMetaMap(root, parsedMeta)

	if err != nil {
		return Page{}, err
	}

	return Page{
		Content: []byte(strings.Join(lines, "\n")),
		Meta:    meta,
	}, nil
}
