package fetcher

import (
	"context"
)

// Fetch は Link を受け取って Title を抽出する
func Fetch(ctx context.Context, src string) (string, error) {
	return "title", nil
}
