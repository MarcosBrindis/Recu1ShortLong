package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotifyUpdatesMiddleware(updates *chan bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Writer.Status() >= http.StatusOK && ctx.Writer.Status() < http.StatusMultipleChoices {

			select {
			case *updates <- true:
				println("Notificación enviada al canal Updates")
			default:
				println("Canal Updates lleno, no se pudo enviar notificación")
			}
		}
	}
}
