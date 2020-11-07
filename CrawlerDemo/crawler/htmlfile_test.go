package crawler

import (
	"bytes"
	"os"
	"path"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	modulePath := pathutils.GetModulePath("CrawlerDemo")
	filePath := path.Join(modulePath, "crawler/temp/writetest")

	defer func() {
		err := os.RemoveAll(path.Dir(filePath))
		assert.Nil(t, err)
	}()

	htmlFile := HTMLFile{
		Path:    filePath,
		Content: []byte("hello"),
	}

	err := htmlFile.Write()
	assert.Nil(t, err)

	file, err := os.Open(filePath)
	assert.Nil(t, err)

	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	file.Close()
	assert.Nil(t, err)

	assert.Equal(t, buf.String(), string(htmlFile.Content))
}
