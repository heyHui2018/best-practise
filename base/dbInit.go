package base

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/heyHui2018/log"
)

var DBEngine *xorm.Engine

func DbInit() {
	var err error
	username := GetConfig().DB.Username
	password := GetConfig().DB.Password
	database := GetConfig().DB.Database
	host := GetConfig().DB.Host
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, database)
	DBEngine, err = xorm.NewEngine("mysql", params)
	if err != nil {
		log.Fatalf("初始化DB出错,error = %v", err)
	}
	log.Infof("初始化DB完成,host = %v,username = %v,database = %v", host, username, database)
	DBEngine.SetMaxOpenConns(GetConfig().DB.MaxOpenConn)
	DBEngine.SetMaxIdleConns(GetConfig().DB.MaxIdleConn)
	// DB.ShowSQL()
}
