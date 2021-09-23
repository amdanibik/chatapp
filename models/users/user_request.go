package users

type UserRegister struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password []byte `form:"password"`
}

type UserChecker struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
