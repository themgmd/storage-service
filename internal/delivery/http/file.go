package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileParams struct {
	Width  int    `json:"width" form:"width"`
	Height int    `json:"height" form:"height"`
	Type   string `json:"type" form:"type"`
}

func (h *Handler) FileHandlerInit(api *gin.RouterGroup) {
	api.GET("/file/:id", h.getFileById)
	api.POST("/file", h.uploadFile)
}

func (h *Handler) getFileById(ctx *gin.Context) {
	var fileParams FileParams

	fileId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BadRequestResponse(&ResponseInput{
			Message: CheckInputData,
		}))
		return
	}

	if err := ctx.ShouldBind(&fileParams); err != nil {
		ctx.JSON(http.StatusBadRequest, BadRequestResponse(&ResponseInput{
			Message: CheckInputData,
			Data:    map[string]string{"error": err.Error()},
		}))
		return
	}

	fmt.Println(fileParams, fileId)

	one, err := h.services.Files.FindOne(fileId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, NotFoundResponse(&ResponseInput{
			Message: FileNotFound, Data: err.Error(),
		}))
		return
	}

	ctx.JSON(http.StatusOK, OkResponse(&ResponseInput{
		Message: FileFound, Data: one,
	}))
}

// Upload Files on storage
func (h *Handler) uploadFile(ctx *gin.Context) {
	var id uint

	// parse mp form
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BadRequestResponse(&ResponseInput{
			Message: CheckInputData,
		}))
		return
	}

	// get file
	files := form.File["file"]

	for _, file := range files {
		// Load file to storage and save in db
		id, err = h.services.Files.UploadFile(file)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerErrorResponse(&ResponseInput{
				Message: InternalServerError,
				Data:    map[string]string{"error": err.Error()},
			}))
			return
		}
	}

	ctx.JSON(http.StatusCreated, CreatedResponse(&ResponseInput{
		Message: FileUploaded,
		Data:    map[string]uint{"id": id},
	}))
}
