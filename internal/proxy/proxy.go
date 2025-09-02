package proxy

import (
	"fmt"

	"golang.org/x/net/proxy"
)

type Proxy struct {
	Host string
	Port string
	User string
	Pass string
}

func (p *Proxy) CreateDialer() (proxy.Dialer, error) {
	return proxy.SOCKS5("tcp", p.GetAddress(), p.GetAuth(), proxy.Direct)
}

func (p *Proxy) GetAddress() string {
	return fmt.Sprintf("%s:%s", p.Host, p.Port)
}

func (p *Proxy) GetAuth() *proxy.Auth {
	return &proxy.Auth{User: p.User, Password: p.Pass}
}
