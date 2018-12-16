package packet

import (
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
func (s *Service) Store(b *entity.Packet) (string, error) {
	return s.repo.Store(b)
}

//Find a Packet
func (s *Service) Find(id string) (*entity.Packet, error) {
	return s.repo.Find(id)
}

//Search Packets
func (s *Service) Search(query string) ([]entity.Packet, error) {
	return s.repo.Search(query)
}

//FindAll Packets
func (s *Service) FindAll() ([]entity.Packet, error) {
	return s.repo.FindAll()
}

//Delete a Packet
func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
