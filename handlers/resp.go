package handlers

import "github.com/nahK994/TCPickle/models"

func HandleRespResponse(response *models.RespResponse) string {
	return response.Response
}
