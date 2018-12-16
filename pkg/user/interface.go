package user

import "github.com/loggibox/loggibox-api/pkg/entity"

//Reader interface
type Reader interface {
	Find(id string) (*entity.User, error)
	Search(query string) ([]entity.User, error)
	FindAll() ([]entity.User, error)
}

//Writer User writer
type Writer interface {
	Store(b *entity.User) (string, error)
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
