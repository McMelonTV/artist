package client

import (
	"artist/internal/client/data"
	"artist/internal/endpoints"
	"artist/internal/proxy"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	JCookie    *http.Cookie
}

func New(proxy *proxy.Proxy, j string) (*Client, error) {
	dialer, err := proxy.CreateDialer()
	if err != nil {
		return nil, err
	}

	return &Client{
		HttpClient: &http.Client{
			Transport: &http.Transport{
				Dial: dialer.Dial,
			},
		},
		JCookie: &http.Cookie{
			Name:  "j",
			Value: j,
		},
	}, nil
}

func (c *Client) GetUserData(cfClearance string) (*data.UserData, error) {
	req, err := http.NewRequest("GET", endpoints.UserDataEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(c.JCookie)
	req.AddCookie(CreateCloudflareClearanceCookie(cfClearance))

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user data")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	userData := &data.UserData{}
	err = json.Unmarshal(body, userData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func CreateCloudflareClearanceCookie(cfClearance string) *http.Cookie {
	return &http.Cookie{
		Name:   "cf_clearance",
		Value:  cfClearance,
		Domain: ".wplace.live",
		Path:   "/",
		//Expires: ???,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
}
