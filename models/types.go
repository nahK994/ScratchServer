package models

type HandleHttpFunc func(request HttpRequest, response *HttpResponse)
type HandlerDetails struct {
	Method string
	Func   HandleHttpFunc
}

type HttpUrlPath string
type RouteMapperType map[HttpUrlPath]HandlerDetails

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
