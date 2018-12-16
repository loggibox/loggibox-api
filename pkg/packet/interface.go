package packet

import "github.com/loggibox/loggibox-api/pkg/entity"

//Reader interface
type Reader interface {
	Find(id string) (*entity.Packet, error)
	Search(query string) ([]entity.Packet, error)
	FindAll() ([]entity.Packet, error)
}

//Writer Packet writer
type Writer interface {
	Store(b *entity.Packet) (string, error)
	Update(id string, b *entity.Packet) (*entity.Packet, error)
	Delete(id string) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
