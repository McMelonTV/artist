package model

import (
	"artist/internal/proxy"

	"gorm.io/gorm"
)

type Proxy struct {
	gorm.Model
	proxy.Proxy
}
