package controllers

import (
	"latihan/configs"
	"latihan/middleware"
	"latihan/models"
	"latihan/models/base"
	"latihan/repositories"

	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func LoginController(c echo.Context) error {
    var user models.User
    c.Bind(&user)
    authenticatedUser, err := repositories.Login(user.Password, user.Email)
    if err == gorm.ErrRecordNotFound {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status:  false,
            Message: "Email/Password not Found",
            Data:    nil,
        })
    } else if err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status:  false,
            Message: "Failed to Authenticate",
            Data:    nil,
        })
    } else {
        // Jika autentikasi berhasil, isi objek user dengan data pengguna yang sesuai
        user = authenticatedUser
    }

    tokenResult := middleware.GenerateJWT(user.ID, user.Username)

    var response models.UserAuthResponse
    response.ID = user.ID
    response.Email = user.Email
    response.Username = user.Username
    response.Token = tokenResult

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status:  true,
        Message: "Success Login",
        Data:    response,
    })
}
