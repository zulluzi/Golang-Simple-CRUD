package controller

import (
	"golang-crud-init/data/request"
	"golang-crud-init/data/response"
	"golang-crud-init/helper"
	"golang-crud-init/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// Create Controller
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindBodyWithJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	WebResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, WebResponse)
}

// Update Controller
func (controller *TagsController) Update(ctx *gin.Context) {
	UpdateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&UpdateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	UpdateTagsRequest.Id = id

	controller.tagsService.Update(UpdateTagsRequest)
	WebResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, WebResponse)
}

// Delete Controller
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)
	WebResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, WebResponse)

}

// FindById Controller
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagsResponse := controller.tagsService.FindById(id)
	WebResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagsResponse,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, WebResponse)
}

// FindAll Controller
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagsResponse := controller.tagsService.FindAll()
	WebResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagsResponse,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, WebResponse)
}
