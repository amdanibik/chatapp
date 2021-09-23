package controllers

import (
	"app/configs"
	"app/models/users"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	var userCreate users.UserRegister
	c.Bind(&userCreate)
	var userDB users.User
	password, _ := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.MinCost)
	userDB.Name = userCreate.Name
	userDB.Email = userCreate.Email
	userDB.Password = password
	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed register user database", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponse{
		true, "success register user database", userDB,
	})
}

func LoginUser(c echo.Context) error {
	var userCheck users.UserChecker
	c.Bind(&userCheck)
	var userDB users.User
	err := configs.DB.First(&userDB, "email = ? ", userCheck.Email).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "Failed get data user", nil,
		})
	}
	userResponse := users.UserData{
		userDB.Id, userDB.Name, userDB.Email, userDB.Password,
	}
	if checkPass := bcrypt.CompareHashAndPassword(userDB.Password, []byte(userCheck.Password)); checkPass != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponseData{
			false, "Invalid Password", userResponse,
		})
	}
	return c.JSON(http.StatusOK, users.UserResponse{
		true, "Success get data user", userDB,
	})
}
