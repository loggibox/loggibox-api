package entity

// HTTPResp ...
type HTTPResp struct {
	Status   string      `json:"status,omitempty"`
	Code     int         `json:"code,omitempty"`
	Messages []string    `json:"messages,omitempty"`
	Result   interface{} `json:"result,omitempty"`
}
