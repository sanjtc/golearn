package htmlutil

import (
	"os"

	"github.com/pantskun/commonutils/pathutils"
)

type HTMLFile struct {
	Path    string
	Content []byte
}

// writeToFile
// 将content写入filePath文件中
func writeToFile(filePath string, content []byte) error {
	file, err := createFile(filePath)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// createFile
// 创建fp文件
func createFile(fp string) (*os.File, error) {
	p := pathutils.GetParentPath(fp)

	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(fp)
}
