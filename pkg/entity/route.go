package entity

// Route ...
type Route struct {
	Start Location `json:"start"`
	End   Location `json:"end"`
}

// Location  lat e long
type Location struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
