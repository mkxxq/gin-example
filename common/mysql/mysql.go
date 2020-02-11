package mysql

import (
	_ "database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"ginexample/common/config"
)

const driverName = "mysql"

var SQLX_DB *sqlx.DB

// Setup method
func Setup() error {
	var err error
	SQLX_DB, err = sqlx.Open(driverName, config.Conf.CloudSql.Source)
	if err != nil {
		log.Printf("failed to connection MySQL error = %v", err)
		return err
	}
	err = SQLX_DB.Ping()
	if err != nil {
		log.Printf("failed to ping MySQL error = %v", err)
		return err
	}
	SQLX_DB.SetMaxIdleConns(config.Conf.CloudSql.MaxIdle)
	SQLX_DB.SetMaxOpenConns(config.Conf.CloudSql.MaxConn)
	SQLX_DB.SetConnMaxLifetime(time.Duration(config.Conf.CloudSql.MaxLifeTime) * time.Second)
	return nil
}
