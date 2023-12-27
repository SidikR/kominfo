package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type Stunting struct {
	template.UUIDModel
	NIK       string `validate:"required" gorm:"type:varchar(50)"`
	Name      string `validate:"required" gorm:"type:varchar(50)"`
	Kecamatan string `validate:"required" gorm:"type:varchar(50)"`
	Desa      string `validate:"required" gorm:"type:varchar(50)"`
	Koodinat  string `json:"koordinat" gorm:"type:varchar(50)"`
	Status    string `validate:"required" gorm:"type:varchar(50)"`
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *Stunting) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}
	return nil
}
