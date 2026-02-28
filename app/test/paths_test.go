package app_test

import (
	"testing"

	"larana.tech/go/electrostatic/app"
)

func TestFormatFilepathtoUrl(t *testing.T) {
	root := "./root"

	data := map[string]string{
		root + "/index.md":            "/",
		root + "/larana.md":           "/larana",
		root + "/faq/index.md":        "/faq",
		root + "/faq/subfile.md":      "/faq/subfile",
		root + "/faq/subdir/index.md": "/faq/subdir",
	}

	for filepath, expected := range data {
		formatted := app.FormatFilepathToRoute(root, filepath)

		if formatted != expected {
			t.Errorf("Wrong output for `%s`, expected `%s`, result is `%s` ", filepath, expected, formatted)
		}
	}
}
