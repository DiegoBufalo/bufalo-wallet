package routes

import (
	"bufalo-wallet-app/controllers"
	"bufalo-wallet-app/database"
	"bufalo-wallet-app/repositories"
	"bufalo-wallet-app/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	db := database.DB

	// Instancia os repositórios
	transactionRepo := repositories.NewTransactionRepository(db)
	userRepo := repositories.NewUserRepository(db)         // Novo repositório para usuários
	categoryRepo := repositories.NewCategoryRepository(db) // Novo repositório para categorias

	// Instancia os serviços, passando os repositórios como dependências
	transactionService := services.NewTransactionService(transactionRepo)
	userService := services.NewUserService(userRepo)             // Novo serviço para usuários
	categoryService := services.NewCategoryService(categoryRepo) // Novo serviço para categorias

	// Instancia os controllers, passando os serviços como dependências
	transactionController := controllers.NewTransactionController(transactionService)
	userController := controllers.NewUserController(userService)             // Novo controller para usuários
	categoryController := controllers.NewCategoryController(categoryService) // Novo controller para categorias

	api := router.Group("/api")

	// Rotas de transações
	api.POST("/transactions", transactionController.CreateTransaction)
	api.GET("/transactions/user/:user_id", transactionController.GetUserTransactions)
	api.PUT("/transactions", transactionController.UpdateTransaction)
	api.DELETE("/transactions/:transaction_id", transactionController.DeleteTransaction)

	// Rotas de usuários
	api.POST("/users", userController.CreateUser)          // Novo endpoint para criar usuário
	api.GET("/users/:user_id", userController.GetUserByID) // Novo endpoint para buscar usuário por ID

	// Rotas de categorias
	// Rotas de categorias
	api.POST("/categories", categoryController.CreateCategory)              // Endpoint para criar uma categoria
	api.GET("/categories", categoryController.GetCategories)                // Endpoint para listar categorias
	api.GET("/categories/:category_id", categoryController.GetCategoryByID) // Endpoint para buscar categoria por ID

	return router
}
