package armor

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
)

type (
	Armor struct {
		Address      string           `json:"address"`
		TLS          *TLS             `json:"tls"`
		ReadTimeout  time.Duration    `json:"read_timeout"`
		WriteTimeout time.Duration    `json:"write_timeout"`
		Plugins      []Plugin         `json:"plugins"`
		Hosts        map[string]*Host `json:"hosts"`
		Logger       *log.Logger      `json:"-"`
		Colorer      *color.Color     `json:"-"`
	}

	TLS struct {
		Address  string `json:"address"`
		CertFile string `json:"cert_file"`
		KeyFile  string `json:"key_file"`
		Auto     bool   `json:"auto"`
		CacheDir string `json:"cache_dir"`
	}

	Host struct {
		Name     string           `json:"-"`
		CertFile string           `json:"cert_file"`
		KeyFile  string           `json:"key_file"`
		Plugins  []Plugin         `json:"plugins"`
		Paths    map[string]*Path `json:"paths"`
		Echo     *echo.Echo       `json:"-"`
	}

	Path struct {
		Plugins []Plugin `json:"plugins"`
	}

	Plugin map[string]interface{}
)

const (
	Version = "0.2.7"
	Website = "https://armor.labstack.com"
)
