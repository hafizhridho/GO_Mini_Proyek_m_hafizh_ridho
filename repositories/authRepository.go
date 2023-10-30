package repositories

import (
	"latihan/configs"
	"latihan/models"
)

func Login(password string, email string) (models.User, error) {
	var user models.User

result := configs.DB.First(&user, "password = ? AND email = ?", password, email)

	if result.Error != nil {
		return models.User{}, result.Error

	}
	return user, nil
}