package configurations

import (
	"strconv"
)

type Redis struct {
	Host        string
	Port        int
	Database    int
	Password    string
	MaxIdle     int
	IdleTimeout int64
}

func (r *Redis) Address() string {
	return r.Host + ":" + strconv.Itoa(r.Port)
}

func (r *Redis) DB() int {
	if r.Database < 0 || r.Database > 15 {
		return 0
	}
	return r.Database
}
