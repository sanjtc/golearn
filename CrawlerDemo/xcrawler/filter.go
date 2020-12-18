package xcrawler

type HTMLFilter func(HTMLElement) bool

type RequestFilter func(Request) bool

type ResponseFilter func(Response) bool

func FilterHTML(element HTMLElement, filters ...HTMLFilter) bool {
	for _, filter := range filters {
		if !filter(element) {
			return false
		}
	}

	return true
}

func FilterRequest(req Request, filters ...RequestFilter) bool {
	for _, filter := range filters {
		if !filter(req) {
			return false
		}
	}

	return true
}

func FilterResponse(resp Response, filters ...ResponseFilter) bool {
	for _, filter := range filters {
		if !filter(resp) {
			return false
		}
	}

	return true
}
