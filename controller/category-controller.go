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

	// Attempt to bind the JSON request body to the struct
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("Could not bind JSON: %v\n", err) // Log the error for debugging purposes
		webResponse := response.Response{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(), // Include the error message in the response
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Proceed with business logic
	err = controller.CategoryService.Create(req)
	if err != nil {
		log.Printf("Failed to create category: %v\n", err) // Log the error for debugging purposes
		webResponse := response.Response{
			Code:   500, // Internal Server Error
			Status: "Could not create",
			Data:   nil,
		}
		ctx.JSON(http.StatusInternalServerError, webResponse) // Return a 500 status code for internal server errors
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: "Category Created",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
