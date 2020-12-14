package xcrawler

import "net/http"

type HTMLFilter func(HTMLElement) bool

type RequestFilter func(*http.Request) bool

type ResponseFilter func(*http.Response) bool

func FilterHTML(element HTMLElement, filters ...HTMLFilter) bool {
	for _, filter := range filters {
		if !filter(element) {
			return false
		}
	}

	return true
}

func FilterRequest(req *http.Request, filters ...RequestFilter) bool {
	for _, filter := range filters {
		if !filter(req) {
			return false
		}
	}

	return true
}

func FilterResponse(resp *http.Response, filters ...ResponseFilter) bool {
	for _, filter := range filters {
		if !filter(resp) {
			return false
		}
	}

	return true
}
