package xhttputil

import (
	"net/http"
	"strconv"
	"strings"
)

func ParseCookie(s string) *http.Cookie {
	cookie := &http.Cookie{}

	cookieArgs := strings.Split(s, ";")
	if len(cookieArgs) == 0 {
		return nil
	}

	for i, cookieArg := range cookieArgs {
		kv := strings.Split(cookieArg, "=")
		kvLen := len(kv)

		if kvLen == 0 {
			return nil
		}

		if i == 0 {
			cookie.Name = kv[0]
			if kvLen > 1 {
				cookie.Value = kv[1]
			}
		} else {
			switch kv[0] {
			case "max-age":
				{
					if kvLen == 1 {
						cookie.MaxAge = 0
						break
					}

					maxAge, err := strconv.Atoi(kv[1])
					if err != nil {
						cookie.MaxAge = 0
					} else {
						cookie.MaxAge = maxAge
					}
				}
			case "path":
				{
					if kvLen == 1 {
						cookie.Path = ""
						break
					}

					cookie.Path = kv[1]
				}
			case "HttpOnly":
				{
					cookie.HttpOnly = true
				}
			}
		}
	}

	return cookie
}
