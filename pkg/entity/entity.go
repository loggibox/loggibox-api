package entity

import (
	"strconv"
)

//ID type
type ID uint

// HTTPResp ...
type HTTPResp struct {
	Status   string      `json:"status,omitempty"`
	Code     int         `json:"code,omitempty"`
	Messages []string    `json:"messages,omitempty"`
	Result   interface{} `json:"result,omitempty"`
}

//ToString convert an ID in a string
func (i ID) String() string {
	return string(i)
}

//StringToID convert a string to an ID
func StringToID(s string) ID {
	id, _ := strconv.ParseUint(s, 10, 32)
	return ID(id)
}

//NewID create a new id
func NewID() ID {
	return ID(1)
}
