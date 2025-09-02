package data

type UserData struct {
	AllianceId   uint32 `json:"allianceId"`
	AllianceRole string `json:"allianceRole"`
	Charges      struct {
		CooldownMs uint32  `json:"cooldownMs"`
		Count      float64 `json:"count"`
		Max        uint32  `json:"max"`
	} `json:"charges"`
	Country           string `json:"country"`
	Discord           string `json:"discord"`
	Droplets          uint32 `json:"droplets"`
	EquippedFlag      uint32 `json:"equippedFlag"`
	ExtraColorsBitmap uint32 `json:"extraColorsBitmap"`
	FavoriteLocations []struct {
		Id        uint32  `json:"id"`
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"favoriteLocations"`
	FlagsBitmap            string  `json:"flagsBitmap"`
	Id                     uint32  `json:"id"`
	IsCustomer             bool    `json:"isCustomer"`
	Level                  float64 `json:"level"`
	MaxFavoriteLocations   uint32  `json:"maxFavoriteLocations"`
	Name                   string  `json:"name"`
	NeedsPhoneVerification bool    `json:"needsPhoneVerification"`
	Picture                string  `json:"picture"`
	PixelsPainted          uint32  `json:"pixelsPainted"`
	ShowLastPixel          bool    `json:"showLastPixel"`
}
