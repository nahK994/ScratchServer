package models

type HttpHandlerFunc func(request HttpRequest, response *HttpResponse)
type HttpHandler struct {
	Method string
	Func   HttpHandlerFunc
}

type HttpUrlPath string
type RouteMapperType map[HttpUrlPath]HttpHandler

type HttpRequest struct {
	Method  string
	UrlPath string
	Body    string
}

type RespRequest struct {
}

type ResponseStatusText map[int]string

type HttpResponse struct {
	StatusCode int
	Body       interface{}
}

type RespResponse struct {
}

type Response struct {
	Http HttpResponse
	Resp RespResponse
}
