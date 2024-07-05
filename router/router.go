package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webserver/controller"
)

func NewRouter(category *controller.WorkCategoryController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	categoryRouter := router.Group("/category")
	categoryRouter.POST("", category.Create)

	//categoryRouter.GET("", tagController.FindAll)
	//categoryRouter.GET("/:tagId", tagController.FindById)
	//categoryRouter.PATCH("/:tagId", tagController.Update)
	//categoryRouter.DELETE("/:tagId", tagController.Delete)

	return service
}
