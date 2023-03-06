package main

import (
	"context"
	"os"
	"test3/config"
	"test3/container"
	"test3/repository/gormrepo"
	"test3/storage"
	_ "time/tzdata"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use: "gits case 3",
}

func init() {
	loadConfig()
}

func main() {
	Execute()
}

func Execute() {
	rootCmd := registerCommands(&defaultAppProvider{})
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
}

func loadConfig() {
	err := config.Load()
	if err != nil {
		logrus.Errorf("Config error: %s", err.Error())
		os.Exit(1)
	}
}

func registerCommands(appProvider AppProvider) *cobra.Command {
	rootCmd.AddCommand(Server(appProvider))
	rootCmd.AddCommand(Migrate(appProvider))

	return rootCmd
}

type AppProvider interface {
	BuildContainer(ctx context.Context, options buildOptions) (*container.Container, func(), error)
}

type buildOptions struct {
	Postgres bool
	RabbitMq bool
}

type defaultAppProvider struct {
}

func (defaultAppProvider) BuildContainer(ctx context.Context, options buildOptions) (*container.Container, func(), error) {
	var db *gorm.DB
	cfg := config.Instance()

	appContainer := container.NewContainer()
	appContainer.SetConfig(cfg)

	if options.Postgres {
		db = storage.GetPostgresDb()
		appContainer.SetDb(db)
		bookRepo := gormrepo.NewBookRepository(db)
		appContainer.SetBookRepo(bookRepo)
		authorRepo := gormrepo.NewAuthorRepository(db)
		appContainer.SetAuthorRepo(authorRepo)
		publisherRepo := gormrepo.NewPublisherRepository(db)
		appContainer.SetPublisherRepo(publisherRepo)
	}

	deferFn := func() {
		if db != nil {
			err := storage.CloseDB(db)
			if err != nil {
				logrus.Errorf("Error when closing db, error: %s", err.Error())
			}
		}
	}

	return appContainer, deferFn, nil
}
