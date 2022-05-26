package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	DbUpdate(id int64, updateEvent *Event) (*Event, error)
	DbCreated(newEvent *Event) (*Event, error)
	DbDelete(id int64) (bool, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) DbUpdate(id int64, updateEvent *Event) (*Event, error) {
	return (*s.repo).DbUpdate(id, updateEvent)
}

func (s *service) DbCreated(newEvent *Event) (*Event, error) {
	return (*s.repo).DbCreated((newEvent))
}

func (s *service) DbDelete(id int64) (bool, error) {
	return (*s.repo).DbDelete(id)
}
