package pathutils

import (
	"os"
	"path"
	"strings"
)

func GetModulePath() string {
	filePath, _ := os.Getwd()
	return path.Dir(filePath)
}

func GetURLPath(url string) string {
	res := strings.Split(url, "://")
	return res[1]
}
