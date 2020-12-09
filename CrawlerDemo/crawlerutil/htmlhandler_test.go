package crawlerutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyForCrawler(t *testing.T) {
	url := "test"
	expected := "/crawler/test"
	got := generateKeyForCrawler(url)
	assert.Equal(t, expected, got)
}
