package xcrawler

import "net/http"

type HTMLHandler func(HTMLElement)

type RequestHandler func(*http.Request)

type ResponseHandler func(*http.Response)
