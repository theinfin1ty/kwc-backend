package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *gin.H `json:"data"`
}

func SuccessResponse(data *gin.H) Response {
	var successResponse Response
	successResponse.Status = 200
	successResponse.Message = "Success"
	successResponse.Data = data
	return successResponse
}

func BadRequestResponse(err string) Response {
	var errorResponse Response
	errorResponse.Status = 400
	errorResponse.Message = err
	return errorResponse
}

func NotFoundResponse(err string) Response {
	var errorResponse Response
	errorResponse.Status = 404
	errorResponse.Message = err
	return errorResponse
}

func InternalServerErrorResponse(err error) Response {
	var errorResponse Response
	errorResponse.Status = 500
	errorResponse.Message = err.Error()
	return errorResponse
}

func UnauthorizedResponse(err error) Response {
	var errorResponse Response
	errorResponse.Status = 401
	errorResponse.Message = err.Error()
	return errorResponse
}
