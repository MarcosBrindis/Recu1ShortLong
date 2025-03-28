package main

import (
	userInfra "recuCorte1/src/user/infrastructure"
	userRouter "recuCorte1/src/user/infrastructure/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar dependencias
	userInfra.InitDependencies()

	// Crear instancia de Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins for testing
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar rutas con los controladores globales
	userRouter.SetupUserRoutes(r, userInfra.CreateUserController, userInfra.GetUserController, userInfra.GetAllUsersController, userInfra.UserPollingController, &userInfra.Updates)

	// Iniciar el servidor
	r.Run(":8080")
}
