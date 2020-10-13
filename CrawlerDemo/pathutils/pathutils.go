package pathutils

import (
	"os"
	"strings"
)

const ModuleName = "CrawlerDemo"

func GetModulePath() string {
	fp, _ := os.Getwd()
	fp = ConvertBackslashToSlash(fp)
	fp = strings.SplitAfter(fp, ModuleName)[0]

	return fp
}

func GetURLPath(url string) string {
	res := strings.Split(url, "://")
	return res[1]
}

func ConvertBackslashToSlash(s string) string {
	ss := strings.Split(s, "\\")
	res := ""

	for i, t := range ss {
		res += t

		if i < len(ss)-1 {
			res += "/"
		}
	}

	return res
}

func GetParentPath(p string) string {
	index := strings.LastIndex(p, "/")
	if index == -1 {
		return p
	}

	if index == len(p)-1 {
		p = p[0:index]
		index = strings.LastIndex(p, "/")
	}

	return p[0:index]
}

func CreateFile(fp string) (*os.File, error) {
	p := GetParentPath(fp)

	err := os.MkdirAll(p, 4)
	if err != nil {
		return nil, err
	}

	return os.Create(fp)
}
