package models

type HttpHandlerFunc func(request Request, response *Response)
type HttpHandler struct {
	Method string
	Func   HttpHandlerFunc
}

type HttpUrlPath string
type RouteMapperType map[HttpUrlPath][]HttpHandler

type Request struct {
	Method  string
	UrlPath string
	Body    string
}

type ResponseStatusText map[int]string

type Response struct {
	StatusCode int
	Body       interface{}
}

type RespResponse struct {
	Response string
}
