package models

var UrlMapper map[string]func(request *Request)

type Request struct {
	Method  string
	UrlPath string
	Body    string
}
