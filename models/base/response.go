package base



type BaseResponse struct {

	Status bool			`json:"status"`
	Message string		`json:"message"`
	Data interface{}	`json:"data"`
}


type UserResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}
