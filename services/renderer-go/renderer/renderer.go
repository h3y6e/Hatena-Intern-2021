package renderer

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	source := []byte(src)
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}
	html := buf.String()
	return html, nil
}
