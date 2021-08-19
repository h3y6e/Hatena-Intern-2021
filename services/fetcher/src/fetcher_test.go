package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetch(t *testing.T) {
	link := "https://hatenacorp.jp/recruit/intern/2021"
	title, err := Fetch(context.Background(), link)
	assert.NoError(t, err)
	assert.Equal(t, "はてなリモートインターンシップ2021 - 株式会社はてな", title)
}
