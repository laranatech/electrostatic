package content

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed content-template/*
var ContentTemplateFs embed.FS

func InitializeTemplate(root string) {
	var f, _ = fs.Sub(ContentTemplateFs, "content-template")
	os.CopyFS(root, f)
}
