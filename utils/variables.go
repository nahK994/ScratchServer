package utils

import "github.com/nahK994/ScratchServer/models"

var RouteMapper map[string]func(models.Request) = make(map[string]func(models.Request))
