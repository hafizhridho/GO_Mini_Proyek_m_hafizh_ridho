package controllers

import (

	"latihan/configs"
	"latihan/middleware"
	"latihan/models"
	"latihan/models/base"
	"latihan/repositories"

	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

    // Hash kata sandi menggunakan bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal hashing password",
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

    // Gunakan hashedPassword saat membuat pengguna baru
    var newDB models.User
    newDB = newDB.MapFromAddUser(user)
    newDB.Password = string(hashedPassword) // Simpan hashed password ke basis data

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

    
    authenticatedUser, err := repositories.Login(user.Email)
    if err == gorm.ErrRecordNotFound {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status:  false,
            Message: "Email tidak ditemukan",
            Data:    nil,
        })
    } else if err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status:  false,
            Message: "Gagal  autentikasi",
            Data:    nil,
        })
    }

    
    if err := bcrypt.CompareHashAndPassword([]byte(authenticatedUser.Password), []byte(user.Password)); err != nil {
        return c.JSON(http.StatusUnauthorized, base.BaseResponse{
            Status:  false,
            Message: "Password salah",
            Data:    nil,
        })
    }

    tokenResult := middleware.GenerateJWT(authenticatedUser.ID, authenticatedUser.Username)

    var response models.UserAuthResponse
    response.ID = authenticatedUser.ID
    response.Email = authenticatedUser.Email
    response.Username = authenticatedUser.Username
    response.Token = tokenResult

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status:  true,
        Message: "Berhasil Login",
        Data:    response,
    })
}
