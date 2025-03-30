package config

import "github.com/yiixun/gotest01/v2/pkg/apex"

// This is an example to register and wire premitive types

const (
	Conf              = "apex.conf"
	DefaultLunchSeats = "apex.conf.default_lunch_seats" // 4
)

func NewConfig() {
	apex.Register(Conf, "./conf/default.yaml")
	apex.Register(DefaultLunchSeats, 4)
}
