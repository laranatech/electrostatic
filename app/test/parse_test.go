package app_test

import (
	"os"
	"testing"

	"larana.tech/go/electrostatic/app"
	"larana.tech/go/electrostatic/config"
)

// TODO: fix tests
func TestParsePageInfoContent(t *testing.T) {
	root := "./root"

	f1, err := os.ReadFile(root + "/larana.md")

	if err != nil {
		t.Error(err.Error())
		return
	}

	f2, err := os.ReadFile(root + "/no-meta-larana.md")

	cfg := &config.Config{}

	if err != nil {
		t.Error(err.Error())
		return
	}

	r1, _ := app.ParsePageInfo(root, f1, cfg)
	r2, _ := app.ParsePageInfo(root, f2, cfg)

	r1Content := string(r1.Content)
	r2Content := string(r2.Content)

	if r1Content != r2Content {
		t.Error("content mismatch\n", r1Content, "\n!=\n", r2Content)
	}
}
