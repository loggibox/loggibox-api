package entity

import "time"

// Packet ...
type Packet struct {
	ID                 string    `json:"id"`
	Weight             float64   `json:"weight"`
	Deadline           time.Time `json:"deadline"`
	DeadlineDays       int       `json:"deadline_days"`
	DeliveryAddress    string    `json:"delivery_address"`
	IDUser             string    `json:"id_user"`
	User               *User     `json:"user,omitempty"`
	RouteStart         string    `json:"route_start"`
	RouteEnd           string    `json:"route_end"`
	Delivered          bool      `json:"delivered"`
	Delivering         bool      `json:"delivering"`
	DistributionCenter bool      `json:"distribution_center"`
}
