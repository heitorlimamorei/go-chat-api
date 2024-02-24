package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/go-chat-api/repository"
	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func CreateUser(c *gin.Context) {
	req := UserRequest{}

	c.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.ErrorF("Error during the validation: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser := schemas.User{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}

	err := repository.CreateUser(&newUser)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "User created successfully", nil)
}

func UpdateUser(c *gin.Context) {
	req := UserRequest{}
	userId := c.Param("user_id")

	c.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.ErrorF("Error during the validation: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser := schemas.User{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}

	err := repository.UpdateUser(&newUser, userId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "User Saved sucessfully", nil)
}

func GetUserById(c *gin.Context) {
	userId := c.Param("user_id")

	if userId == "" {
		logger.Error("Request error: ID is required in the query parameters")
		sendError(c, http.StatusBadRequest, "Id is required")
		return
	}

	user, err := repository.GetChatById(userId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "User retrieved successfully", user)
}

func GetUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	email := c.Query("email")

	if email != "" {
		user, err := repository.GetUserByEmail(email)

		if err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		sendSucess(c, "User retrieved successfully", user)
		return
	}

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Users retrieved successfully", users)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("user_id")

	if userId == "" {
		logger.Error("Request error: ID is required in the query parameters")
		sendError(c, http.StatusBadRequest, "Id is required")
		return
	}

	err := repository.DeleteUSer(userId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "User deleted successfully", nil)
}
