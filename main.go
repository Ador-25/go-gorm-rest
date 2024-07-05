package main

import (
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
	"webserver/controller"
	"webserver/database"
	"webserver/repository"
	"webserver/router"
	"webserver/service"
)

func main() {
	db := database.GetDB()
	err := database.Migrate()
	if err != nil {
		log.Fatal("MIGRATION ERROR =>", err)
	}
	validate := validator.New()

	//Init Repository
	categoryRepo := repository.GetCategoryRepoImpl(db)

	//Init Service
	categoryService := service.GetCategoryServiceImp(categoryRepo, validate)

	//Init controller
	categoryController := controller.GetWorkCategoryControllerImpl(categoryService)
	//Router
	routes := router.NewRouter(categoryController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
