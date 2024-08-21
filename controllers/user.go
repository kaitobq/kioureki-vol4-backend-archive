package controllers

import (
	"backend/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase *usecases.UserUsecase
}

func NewUserController(userUsecase *usecases.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

type signUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (uc *UserController) SignUp(c *gin.Context) {
	var input signUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.CreateUser(input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}