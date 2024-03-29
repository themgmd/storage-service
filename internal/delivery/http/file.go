package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/onemgvv/storage-service/internal/domain"
)

func (h *Handler) FileHandlerInit(api *gin.RouterGroup) {
	api.GET("/file/:id", h.getFileById)
	api.GET("/files", h.getFiles)
	api.POST("/file", h.uploadFile)
}

func (h *Handler) getFileById(ctx *gin.Context) {
	var fileParams domain.FileParams

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

	one, err := h.services.Files.GetFile(fileId, fileParams)
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

func (h *Handler) getFiles(ctx *gin.Context) {
	var params domain.GetFilesParams

	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, BadRequestResponse(&ResponseInput{
			Message: CheckInputData,
			Data:    map[string]string{"error": err.Error()},
		}))
		return
	}

	ids := h.services.Files.AllFiles(params.Type)

	ctx.JSON(http.StatusOK, OkResponse(&ResponseInput{
		Message: FilesIdsFound,
		Data:    ids,
	}))
}
