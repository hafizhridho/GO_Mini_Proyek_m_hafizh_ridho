package controllers

import (
	"latihan/configs"
	"latihan/models"
	"latihan/models/base"

	"net/http"

	"github.com/labstack/echo/v4"
)


func Register(c echo.Context) error {
	var user models.AddUser
	c.Bind(&user)

	if user.Username == "" {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "Harap isi username",
			Data: nil,
		})
	}
	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "harap isi email",
			Data: nil,
		})
	}
	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "Harap isi Password",
			Data: nil,
		})
	}

	var existUser models.User
	hasil := configs.DB.Where("email = ?", user.Email).First(&existUser)
	if hasil.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "email telah dipakai",
			Data: nil,
		})
	}


	var newDB models.User
	newDB = newDB.MapFromAddUser(user)
	result := configs.DB.Create(&newDB)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "error",
			Data: nil,
		})
	}
	var response models.UserResponse
	response.MapFromDB(newDB)
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status: true,
		Message: "Berhasil membuat user",
		Data: newDB,
	})
}

func Get(c echo.Context) error {
    return c.JSON(http.StatusOK, nil)
}