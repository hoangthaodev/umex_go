package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"umex.com/global"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db

	// Set pool
	setPool()
	// migration
	migrationTable()
}

func setPool() {
	m := global.Config.Mysql
	sqlDb := global.Mdb
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnectMaxLifeTime))
}

func migrationTable() {

}
