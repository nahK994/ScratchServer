package models

type HandleHttpFunc func(request Request, response *Response)
type HandlerDetails struct {
	Method string
	Func   HandleHttpFunc
}

type HttpUrlPath string
type RouteMapperType map[HttpUrlPath]HandlerDetails

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
