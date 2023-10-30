package models

import (
	"time"

	"gorm.io/gorm"
)

type Tugas struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TaskName  string
	Deskripsi string
	Deadline  time.Time
	Status    bool
	ListID    uint `gorm:"index"`
	List      List `gorm:"foreignKey:ListID"`



}

type AddTask struct {
    TaskName  string `json:"taskName"`
	Deskripsi string `json:"deskripsi"`
	Deadline  time.Time `json:"deadline"`
	Status    bool   `json:"status"`
	ListID    uint   `json:"listID"`
}

func (tugas Tugas) MapFromAddTask (task AddTask) Tugas {
	var newDb Tugas
	newDb.TaskName = task.TaskName
	newDb.Deskripsi = task.Deskripsi
	newDb.Deadline = task.Deadline
	newDb.Status = task.Status
	newDb.ListID = task.ListID
	return newDb
}

type TaskResponse struct {
	ID        uint      `json:"id"`
	TaskName  string `json:"taskName"`
	Deskripsi string `json:"deskripsi"`
	Deadline  time.Time `json:"deadline"`
	Status    bool   `json:"status"`
	List	List		`json:"list"`
	ListID    uint   `json:"listID"`
	
}

func (taskResponse *TaskResponse) MapFROMDb (task Tugas){
	taskResponse.TaskName = task.TaskName
	taskResponse.Deskripsi = task.Deskripsi
	taskResponse.Deadline = task.Deadline
	taskResponse.Status = task.Status
	taskResponse.ListID = task.ListID
	taskResponse.List = task.List
	taskResponse.List = List{
        ID:      task.List.ID,
		ListName: task.List.ListName,
        
    }
	
	
}


type RequestTask struct {
	TaskName  string    `json:"taskName"`
	Deskripsi string    `json:"deskripsi"`
	Deadline  time.Time `json:"deadline"`
}                