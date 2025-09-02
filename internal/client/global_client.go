package client

import (
	"artist/internal/endpoints"
	"fmt"
	"io"
	"net/http"
)

var GC = &GlobalClient{}

type GlobalClient struct {
}

func (c *GlobalClient) GetTileData(tileX, tileY int) ([]byte, error) {
	res, err := http.Get(endpoints.GetTileDataEndpoint(tileX, tileY))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get tile data for %d %d", tileX, tileY)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
