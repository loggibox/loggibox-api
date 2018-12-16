package entity

//User data
type User struct {
	ID         string `json:"id"`
	Age        uint   `json:"age"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	RouteStart string `json:"route_start"`
	RouteEnd   string `json:"route_end"`
}
