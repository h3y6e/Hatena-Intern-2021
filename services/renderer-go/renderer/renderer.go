package renderer

import (
	"bytes"
	"context"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	source := []byte(src)
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			mathjax.MathJax,
		),
	)
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return "", err
	}
	html := buf.String()
	return html, nil
}
