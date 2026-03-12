package main

import (
	"context"
	"fmt"
	"go-mux-mongodb-user-manager-api/internal/infra/database/mongodb"
	"go-mux-mongodb-user-manager-api/internal/infra/security"
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
	"go-mux-mongodb-user-manager-api/internal/web"
	"go-mux-mongodb-user-manager-api/pkg/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	err := configs.Connect()
	returnFatalError(err)

	col := initDatabase()
	mongorepo := mongodb.NewMongoRepository(col)

	createUseCase := users_manager.NewCreateUserServices(mongorepo, security.NewHashingService())
	getAllUseCase := users_manager.NewGetAllServices(mongorepo)
	getLoginUseCase := users_manager.NewLoginUserServices(mongorepo, security.NewHashingService())
	updateNameUseCase := users_manager.NewUpdateNameServices(mongorepo, security.NewHashingService())
	updateEmailUseCase := users_manager.NewUpdateEmailServices(mongorepo, security.NewHashingService())

	type handler struct {
		*users_manager.CreateUserServices
		*users_manager.GetAllServices
		*users_manager.LoginUserServices
		*users_manager.UpdateNameServices
		*users_manager.UpdateEmailServices
	}

	hub := handler{
		createUseCase,
		getAllUseCase,
		getLoginUseCase,
		updateNameUseCase,
		updateEmailUseCase,
	}

	usecases := web.NewUserUseCasesRepository(hub)

	err = web.Routers(usecases)
	returnFatalError(err)

	fmt.Println("server running on port " + configs.Env.Port)
	fmt.Println("server running on address http://localhost:" + configs.Env.Port)
}

func initDatabase() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongodb.NewMongoConnection(ctx)
	returnFatalError(err)

	configs.LoadEnv(client)

	col := mongodb.NewCollection(client, configs.Env.DbCol)

	errIndex := mongodb.IndexUnique(col)
	returnFatalError(errIndex)

	return col
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
