package data

type DrawData struct {
	Colors []uint32 `json:"colors"`
	Coords []uint32 `json:"coords"`
	T      string   `json:"t"`
}

type DrawResponse struct {
	Painted int64 `json:"painted"`
}
