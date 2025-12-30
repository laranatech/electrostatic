package pages_test

import (
	"os"
	"testing"

	"github.com/laranatech/electrostatic/pages"
)

func TestParsePageInfoMeta(t *testing.T) {
	root := "./root"

	f, err := os.ReadFile(root + "/larana.md")

	if err != nil {
		t.Error(err.Error())
		return
	}

	expected := map[string]string{
		"title":       "What is larana",
		"keywords":    "larana, gorana, framework",
		"description": "some description",
		"date":        "2025-01-01",
	}

	result, err := pages.ParsePageInfo(root, f)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if !matchMeta(expected, result.Meta) {
		t.Error("meta mismatch\n", expected, "\n!=\n", result.Meta)
	}
}

func matchMeta(a, b map[string]string) bool {
	for key := range a {
		if b[key] != a[key] {
			return false
		}
	}

	for key := range b {
		if b[key] != a[key] {
			return false
		}
	}

	return true
}

func TestParsePageInfoContent(t *testing.T) {
	root := "./root"

	f1, err := os.ReadFile(root + "/larana.md")

	if err != nil {
		t.Error(err.Error())
		return
	}

	f2, err := os.ReadFile(root + "/no-meta-larana.md")

	if err != nil {
		t.Error(err.Error())
		return
	}

	r1, _ := pages.ParsePageInfo(root, f1)
	r2, _ := pages.ParsePageInfo(root, f2)

	r1Content := string(r1.Content)
	r2Content := string(r2.Content)

	if r1Content != r2Content {
		t.Error("content mismatch\n", r1Content, "\n!=\n", r2Content)
	}
}
