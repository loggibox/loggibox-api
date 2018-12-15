package user

import (
	"strings"

	"github.com/loggibox/loggibox-api/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store an User
func (s *Service) Store(b *entity.User) (entity.ID, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		return entity.NewID(), err
	}
	b.ID = entity.ID(len(all) + 1)
	return s.repo.Store(b)
}

//Find a User
func (s *Service) Find(id entity.ID) (*entity.User, error) {
	return s.repo.Find(id)
}

//Search Users
func (s *Service) Search(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll Users
func (s *Service) FindAll() ([]*entity.User, error) {
	return s.repo.FindAll()
}

//Delete a User
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
