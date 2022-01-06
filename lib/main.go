package lib

import (
	"bytes"
	"fmt"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Renderer renders markdown to HTML.
type Renderer struct {
	md goldmark.Markdown
	p  *bluemonday.Policy
}

// NewRenderer returns a new Renderer.
func NewRenderer() *Renderer {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	p := bluemonday.UGCPolicy()
	return &Renderer{md: md, p: p}
}

// RenderAndSanitize renders the markdown to HTML from source and sanitizes it.
func (r *Renderer) RenderAndSanitize(source []byte) (sanitized []byte, err error) {
	unsanitized := new(bytes.Buffer)
	err = r.md.Convert(source, unsanitized)
	if err != nil {
		err = fmt.Errorf("md render: %w", err)
		return
	}
	return r.p.SanitizeReader(unsanitized).Bytes(), nil
}
