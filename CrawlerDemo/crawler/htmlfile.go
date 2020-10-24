package crawler

import (
	"os"

	"github.com/pantskun/commonutils/pathutils"
)

type HTMLFile struct {
	Path    string
	Content []byte
}

func (f *HTMLFile) Write() error {
	file, err := CreateFile(f.Path)
	if err != nil {
		return err
	}

	_, err = file.Write(f.Content)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

func CreateFile(fp string) (*os.File, error) {
	p := pathutils.GetParentPath(fp)

	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(fp)
}
