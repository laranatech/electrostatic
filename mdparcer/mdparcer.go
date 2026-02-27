package mdparcer

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gomarkdown/markdown"
	mdHtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"larana.tech/go/electrostatic/config"
)

type CodeBlock struct {
	Code string
	Lang string
	Id   string
}

func MdToHTML(input []byte, cfg *config.Config) []byte {
	codeBlocks, md := ParseCodeBlocks(input)

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := mdHtml.CommonFlags | mdHtml.HrefTargetBlank | mdHtml.NofollowLinks | mdHtml.NoreferrerLinks | mdHtml.NoopenerLinks

	if cfg.Laziness.Images {
		htmlFlags = htmlFlags | mdHtml.LazyLoadImages
	}

	opts := mdHtml.RendererOptions{Flags: htmlFlags}
	renderer := mdHtml.NewRenderer(opts)

	h := markdown.Render(doc, renderer)

	if !cfg.Laziness.FirstImage {
		h = []byte(strings.Replace(string(h), "loading=\"lazy\"", "", 1))
	}

	return RenderCode(h, codeBlocks)
}

func ParseCodeBlocks(md []byte) ([]CodeBlock, []byte) {
	lines := strings.Split(string(md), "\n")

	newLines := make([]string, 0, len(lines))

	var blockLines []string

	blocks := make([]CodeBlock, 0, len(lines))

	lang := ""

	for i, line := range lines {
		if blockLines != nil {
			if line == "```" {
				id := fmt.Sprintf("CODE_BLOCK_%d", i)

				blocks = append(blocks, CodeBlock{
					Lang: lang,
					Id:   id,
					Code: strings.Join(blockLines, "\n"),
				})

				newLines = append(newLines, id)

				blockLines = nil
				continue
			}
			blockLines = append(blockLines, line)
			continue
		}

		if len(line) > 2 && line[0:3] == "```" {
			lang = strings.Replace(line, "```", "", 1)
			blockLines = make([]string, 0, len(lines))
			continue
		}

		newLines = append(newLines, line)
	}

	return blocks, []byte(strings.Join(newLines, "\n"))
}

func FormatCode(block CodeBlock) ([]byte, error) {
	lexer := lexers.Get(block.Lang)

	if lexer == nil {
		lexer = lexers.Fallback
	}

	lexer = chroma.Coalesce(lexer)

	style := styles.Get("monokai")

	iterator, err := lexer.Tokenise(nil, block.Code)

	if err != nil {
		return []byte{}, err
	}

	formatter := html.New(
		html.Standalone(false),
		html.WithClasses(true),
		html.WithLineNumbers(true),
	)

	var buf bytes.Buffer

	err = formatter.Format(&buf, style, iterator)

	return buf.Bytes(), err
}

func RenderCode(h []byte, blocks []CodeBlock) []byte {
	content := string(h)
	for _, v := range blocks {
		code, err := FormatCode(v)

		if err != nil {
			log.Println(err)
			continue
		}

		content = strings.Replace(content, v.Id, string(code), 1)
	}

	return []byte(content)
}
