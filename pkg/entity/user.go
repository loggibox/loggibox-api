package entity

//User data
type User struct {
	ID                ID     `json:"id"`
	Birthday          int    `json:"birthday"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Gender            string `json:"gender"`
	Nacionality       string `json:"nacionality"`
	CivilStatus       string `json:"civilStatus"`
	Address           string `json:"address"`
	City              string `json:"city"`
	State             string `json:"state"`
	Phone             string `json:"phone"`
	FbInstitutionName string `json:"fbInstitutionName"`
}
