package packet

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

// Store a Packet
func (r *FirebaseRepo) Store(a *entity.Packet) (entity.ID, error) {
	// r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a Packet
func (r *FirebaseRepo) Find(id entity.ID) (*entity.Packet, error) {
	// if r.m[id.String()] == nil {
	// 	return nil, entity.ErrNotFound
	// }
	return nil, nil
}

//Search Packets
func (r *FirebaseRepo) Search(query string) ([]*entity.Packet, error) {
	var d []*entity.Packet
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

//FindAll Packets
func (r *FirebaseRepo) FindAll() ([]*entity.Packet, error) {
	var d []*entity.Packet
	ctx := context.Background()
	ref := r.client.NewRef("/packets")
	var packets map[string]entity.Packet
	if err := ref.Get(ctx, &packets); err != nil {
		log.Fatalln("Error reading from database:", err)
		return nil, err
	}

	for _, value := range packets {
		d = append(d, &value)
	}
	fmt.Println(packets)
	return d, nil
}

//Delete a Packet
func (r *FirebaseRepo) Delete(id entity.ID) error {
	// if r.m[id.String()] == nil {
	// 	return entity.ErrNotFound
	// }
	// r.m[id.String()] = nil
	return nil
}
