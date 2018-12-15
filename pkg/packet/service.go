package packet

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

//Store an Packet
func (s *Service) Store(b *entity.Packet) (entity.ID, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		return entity.NewID(), err
	}
	b.ID = entity.ID(len(all) + 1)
	return s.repo.Store(b)
}

//Find a Packet
func (s *Service) Find(id entity.ID) (*entity.Packet, error) {
	return s.repo.Find(id)
}

//Search Packets
func (s *Service) Search(query string) ([]*entity.Packet, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll Packets
func (s *Service) FindAll() ([]*entity.Packet, error) {
	return s.repo.FindAll()
}

//Delete a Packet
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
