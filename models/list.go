package models


import (
	"time"

	"gorm.io/gorm"
)

type List struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time		`gorm:"autoCreateTime"`
	UpdatedAt time.Time		`gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ListName string
	UserID	uint	`gorm:"index"`
	User  	User	`gorm:"foreignKey:UserID"`
	Tugas[]Tugas 	`gorm:"foreignKey:ListID"`

	

}

type AddList struct {
	UserID		uint
	Listname	string
}

func (model List) MapFromDb (list List) List {
	var newDB List
	newDB.UserID = list.UserID
	newDB.ListName = list.ListName
	return newDB
}
type ListResponse struct {
    ID        uint `json:"id"`
    CreatedAt time.Time     `json:"createdAT"`
    UpdatedAt time.Time     `json:"updatedAT"`
    ListName  string        `json:"listname"`
    User      User          `json:"user"`
    UserID    uint          `json:"userID"`
}

func (listResponse *ListResponse) MapFromList(list List) {
    listResponse.ID = list.ID
    listResponse.CreatedAt = list.CreatedAt
    listResponse.UpdatedAt = list.UpdatedAt
    listResponse.ListName = list.ListName
    listResponse.User = User{
        ID:       list.User.ID,
        Username: list.User.Username,
        Email:    list.User.Email,
    }

    listResponse.UserID = list.UserID
}
