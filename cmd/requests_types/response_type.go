package req_types

// Price is interface (but will be used as float64) 'cuz it can be quoted or unquoted
type RespImp struct {
	Id     uint   `json:"id"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	// Price expects as float64, but it may be quoted or unquoted
	Price interface{} `json:"price"`
}

// Response from ad partners
type SuccesResponse struct {
	Id  string    `json:"id"`
	Imp []RespImp `json:"imp"`
}
