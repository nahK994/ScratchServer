package errors

type UrlNotFound struct{}

func (e UrlNotFound) Error() string {
	return "URL not found"
}

type MethodNotAllowed struct{}

func (e MethodNotAllowed) Error() string {
	return "Method not allowed"
}
