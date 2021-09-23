package users

type UserData struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type UserResponses struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UserResponseData struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}
