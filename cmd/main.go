package main

import (
	"github.com/TwiLightDM/users-service/internal/configs"
	"github.com/TwiLightDM/users-service/internal/database"
	"github.com/TwiLightDM/users-service/internal/transport/grpc"
	"github.com/TwiLightDM/users-service/internal/user"
	"log"
)

func main() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDB(cfg.HostDatabase, cfg.PortDatabase, cfg.User, cfg.Password, cfg.DbName)
	if err != nil {
		log.Fatal(err)
	}

	repo := user.NewUserRepository(db)
	svc := user.NewUserService(repo)

	if err = grpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
