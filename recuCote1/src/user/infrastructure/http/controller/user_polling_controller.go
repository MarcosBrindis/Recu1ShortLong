package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"recuCorte1/src/user/application"

	"github.com/gin-gonic/gin"
)

// UserPollingController
type UserPollingController struct {
	GetAllUsersUsecase *application.GetAllUsersUsecase
	Updates            *chan bool
}

func NewUserPollingController(getAllUsecase *application.GetAllUsersUsecase, updates *chan bool) *UserPollingController {
	return &UserPollingController{
		GetAllUsersUsecase: getAllUsecase,
		Updates:            updates,
	}
}

// HandleShortPoll
func (c *UserPollingController) HandleShortPoll(ctx *gin.Context) {
	reqCtx := context.Background()
	users, err := c.GetAllUsersUsecase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// HandleLongPoll
func (c *UserPollingController) HandleLongPoll(ctx *gin.Context) {
	select {
	case <-*c.Updates:
		println("Se recibió una notificación de cambio")
		*c.Updates = make(chan bool, 1)
	case <-time.After(5 * time.Second):
		println("Timeout: No hubo cambios en 30 segundos")
	}

	reqCtx := context.Background()
	users, err := c.GetAllUsersUsecase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// HandleCountShortPollStreaming
func (c *UserPollingController) HandleCountShortPollStreaming(ctx *gin.Context) {

	writer := ctx.Writer
	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")

	flusher, ok := writer.(http.Flusher)
	if !ok {
		ctx.String(http.StatusInternalServerError, "Streaming no soportado")
		return
	}

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:

			fmt.Fprintf(writer, "data: %s\n\n", "timeout")
			flusher.Flush()
			return
		case <-ticker.C:
			reqCtx := context.Background()
			users, err := c.GetAllUsersUsecase.Execute(reqCtx)
			if err != nil {
				fmt.Fprintf(writer, "data: error: %s\n\n", err.Error())
				flusher.Flush()
				continue
			}

			countTrue := 0
			countFalse := 0
			for _, user := range users {
				if user.Sexo {
					countTrue++
				} else {
					countFalse++
				}
			}

			fmt.Fprintf(writer, "data: {\"count_true\": %d, \"count_false\": %d}\n\n", countTrue, countFalse)
			flusher.Flush()
		}
	}
}
