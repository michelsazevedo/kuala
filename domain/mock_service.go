package domain

import (
	"errors"
	"time"
)

type ServiceMock struct{}

func (*ServiceMock) Create(job *Job) error {
	if job.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

func (*ServiceMock) Update(job *Job) error {
	if job.Id != 1 {
		return errors.New("Job not found")
	}

	if job.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

func (*ServiceMock) Delete(id int64) error {
	if id != 1 {
		return errors.New("Job not found")
	}

	return nil
}

func (*ServiceMock) Find(id int64) (*Job, error) {
	if id != 1 {
		return nil, errors.New("Job not found")
	}

	now := time.Now()
	job := &Job{
		Id:          1,
		Title:       "Developer",
		Description: "Developer",
		CompanyId:   1,
		Tags:        []string{"Backend"},
		Featured:    false,
		PublishedAt: now,
		CreatedAt:   now,
		ExpiresAt:   now.AddDate(0, 0, 30),
	}
	return job, nil
}

func (*ServiceMock) FindAll() ([]*Job, error) {
	now := time.Now()
	mockedJobs := []*Job{{
		Id:          1,
		Title:       "Developer",
		Description: "Developer",
		CompanyId:   1,
		Tags:        []string{"Backend"},
		Featured:    false,
		PublishedAt: now,
		CreatedAt:   now,
		ExpiresAt:   now.AddDate(0, 0, 30),
	}}
	return mockedJobs, nil
}

func NewServiceMock() Service {
	return &ServiceMock{}
}
