package grpcServer

import (
	"fmt"
	productsv1 "github.com/Lineblaze/products_protos/gen/go/products"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"products/config"
	grpcHandler "products/internal/delivery/grpc"
	repository "products/internal/repository"
	useCase "products/internal/usecase"
	"products/pkg/logger"
	storage "products/pkg/storage/postgres"
)

type Server struct {
	grpcServer *grpc.Server
	cfg        *config.Config
	apiLogger  *logger.ApiLogger
}

func NewServer(cfg *config.Config, logger *logger.ApiLogger) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		cfg:        cfg,
		apiLogger:  logger,
	}
}

func (s *Server) MapHandlers(logger *logger.ApiLogger) error {
	db, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}
	repo := repository.NewPostgresRepository(db, logger)
	useCase := useCase.NewUseCase(repo, logger)
	handler := grpcHandler.NewHandler(useCase, logger)

	productsv1.RegisterProductServiceServer(s.grpcServer, handler)

	return nil
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.cfg.Server.Port))
	if err != nil {
		s.apiLogger.Fatalf("Failed to listen on port %s: %v", s.cfg.Server.Port, err)
		return err
	}

	s.apiLogger.Infof("Starting GRPC server on port %s", s.cfg.Server.Port)

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			s.apiLogger.Fatalf("Failed to serve GRPC server: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done

	s.apiLogger.Info("Shutting down GRPC server gracefully...")
	s.grpcServer.GracefulStop()
	s.apiLogger.Info("Server gracefully stopped")

	return nil
}
