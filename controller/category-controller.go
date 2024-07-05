package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webserver/dto/request"
	"webserver/dto/response"
	"webserver/service"
)

type WorkCategoryController struct {
	CategoryService service.IWorkCategoryService
}

func GetWorkCategoryControllerImpl(service service.IWorkCategoryService) *WorkCategoryController {
	return &WorkCategoryController{
		CategoryService: service,
	}
}
func (controller *WorkCategoryController) Create(ctx *gin.Context) {
	req := request.CreateWorkCategoryRequest{}
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("could not bind json")
		return // return error
	}

	controller.CategoryService.Create(req)

	webResponse := response.Response{
		Code:   200,
		Status: "category created",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
