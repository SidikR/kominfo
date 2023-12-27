package model

import (
	"fmt"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	Name       string `validate:"required" gorm:"type:varchar(50)"`
	Deskripsi    string `validate:"required" gorm:"type:varchar(50)"`
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *Program) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	return nil
}
