package service

import (
	"github.com/go-playground/validator/v10"
	"webserver/dto/request"
	"webserver/model"
	"webserver/repository"
)

type IWorkCategoryService interface {
	Create(request request.CreateWorkCategoryRequest) error
}

type WorkCategoryService struct {
	Repo     repository.IWorkCategoryRepository
	Validate *validator.Validate
}

func GetCategoryServiceImp(categoryRepo repository.IWorkCategoryRepository, validate *validator.Validate) IWorkCategoryService {
	return &WorkCategoryService{
		Repo:     categoryRepo,
		Validate: validate,
	}
}

func (w WorkCategoryService) Create(request request.CreateWorkCategoryRequest) error {
	err := w.Validate.Struct(request)
	if err != nil {
		return err
	}
	w.Repo.Save(model.WorkCategory{
		Name: request.Name,
	})
	return nil
}
