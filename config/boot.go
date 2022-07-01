package config

import (
	"log"

	"github.com/michelsazevedo/kuala/api"
	"github.com/michelsazevedo/kuala/domain"
	"github.com/michelsazevedo/kuala/repository"
)

type Boot struct {
	Handler  api.JobHandler
	Settings Settings
}

func NewBoot() *Boot {
	conf, err := NewConfig("./config/config.yaml")
	if err != nil {
		log.Default().Fatal(err)
	}

	repo, _ := repository.NewPostgresRepository(conf.Database.Host,
		conf.Database.Username, conf.Database.Password, conf.Database.Database)
	service := domain.NewJobService(repo)
	handler := api.NewHandler(service)

	return &Boot{Handler: handler, Settings: conf.Settings}
}
