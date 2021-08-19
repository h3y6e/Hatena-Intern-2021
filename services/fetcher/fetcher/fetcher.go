package fetcher

import (
	"context"
)

// Fetch は URL を受け取って title を抽出する
func Fetch(ctx context.Context, src string) (string, error) {
	return "title", nil
}
