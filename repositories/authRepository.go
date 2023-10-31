package repositories

import (
	"errors"
	"latihan/configs"
	"latihan/models"

	"gorm.io/gorm"
)
func Login(email string) (models.User, error) {
	var user models.User

	result := configs.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Pengguna tidak ditemukan, kembalikan kesalahan sesuai
			return models.User{}, gorm.ErrRecordNotFound
		}
		// Ada kesalahan lain, kembalikan kesalahan tersebut
		return models.User{}, result.Error
	}

	return user, nil
}
