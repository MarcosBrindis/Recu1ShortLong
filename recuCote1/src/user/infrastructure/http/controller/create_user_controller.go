package controller

import (
	"context"
	"net/http"
	"recuCorte1/src/user/application"
	"recuCorte1/src/user/domain/model"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	UseCase *application.CreateUserUsecase
}

func NewCreateUserController(useCase *application.CreateUserUsecase) *CreateUserController {
	return &CreateUserController{UseCase: useCase}
}

func (c *CreateUserController) HandleCreate(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}
