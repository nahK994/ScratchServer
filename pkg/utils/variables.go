package utils

import (
	"github.com/nahK994/SimpleServer/pkg/types"
)

var HttpRouteMapper types.RouteMapperType = make(types.RouteMapperType)

var StatusText = types.ResponseStatusText{
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",

	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",

	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
}
