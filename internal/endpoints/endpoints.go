package endpoints

import "fmt"

var BackendBase = "https://backend.wplace.live"
var TileDataEndpointFormat = BackendBase + "/files/s0/tiles/%d/%d.png"
var UserDataEndpoint = BackendBase + "/me"

func GetTileDataEndpoint(tileX, tileY int) string {
	return fmt.Sprintf(TileDataEndpointFormat, tileX, tileY)
}
