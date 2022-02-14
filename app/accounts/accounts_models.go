package accounts

import (
	"github.com/JoseRFelix/cashbud-api/app/common"
	"github.com/JoseRFelix/cashbud-api/app/users"

	"gorm.io/gorm"
)

type AccountModel struct {
	gorm.Model
	Name    string          `json:"name"`
	Balance float32         `json:"balance" gorm:"default:0"`
	OwnerID int             `json:"owner_id"`
	Owner   users.UserModel `json:"owner"`
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

func CreateAccount(data CreateAccountDTO) error {
	model := AccountModel{Name: data.Name}
	err := common.DB.Create(&model).Error
	return err
}
