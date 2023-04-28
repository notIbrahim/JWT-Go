package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"jwt-go/pkg/helper"
)

type Users struct {
	Model
	Fullname string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password must be 6 characters or more"`
	Level    string    `gorm:"not null" json:"level" form:"level" valid:"required~Level is required"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (ReferenceUser *Users) BeforeCreate(ReferenceDB *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(ReferenceUser)

	if err != nil {
		err = errCreate
		return
	}

	ReferenceUser.Password = helper.LookupPassword(ReferenceUser.Password)
	return nil
}
