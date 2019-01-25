package configurations

import (
	"fmt"
)

// mysql database configuration
type Mysql struct {
	Type          string // type 业务逻辑标识类型，预留（可以是读或写或其它）
	UserName      string
	Password      string
	Database      string
	Address       string
	Parameters    string
	MaxIdle       int
	MaxOpen       int
	Debug         bool
	MigrationsDir string
}

//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (conf *Mysql) String() string {
	return fmt.Sprintf("%s:%s@%s/%s?%s", conf.UserName, conf.Password, conf.Address, conf.Database, conf.Parameters)
}
