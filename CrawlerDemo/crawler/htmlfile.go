package crawler

import "os"

type HTMLFile struct {
	Path    string
	Content []byte
}

func (f *HTMLFile) Write() error {
	file, err := os.Create(f.Path)
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
