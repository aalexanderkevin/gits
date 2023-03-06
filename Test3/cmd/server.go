package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"test3/config"
	handler "test3/handlers"
	"test3/transport/grpc/gitspb"
	"test3/usecase"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func Server(appProvider AppProvider) *cobra.Command {
	cliCommand := &cobra.Command{
		Use:   "server",
		Short: "Run the gRPC server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			app, closeResourcesFn, err := appProvider.BuildContainer(ctx, buildOptions{
				Postgres: true,
			})
			if err != nil {
				return err
			}
			if closeResourcesFn != nil {
				defer closeResourcesFn()
			}

			cfg := config.Instance()

			usecaseBook := usecase.NewBook(app)
			usecaseAuthor := usecase.NewAuthor(app)
			usecasePublisher := usecase.NewPublisher(app)

			// Start Http Server
			lis, err := net.Listen("tcp", cfg.Service.Host+":"+cfg.Service.Port)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			server := grpc.NewServer()
			gitspb.RegisterGitsServer(server, handler.NewGrpcHandler(usecaseBook, usecaseAuthor, usecasePublisher))

			term := make(chan os.Signal)
			go func() {
				log.Println("grpc server listening at " + lis.Addr().String())
				if err := server.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
					term <- syscall.SIGINT
				}
			}()

			signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)
			<-term
			log.Println("shutting down")
			server.GracefulStop()
			return nil
		},
	}
	return cliCommand
}
