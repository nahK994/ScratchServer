package models

type HandleHttpFunc func(r Request)
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
