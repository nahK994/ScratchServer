package models

type RespHandlerFunc func(request RespRequest, response *RespResponse)

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
	Request string
}

type ResponseStatusText map[int]string

type HttpResponse struct {
	StatusCode int
	Body       interface{}
}

type RespResponse struct {
	Response string
}

type Response struct {
	Http HttpResponse
	Resp RespResponse
}
