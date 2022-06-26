package http

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Form struct {
	FileName string                  `form:"name" binding:"required"`
	FileType string                  `form:"type"  binding:"required"`
	Files    []*multipart.FileHeader `form:"files" binding:"required"`
}

func (h *Handler) FileHandlerInit(api *gin.RouterGroup) {
	api.GET("/file/:id", h.getFileById)
	api.POST("/file", h.uploadFile)
}

func (h *Handler) getFileById(ctx *gin.Context) {
	fileId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BadRequestResponse(&ResponseInput{
			Message: CheckInputData,
		}))
	}

	one, err := h.services.Files.FindOne(fileId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, NotFoundResponse(&ResponseInput{
			Message: FileNotFound, Data: err.Error(),
		}))
	}

	ctx.JSON(http.StatusOK, OkResponse(&ResponseInput{
		Message: FileFound, Data: one,
	}))
}

func (h *Handler) uploadFile(ctx *gin.Context) {
	var form Form

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, BadRequestResponse(&ResponseInput{
			Message: CheckInputData,
		}))
	}

	for _, formFile := range form.Files {
		openedFile, _ := formFile.Open()

		file, _ := ioutil.ReadAll(openedFile)
		saved, err := h.storage.SaveFile(h.storage.Directory, form.FileName, file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				InternalServerErrorResponse(&ResponseInput{
					Message: InternalServerError,
				}),
			)
		}

		ctx.JSON(http.StatusCreated,
			CreatedResponse(&ResponseInput{
				Message: FileUploaded,
				Data:    map[string]string{"filePath": saved},
			}),
		)
	}
}
