package config

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/michelsazevedo/kuala/api"
	"github.com/michelsazevedo/kuala/domain"
	"github.com/michelsazevedo/kuala/repository"
)

type Application struct {
	Routes   chi.Router
	Settings Settings
}

func NewApplication() *Application {
	conf, err := NewConfig("./config/config.yaml")
	if err != nil {
		log.Default().Fatal(err)
	}

	repo, _ := repository.NewPostgresRepository(conf.Database.Host,
		conf.Database.Username, conf.Database.Password, conf.Database.Database)
	service := domain.NewJobService(repo)
	handler := api.NewHandler(service)

	return &Application{Routes: Routes(handler), Settings: conf.Settings}
}
