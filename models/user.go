package models


import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint 			`gorm:"primarykey"`
	CreatedAt time.Time		`gorm:"autoCreateTime"`
	UpdatedAt time.Time		`gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Username         string
	Email           string
	Password         string
	List []List `gorm:"foreignKey:UserID"` 
	
}

type AddUser struct {
	Username         string
	Email           string
	Password         string
	
}

func (pengguna User) MapFromAddUser (user AddUser) User {
	var newDb User
	newDb.Username = user.Username
	newDb.Email = user.Email
	newDb.Password = user.Password
	return newDb 
}


type UserResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAT"`
	UpdatedAt time.Time      `json:"updateAT"`
	Username  string			`json:"username"`
	Email     string			`json:"Username"`
	Password  string			`json:"password"`

}

func (userResponse *UserResponse) MapFromDB (user User)  {
	userResponse.ID = user.ID
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Password = user.Password

}


type AuthResponse struct {
	ID        uint           `gorm:"primarykey"`
	Username  string
	Email     string
	Token  	string
}