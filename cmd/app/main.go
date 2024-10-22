package main

import (
	"log"
	"products/config"
	"products/internal/grpcServer"
	"products/pkg/logger"
)

func main() {
	log.Println("Starting server")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot cload config: %v", err.Error())
	}

	appLogger := logger.NewApiLogger(cfg)
	err = appLogger.InitLogger()
	if err != nil {
		log.Fatalf("Cannot init logger: %v", err.Error())
	}

	s := grpcServer.NewServer(cfg, appLogger)
	if err = s.MapHandlers(appLogger); err != nil {
		appLogger.Errorf("Error mapping handlers: %v", err)
		return
	}

	if err = s.Run(); err != nil {
		appLogger.Errorf("Server run error: %v", err)
	}
}
