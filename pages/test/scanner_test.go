package pages_test

import (
	"testing"

	"github.com/laranatech/electrostatic/pages"
)

func TestScanAllFilepaths(t *testing.T) {
	root := "./root"

	expected := []string{
		root + "/dir/index.md",
		root + "/dir2/index.md",
		root + "/dir2/subdir/index.md",
		root + "/dir2/subdir/something.md",
		root + "/index.md",
		root + "/larana.md",
		root + "/no-meta-larana.md",
	}

	result, err := pages.ScanAllFilepaths(root)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if !matchPaths(result, expected) {
		t.Error("Result mismatch\n", expected, "\n!=\n", result)
	}
}

func matchPaths(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
