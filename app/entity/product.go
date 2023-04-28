package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	Model
	Title       string `json:"title" form:"title" valid:"required~Product's title is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint
	User        *Users
}

func (ReferenceProduct *Product) BeforeCreate(ReferenceDB *gorm.DB) (err error) {
	_, ErrCreate := govalidator.ValidateStruct(ReferenceProduct)
	if ErrCreate != nil {
		err = ErrCreate
		return
	}
	return nil
}

func (ReferenceProduct *Product) BeforeUpdate(ReferenceDB *gorm.DB) (err error) {
	_, ErrUpdate := govalidator.ValidateStruct(ReferenceProduct)

	if ErrUpdate != nil {
		err = ErrUpdate
		return
	}

	return nil
}
