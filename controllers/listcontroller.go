package controllers

import (
	"latihan/configs"
	"latihan/models"
	"latihan/models/base"
	

	"net/http"

	"github.com/labstack/echo/v4"
)
func CreateList(c echo.Context) error {
	var newList models.List
	c.Bind(&newList)

	if newList.ListName == "" {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "Harapi di isi ",
			Data: nil,
		})
	}

	
var existingUser models.User
if err := configs.DB.First(&existingUser, newList.UserID).Error; err != nil {
    return c.JSON(http.StatusBadRequest, base.BaseResponse{
		Status: false,
		Message:  "User dengan UserID yang diberikan tidak ditemukan",
		Data: nil,
	})
}


newList.User = existingUser 
if err := configs.DB.Create(&newList).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, base.BaseResponse{
		Status: false,
		Message: "Gagal membuat daftar",
		Data: nil,
	})
}


	return c.JSON(http.StatusCreated, base.BaseResponse{
        Status: true,
        Message: "Berhasil",
        Data: newList,
    })
}
func GetAllLists(c echo.Context) error {
    var lists []models.List

   
    if err := configs.DB.Preload("User").Find(&lists).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal mendapatkan daftar",
            Data: nil,
        })
    }

 
    var response []models.ListResponse
    for _, list := range lists {
        var listResponse models.ListResponse
        listResponse.MapFromList(list)
        response = append(response, listResponse)
    }

    return c.JSON(http.StatusOK, response)
}
func UpdateList(c echo.Context) error {
    listID := c.Param("id") 
    var updatedList models.List
    if err := c.Bind(&updatedList); err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "Invalid request data",
            Data: nil,
        })
    }

    var existingList models.List
    if err := configs.DB.Preload("User").First(&existingList, listID).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "List not found",
            Data: nil,
        })
    }

   
    existingList.ListName = updatedList.ListName

    if err := configs.DB.Save(&existingList).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Failed to update the list",
            Data: nil,
        })
    }

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "List updated successfully",
        Data: existingList,
    })
}
func DeleteList(c echo.Context) error {
    listID := c.Param("id") 

    var existingList models.List
    if err := configs.DB.First(&existingList, listID).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "List not found",
            Data: nil,
        })
    }

   
    if err := configs.DB.Delete(&existingList).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal menghapus",
            Data: nil,
        })
    }

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Berhasil Menghapus",
        Data: nil, 
    })
}
func GetListByID(c echo.Context) error {
    listID := c.Param("id")

    var existingList models.List
    if err := configs.DB.Preload("User").First(&existingList, listID).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Daftar tidak ditemukan",
            Data: nil,
        })
    }

    // Anda dapat menggabungkan data daftar dengan tugas jika diperlukan

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Berhasil mendapatkan detail daftar",
        Data: existingList,
    })
}


func GetListByUserID(c echo.Context) error {
    userID := c.Param("userID")
    
    var daftar []models.List
    if err := configs.DB.Where("user_id = ?", userID).Find(&daftar).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status:  false,
            Message: "gagal mencari daftar",
            Data:    nil,
        })
    }
    
    return c.JSON(http.StatusOK, base.BaseResponse{
        Status:  true,
        Message: "Success",
        Data:    daftar,
    })
}