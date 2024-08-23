package controllers

import (
	"backend/usecases"
	"backend/usecases/request"
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

func (uc *UserController) SignUp(c *gin.Context) {
	req, err := request.NewSignUpRequst(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.UserUsecase.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": res})
}

func (uc *UserController) SignIn(c *gin.Context) {
	req, err := request.NewSignInRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.UserUsecase.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": res})
}

func (uc *UserController) VerifyToken(c *gin.Context) {
	valid, err := uc.UserUsecase.TokenService.TokenValid(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": true})
}

func (uc *UserController) GetJoinedOrganizations(c *gin.Context) {
	userID, err := uc.UserUsecase.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.UserUsecase.GetUserJoinedOrganization(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (uc *UserController) GetRecievedInvitations(c *gin.Context) {
	userID, err := uc.UserUsecase.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	organizations, err := uc.UserUsecase.GetRecievedInvitationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"organizations": organizations})
}
