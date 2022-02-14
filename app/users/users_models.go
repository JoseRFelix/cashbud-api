package users

import (
	"github.com/JoseRFelix/cashbud-api/app/common"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex"`
}

func AutoMigrate() {
	common.DB.AutoMigrate(&UserModel{})
}

func GetAllUsers() ([]UserModel, error) {
	var models []UserModel
	err := common.DB.Find(&models).Error
	return models, err
}
