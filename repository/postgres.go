package repository

import (
	"bytes"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/michelsazevedo/kuala/domain"
	"github.com/pkg/errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type postgresRepository struct {
	db *pg.DB
}

func NewPostgresRepository(host, username, password, database string) (domain.Repository, error) {
	db := pg.Connect(&pg.Options{
		Addr:     host,
		User:     username,
		Password: password,
		Database: database,
	})

	if schemaErr := createSchema(db); schemaErr != nil {
		return nil, errors.New("error to load database schema")
	}

	repo := &postgresRepository{
		db: db,
	}

	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		errors.New("postgres is down")
	}

	return repo, nil
}

func createSchema(db *pg.DB) error {
	for _, models := range []interface{}{(*domain.Job)(nil)} {
		if err := db.Model(models).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			panic(err)
		}
	}

	return nil
}

func (r *postgresRepository) Find(id int64) (*domain.Job, error) {
	job := &domain.Job{Id: id}

	if err := r.db.Model(job).WherePK().Select(); err != nil {
		return nil, err
	}

	return job, nil
}

func (r *postgresRepository) FindAll() ([]*domain.Job, error) {
	jobs := []*domain.Job{}

	if err := r.db.Model(&jobs).Order("job.created_at DESC").Select(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *postgresRepository) Create(job *domain.Job) error {
	if err := Validates(job); err != nil {
		return err
	}

	_, err := r.db.Model(job).Insert()
	return err
}

func (r *postgresRepository) Update(job *domain.Job) error {
	if err := Validates(job); err != nil {
		return err
	}

	_, err := r.db.Model(job).WherePK().Update()
	return err
}

func (r *postgresRepository) Delete(id int64) error {
	_, err := r.db.Model(&domain.Job{Id: id}).WherePK().Delete()
	return err
}

func Validates(job *domain.Job) error {
	validate = validator.New()
	err := validate.Struct(job)

	if err != nil {
		buf := bytes.NewBufferString("Validation Error: ")

		for _, err := range err.(validator.ValidationErrors) {
			buf.WriteString(fmt.Sprintf("Field %s is %s. ", err.Field(), err.ActualTag()))
		}

		return errors.New(buf.String())
	}

	return nil
}
