package service

import (
	"golang-crud-init/data/request"
	"golang-crud-init/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tags int) response.TagsResponse
	FindAll() []response.TagsResponse
}
