package repository

import (
	"gorm.io/gorm"
	"log"
	"webserver/model"
)

// make find all paginated
type IWorkCategoryRepository interface {
	Save(category model.WorkCategory) error
	//Update(workCategory model.WorkCategory)
	//Delete(workCategoryId int)
	//FindById(tagsId int) (categories model.WorkCategory, err error)
	//FindAll() []model.WorkCategory
}

type WorkCategoryRepository struct {
	Db *gorm.DB
}

func GetCategoryRepoImpl(Db *gorm.DB) IWorkCategoryRepository {
	return &WorkCategoryRepository{Db: Db}
}

func (t WorkCategoryRepository) Save(category model.WorkCategory) error {
	result := t.Db.Table("work_category").Create(&category)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}
