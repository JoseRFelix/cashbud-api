package accounts

import (
	"github.com/JoseRFelix/cashbud-api/app/common"
	"gorm.io/gorm"
)

type AccountModel struct {
	gorm.Model
	ID      int     `json:"id" gorm:"primary_key"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
	Owner   string  `json:"owner"`
}

func AutoMigrate() {
	common.DB.AutoMigrate(&AccountModel{})
}

func GetAllAccounts() ([]AccountModel, error) {
	var models []AccountModel
	err := common.DB.Find(&models).Error
	return models, err
}

func GetOneAccount(condition interface{}) (AccountModel, error) {
	var model AccountModel
	err := common.DB.Where(condition).Find(&model).Error
	return model, err
}
