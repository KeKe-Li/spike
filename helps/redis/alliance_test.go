package redis

import (
	"testing"
	"fmt"
	"earlydata.com/gopher/configurations"
	"earlydata.com/gopher/connection-pool"
	"encoding/json"
)

func TestUcan(t *testing.T) {
	//Host = "localhost"
	//Port = 6379
	//Database = 12
	//Password =""
	//MaxIdle = 10
	//IdleTimeout = 20000000000
	redisConf := &configurations.Redis{}
	redisConf.Host = "localhost"
	redisConf.Port = 6379
	redisConf.Database = 12
	redisConf.MaxIdle = 10
	redisConf.IdleTimeout = 20000000000
	redispool := connection_pool.NewRedisPool(redisConf)
	alliance := Alliance{Redis:redispool, Key:"ALLIANCE::AUTHORITY"}
	data := alliance.GetAll()
	authorities := make([]Authority, 0)
	json.Unmarshal([]byte(data), &authorities)
	re := Ucan(authorities, "123", "456", "321")
	fmt.Println(re)
}