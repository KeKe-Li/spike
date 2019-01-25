package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"flag"
	"path/filepath"

	"spike/models"
	"spike/config"
	"spike/helps/configurations"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const database =  "sqlite3"

func init(){
	var confFile string
	var pidFile string
	flag.Usage = func() 	{
		fmt.Println("Usage: " + os.Args[0] + " [options]\n\nOptions:")
		fmt.Println("\t-conf \t the configuration file. (Default: app.conf)")
		fmt.Println("\t-pid  \t the pid file. (Default: app.pid)")
	}
	flag.StringVar(&confFile, "conf", "app.conf", "path to configuration file.")
	flag.StringVar(&pidFile, "pid", "app.pid", "path to pid file.")
	flag.Parse()
	absConfPath, _ := filepath.Abs(confFile)
	absPidPath, _ := filepath.Abs(pidFile)

	appConfig = &config.ApplicationConfig{}
	// load configuration
	parser := &configurations.TomlParser{}
	err := parser.LoadConfiguration(appConfig, absConfPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	//// init the mysql
	//mysqlDB, err = connection_pool.NewMysql(appConfig.Mysql)
	//if err != nil {
	//	panic(err.Error())
	//}
	sqlite, err = gorm.Open(database, appConfig.Sqlite)
	if err != nil {
		fmt.Println(err.Error())
	}

	sqlite.AutoMigrate(
		&models.MicroService{},
		&models.GpgKey{},
		&models.User{},
	)

	if err := ioutil.WriteFile(absPidPath, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
		fmt.Println(err.Error())
	}
}
