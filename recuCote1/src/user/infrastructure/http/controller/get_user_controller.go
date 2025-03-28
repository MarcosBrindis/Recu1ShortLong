package controller

import (
	"context"
	"net/http"
	"recuCorte1/src/user/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	UseCase *application.GetUserUsecase
}

func NewGetUserController(useCase *application.GetUserUsecase) *GetUserController {
	return &GetUserController{UseCase: useCase}
}

func (c *GetUserController) HandleGet(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	user, err := c.UseCase.Execute(reqContext, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
