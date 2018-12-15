package entity

//User data
type User struct {
	ID         ID     `json:"id,omitempty"`
	Age        uint   `json:"age,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"address,omitempty"`
	RouteStart string `json:"route_start,omitempty"`
	RouteEnd   string `json:"route_end,omitempty"`
}
