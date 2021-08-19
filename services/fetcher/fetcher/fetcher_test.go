package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetch(t *testing.T) {
	link := "https://google.com/"
	title, err := Fetch(context.Background(), link)
	assert.NoError(t, err)
	assert.Equal(t, "title", title)
}
