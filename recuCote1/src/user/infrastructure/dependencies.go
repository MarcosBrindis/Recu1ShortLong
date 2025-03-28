package infrastructure

import (
	"recuCorte1/src/user/application"
	"recuCorte1/src/user/infrastructure/arreglo" // Importar el paquete arreglo
	"recuCorte1/src/user/infrastructure/http/controller"
)

var (
	CreateUserController  *controller.CreateUserController
	GetUserController     *controller.GetUserController
	GetAllUsersController *controller.GetAllUsersController
	UserPollingController *controller.UserPollingController
	Updates               chan bool
)

func InitDependencies() {
	// Inicializar el repositorio de usuario (en memoria)
	userRepo := arreglo.NewUserRepository()

	// Crear los casos de uso
	createUserUsecase := application.CreateUserUsecase{Repository: userRepo}
	getUserUsecase := application.GetUserUsecase{Repository: userRepo}
	getAllUsersUsecase := application.GetAllUsersUsecase{Repository: userRepo}

	// Crear el canal de notificaciones
	Updates = make(chan bool, 1)

	// Crear instancias de los controladores
	CreateUserController = controller.NewCreateUserController(&createUserUsecase)
	GetUserController = controller.NewGetUserController(&getUserUsecase)
	GetAllUsersController = controller.NewGetAllUsersController(&getAllUsersUsecase)

	// Crear controlador de polling
	UserPollingController = controller.NewUserPollingController(&getAllUsersUsecase, &Updates)
}
