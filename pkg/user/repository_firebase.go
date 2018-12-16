package user

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/db"
	"github.com/loggibox/loggibox-api/pkg/entity"
)

//FirebaseRepo in memory repo
type FirebaseRepo struct {
	client *db.Client
}

//NewFirebaseRepo create new repository
func NewFirebaseRepo(client *db.Client) *FirebaseRepo {
	return &FirebaseRepo{
		client: client,
	}
}

// Store a User
func (r *FirebaseRepo) Store(a *entity.User) (string, error) {
	// r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a User
func (r *FirebaseRepo) Find(id string) (*entity.User, error) {
	// if r.m[id.String()] == nil {
	// 	return nil, entity.ErrNotFound
	// }
	return nil, nil
}

//Search Users
func (r *FirebaseRepo) Search(query string) ([]entity.User, error) {
	var d []entity.User
	// for _, j := range r.m {
	// 	if strings.Contains(strings.ToLower(j.Name), query) {
	// 		d = append(d, j)
	// 	}
	// }
	// if len(d) == 0 {
	// 	return nil, entity.ErrNotFound
	// }

	return d, nil
}

//FindAll Users
func (r *FirebaseRepo) FindAll() ([]entity.User, error) {
	var d []entity.User
	ctx := context.Background()
	ref := r.client.NewRef("/users")
	var users map[string]entity.User
	if err := ref.Get(ctx, &users); err != nil {
		log.Fatalln("Error reading from database:", err)
		return nil, err
	}

	for _, value := range users {
		d = append(d, value)
	}
	fmt.Println(users)
	return d, nil
}

//Delete a User
func (r *FirebaseRepo) Delete(id string) error {
	// if r.m[id.String()] == nil {
	// 	return entity.ErrNotFound
	// }
	// r.m[id.String()] = nil
	return nil
}
