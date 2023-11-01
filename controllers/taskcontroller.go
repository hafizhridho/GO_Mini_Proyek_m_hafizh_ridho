package controllers

import (
	"latihan/configs"
	"latihan/models"
	"latihan/models/base"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)
func CreateTugas(c echo.Context) error {
    
    var request models.AddTask
    if err := c.Bind(&request); err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "invalid request data",
			Data: nil,
		})
    }

   
    if request.TaskName == "" {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "berilah nama tugas",
			Data: nil,
		})
    }

    if request.Deskripsi == "" {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "berilah deskripsi",
			Data: nil,
		})
    }

    
   

    var existingList models.List
    if err := configs.DB.First(&existingList, request.ListID).Error; err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "id tidak ditemukan",
			Data: nil,
		})
    }
	var newDB models.Tugas
	newDB = newDB.MapFromAddTask(request)
	result := configs.DB.Create(&newDB)
	
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "error",
			Data: nil,
		})
	}
	
	
	if err := configs.DB.Preload("List").First(&newDB).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status: false,
			Message: "gagal memuat tugas",
			Data: nil,
		})
	}
	
	var response models.TaskResponse
	response.MapFROMDb(newDB)
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status: true,
		Message: "berhasil membuat tugas",
		Data: newDB,
	})
}


func GetTugas(c echo.Context) error {
    var tasks []models.Tugas
    if err := configs.DB.Preload("List").Find(&tasks).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "gagal mengambil daftar tugas",
            Data: nil,
        })
    }

  
    var  respon []models.TaskResponse
    for _, task := range tasks {
        var tasRespon models.TaskResponse
        tasRespon.MapFROMDb(task)
        respon = append(respon, tasRespon)
    }
    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "daftar tugas berhasil diambil",
        Data: tasks,
    })
}
func UpdateTask(c echo.Context) error {
    tugasID := c.Param("id")

    
    var update models.RequestTask
    if err := c.Bind(&update); err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "  data invalid",
            Data: nil,
        })
    }

   
    var existingTugas models.Tugas
    if err := configs.DB.First(&existingTugas, tugasID).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Tugas tidak ditemukan",
            Data: nil,
        })
    }

    
    if update.TaskName != "" {
        existingTugas.TaskName = update.TaskName
    }
    if update.Deskripsi != "" {
        existingTugas.Deskripsi = update.Deskripsi
    }
    if !update.Deadline.IsZero() {
        existingTugas.Deadline = update.Deadline
    }

    
    if err := configs.DB.Save(&existingTugas).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal mengupdate tugas",
            Data: nil,
        })
    }

    if err := configs.DB.First(&existingTugas, tugasID).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal memuat tugas yang telah diupdate",
            Data: nil,
        })
    }

    
    var response models.TaskResponse
    response.ID = existingTugas.ID
    response.TaskName = existingTugas.TaskName
    response.Deskripsi = existingTugas.Deskripsi
    response.Deadline = existingTugas.Deadline

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Berhasil mengupdate tugas",
        Data: response,
    })
}

func DeleteTugas(c echo.Context) error {
    
    tugasID := c.Param("id")

    var existingTugas models.Tugas
    if err := configs.DB.First(&existingTugas, tugasID).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Tugas tidak ditemukan",
            Data: nil,
        })
    }

   
    if err := configs.DB.Delete(&existingTugas).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal menghapus tugas",
            Data: nil,
        })
    }

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Tugas berhasil dihapus",
        Data: nil,
    })
}
func GetTaskById(c echo.Context) error {
    taskId := c.Param("id")

    var existTask models.Tugas
    if err := configs.DB.Preload("List").First(&existTask, taskId).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Tugas tidak ditemukan",
            Data: nil,
        })
    }

   

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Berhasil mendapatkan detail tugas",
        Data: existTask,
    })
}

func UpdateTugasStatus(c echo.Context) error {
    tugasID := c.Param("id")
    status := c.QueryParam("status")

    var existingTugas models.Tugas
    if err := configs.DB.First(&existingTugas, tugasID).Error; err != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Tugas tidak ditemukan",
            Data: nil,
        })
    }

    // Pastikan tugas memiliki informasi deadline
    if existingTugas.Deadline.IsZero() {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "Tugas tidak memiliki deadline",
            Data: nil,
        })
    }

    now := time.Now()

    if status == "selesai" {
        // Tugas selesai
        existingTugas.Status = true
        if now.Before(existingTugas.Deadline) {
            // Tugas dikerjakan tepat waktu
            return c.JSON(http.StatusOK, base.BaseResponse{
                Status: true,
                Message: "Anda mengerjakan tugas tepat waktu",
                Data: existingTugas,
            })
        } else {
            // Tugas terlambat
            return c.JSON(http.StatusOK, base.BaseResponse{
                Status: true,
                Message: "Anda terlambat menyelesaikan tugas",
                Data: existingTugas,
            })
        }
    } else if status == "belum-selesai" {
        
        existingTugas.Status = false
    } else {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "Status tidak valid",
            Data: nil,
        })
    }

    if err := configs.DB.Save(&existingTugas).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal mengupdate status tugas",
            Data: nil,
        })
    }

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Berhasil mengupdate status tugas",
        Data: existingTugas,
    })
}

func GetTasksByListID(c echo.Context) error {
    listID := c.Param("listID")
    
    var tasks []models.Tugas
    if err := configs.DB.Where("list_id = ?", listID).Find(&tasks).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status:  false,
            Message: "gagal",
            Data:    nil,
        })
    }
    
    return c.JSON(http.StatusOK, base.BaseResponse{
        Status:  true,
        Message: "berhasil",
        Data:    tasks,
    })
}
