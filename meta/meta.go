package meta

import (
	"errors"
	"fmt"
	"strings"

	"larana.tech/go/electrostatic/config"
)

type ParsingResult struct {
	Meta map[string]string
	Text string
}

func Parse(text string) (ParsingResult, error) {
	res := ParsingResult{
		Meta: map[string]string{},
		Text: text,
	}

	lines := make([]string, 0, 100)

	closed := false

	for i, v := range strings.Split(text, "\n") {
		if i == 0 {
			if v != "---" {
				return res, nil
			}
			continue
		}

		if v == "---" {
			closed = true
			continue
		}

		if closed {
			lines = append(lines, v)
			continue
		}

		s := strings.Split(v, ":")

		if len(s) < 2 {
			return res, errors.New("meta string should have key and value")
		}

		if len(s) == 2 {
			res.Meta[s[0]] = strings.Trim(s[1], " ")
			continue
		}

		res.Meta[s[0]] = strings.Trim(strings.Join(s[1:], ":"), " ")
	}

	res.Text = strings.Join(lines, "\n")

	return res, nil
}

func MapTags(tags []config.MetaTag) map[string]config.MetaTag {
	m := map[string]config.MetaTag{}

	for _, v := range tags {
		m[v.Key] = v
	}

	return m
}

func Format(tags map[string]string, cfg *config.Config) string {
	res := make([]string, 0, len(tags))

	cfgTags := MapTags(cfg.Meta.Tags)

	// TODO: sort, otherwise tests fail

	for k, v := range cfgTags {
		t, found := tags[k]

		var val string

		if !found {
			if v.Fallback == "" {
				continue
			}
			val = v.Fallback
		} else {
			val = t
		}

		r := fmt.Sprintf(v.Template, val)

		switch v.Tag {
		case "title":
			r = fmt.Sprintf("<title>%s</title>", r)
		default: // aka "meta"
			r = fmt.Sprintf(`<meta name="%s" content="%s" />`, v.Key, r)
		}

		res = append(res, r)
	}

	return strings.Join(res, "\n")
}
