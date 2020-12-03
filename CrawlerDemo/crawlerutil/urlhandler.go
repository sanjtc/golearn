package crawlerutil

type URLHandler func(url string)

func NewHandlerWithFilters(handler URLHandler, filters ...URLFilter) URLHandler {
	return func(url string) {
		if !FilterURL(url, filters...) {
			return
		}

		handler(url)
	}
}
