package controller

import (
	"context"
	"net/http"
	"recuCorte1/src/user/application"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	UseCase *application.GetAllUsersUsecase
}

func NewGetAllUsersController(useCase *application.GetAllUsersUsecase) *GetAllUsersController {
	return &GetAllUsersController{UseCase: useCase}
}

func (c *GetAllUsersController) HandleGetAll(ctx *gin.Context) {
	reqContext := context.Background()
	users, err := c.UseCase.Execute(reqContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
