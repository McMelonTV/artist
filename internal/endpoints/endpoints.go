package endpoints

import "fmt"

var BackendBase = "https://backend.wplace.live"
var TileDataEndpointFormat = BackendBase + "/files/s0/tiles/%d/%d.png"
var UserDataEndpoint = BackendBase + "/me"
var DrawEndpointFormat = BackendBase + "/s0/pixel/%d/%d"

func GetTileDataEndpoint(tileX, tileY int) string {
	return fmt.Sprintf(TileDataEndpointFormat, tileX, tileY)
}

func GetDrawEndpoint(tileX, tileY int) string {
	return fmt.Sprintf(DrawEndpointFormat, tileX, tileY)
}
