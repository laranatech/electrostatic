package meta_tests

import (
	"testing"

	"larana.tech/go/electrostatic/config"
	"larana.tech/go/electrostatic/meta"
)

var cfg = &config.Config{
	Meta: config.Meta{
		Tags: []config.MetaTag{
			{
				Key:      "title",
				Template: "%s | Sitename",
				Tag:      "title",
				Fallback: "No title",
			},
			{
				Key:      "keywords",
				Template: "%s, key 3, key 4",
				Tag:      "meta",
				Fallback: "No keywords",
			},
			{
				Key:      "description",
				Template: "%s",
				Tag:      "meta",
				Fallback: "No description",
			},
		},
	},
}

func TestFormatMetaTags(t *testing.T) {
	expected := `<title>title | Sitename</title>
<meta name="keywords" value="key 1, key 2, key 3, key 4" />
<meta name="description" value="test description" />`

	raw := map[string]string{
		"title":       "title",
		"keywords":    "key 1, key 2",
		"description": "test description",
	}

	result := meta.Format(raw, cfg)

	if expected == result {
		return
	}

	t.Errorf(
		"Failed to format meta. Expected:\n```%s```\nResult:\n```%s```",
		expected,
		result,
	)
}

func TestFormatMetaTagsFallbacks(t *testing.T) {
	expected := `<title>No title | Sitename</title>
<meta name="keywords" value="key 1, key 2, key 3, key 4" />
<meta name="description" value="No description" />`

	raw := map[string]string{
		"keywords": "key 1, key 2",
	}

	result := meta.Format(raw, cfg)

	if expected == result {
		return
	}

	t.Errorf(
		"Failed to format meta. Expected:\n```%s```\nResult:\n```%s```",
		expected,
		result,
	)
}

func TestParseMeta(t *testing.T) {
	text := `---
title: What is larana
keywords: larana, gorana, framework
description: some description
date: 2025-01-01
---

# larana

Some text
`

	expectedMeta := map[string]string{
		"title":       "What is larana",
		"keywords":    "larana, gorana, framework",
		"description": "some description",
		"date":        "2025-01-01",
	}

	expectedText := `
# larana

Some text
`

	result, err := meta.Parse(text)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if !matchMeta(expectedMeta, result.Meta) {
		t.Errorf(
			"Meta mismatch.\nExpected:\n```%s```\nResult:\n```%s```\n",
			expectedMeta,
			result.Meta,
		)
	}

	if expectedText != result.Text {
		t.Errorf(
			"Text mismatch.\nExpected:\n```%s```\nResult:\n```%s```\n",
			expectedText,
			result.Text,
		)
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
