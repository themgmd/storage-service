package http

import "net/http"

var (
	CheckInputData      = "Hoops! Something was wrong with input data. Check it and try again!"
	FileNotFound        = "Hoops! The file not in the storage"
	FileFound           = "The file is found in the storage"
	FileUploaded        = "The file is uploaded to the storage"
	FilesIdsFound       = "The files ids is found"
	InternalServerError = "Hoops! Something was wrong on server. Please try upload another file or write report"
)

type (
	BaseResponse struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
		Data       any    `json:"data"`
	}

	ResponseInput BaseResponse
)

func OkResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusOK,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func CreatedResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func BadRequestResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusBadRequest,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func NotFoundResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusNotFound,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func InternalServerErrorResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    input.Message,
		Data:       input.Data,
	}
}
