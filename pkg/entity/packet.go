package entity

import "time"

// Packet ...
type Packet struct {
	ID                 ID        `json:"id,omitempty"`
	Weight             float64   `json:"weight,omitempty"`
	Deadline           time.Time `json:"deadline,omitempty"`
	DeliveryAddress    string    `json:"delivery_address,omitempty"`
	IDUser             ID        `json:"id_user,omitempty"`
	RouteStart         string    `json:"route_start,omitempty"`
	RouteEnd           string    `json:"route_end,omitempty"`
	Delivered          bool      `json:"delivered,omitempty"`
	Delivering         bool      `json:"delivering,omitempty"`
	DistributionCenter bool      `json:"distribution_center,omitempty"`
}
