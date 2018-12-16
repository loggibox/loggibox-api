package packet

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

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
func (r *FirebaseRepo) Store(a *entity.Packet) (string, error) {
	ctx := context.Background()
	ref := r.client.NewRef("/")
	packetsRef := ref.Child("packets")
	// Generate a reference to a new location and add some data using Push()
	newPacketRef, err := packetsRef.Push(ctx, a)
	if err != nil {
		log.Fatalln("Error pushing child node:", err)
	}
	// // Get the unique key generated by Push()
	packetID := newPacketRef.Key
	packetsRef = ref.Child("/packets/" + packetID)
	if err := packetsRef.Update(ctx, map[string]interface{}{
		"id": packetID,
	}); err != nil {
		log.Fatalln("Error updating child:", err)
		return a.ID, err
	}

	return packetID, nil
}

//Find a Packet
func (r *FirebaseRepo) Find(id string) (*entity.Packet, error) {
	var d *entity.Packet
	ctx := context.Background()
	ref := r.client.NewRef("/packets")
	var packets map[string]entity.Packet
	if err := ref.Get(ctx, &packets); err != nil {
		log.Fatalln("Error reading from database:", err)
		return nil, err
	}

	fmt.Println(packets)
	return d, nil
}

//Search Packets
func (r *FirebaseRepo) Search(query string) ([]entity.Packet, error) {
	var d []entity.Packet
	ctx := context.Background()
	ref := r.client.NewRef("/packets")

	fmt.Println(query)
	filter := strings.Split(query, "=")
	fBool, err := strconv.ParseBool(filter[1])

	results, err := ref.OrderByChild(filter[0]).EqualTo(fBool).GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
		return nil, err
	}
	for _, r := range results {
		var packet entity.Packet
		if err := r.Unmarshal(&packet); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
			continue
		}
		d = append(d, packet)
	}

	return d, nil
}

//FindAll Packets
func (r *FirebaseRepo) FindAll() ([]entity.Packet, error) {
	var d []entity.Packet

	ctx := context.Background()
	ref := r.client.NewRef("/packets")
	var packets map[string]entity.Packet
	if err := ref.Get(ctx, &packets); err != nil {
		log.Fatalln("Error reading from database:", err)
		return nil, err
	}

	for _, value := range packets {
		// fmt.Println(value)
		d = append(d, value)
	}

	// fmt.Println(packets)
	return d, nil
}

//Delete a Packet
func (r *FirebaseRepo) Delete(id string) error {
	// if r.m[id.String()] == nil {
	// 	return entity.ErrNotFound
	// }
	// r.m[id.String()] = nil
	return nil
}
