package domain

type Service interface {
	Find(id int64) (*Job, error)
	FindAll() ([]*Job, error)
	Create(job *Job) error
	Update(job *Job) error
	Delete(id int64) error
}

type Repository interface {
	Find(id int64) (*Job, error)
	FindAll() ([]*Job, error)
	Create(job *Job) error
	Update(job *Job) error
	Delete(id int64) error
}

type service struct {
	jobRepository Repository
}

func NewJobService(jobRepository Repository) Service {
	return &service{jobRepository: jobRepository}
}

func (s *service) Find(id int64) (*Job, error) {
	return s.jobRepository.Find(id)
}

func (s *service) FindAll() ([]*Job, error) {
	return s.jobRepository.FindAll()
}

func (s *service) Create(job *Job) error {
	return s.jobRepository.Create(job)
}

func (s *service) Update(job *Job) error {
	return s.jobRepository.Update(job)
}

func (s *service) Delete(id int64) error {
	return s.jobRepository.Delete(id)
}
