package req_types

type Tile struct {
	Id    uint    `json:"id"`
	Width uint    `json:"width"`
	Ratio float64 `json:"ratio"`
}

type AdContext struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

type IncomingRequest struct {
	Id      *string   `json:"id"`
	Tiles   []Tile    `json:"tiles"`
	Context AdContext `json:"context"`
}

// Based in Tile
type Imp struct {
	Id        uint `json:"id"`        // same as related `Tile.Id`
	Minwidth  uint `json:"minwidth"`  // `Tile.Width`
	Minheight uint `json:"minheight"` // math.Floor(Tile.Width * Tile.Ratio)
}

type OutgoingRequest struct {
	Id      string    `json:"id"`
	Imp     []Imp     `json:"imp"`
	Context AdContext `json:"context"`
}
