package model

import "gorm.io/gorm"

// add image url for it
// take data from form
type WorkCategory struct {
	gorm.Model
	Name string `json:"name"`
}
