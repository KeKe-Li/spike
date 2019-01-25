package redis

import (
	"github.com/garyburd/redigo/redis"
)

type Alliance struct {
	Redis *redis.Pool `json:"redis"`
	Key   string `json:"key"`
}

type Authority struct {
	DisplayName string //属性
	Value       string //属性值
	Children    []Authority `json:"children"`
}

func (ctrl Alliance)GetAll() string {
	value, err := redis.String(ctrl.Redis.Get().Do("get", ctrl.Key))
	if err != nil {
		return ""
	}
	return value
}

func (ctrl Alliance)Save(value interface{}) error {
	_, err := ctrl.Redis.Get().Do("set", ctrl.Key, value)
	if err != nil {
		return err
	}
	return nil
}

func Ucan(authorities []Authority, values ...string) bool {
	if len(values) <= 0 {
		return false
	}
	now := values[:1]
	end := values[1:]
	can := false
	for _, item := range authorities {
		if item.Value == now[0] {
			can = true
			break
		}
	}
	if can && len(end) > 0 {
		return Ucan(authorities[0].Children, end...)
	}

	return can
}



