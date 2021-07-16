package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlDBConfig struct {
	Username    string
	Password    string
	Port        int
	Host        string
	Database    string
	maxOpen     int
	maxIdle     int
	maxLifetime time.Duration
}

var DB *sql.DB

var MdbConfig = &MysqlDBConfig{
	"root",
	"123456",
	3306,
	"127.0.0.1",
	"todolist",
	100,
	10,
	60,
}


func InitMysql() error {
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", MdbConfig.Username, MdbConfig.Password, MdbConfig.Host,
		MdbConfig.Port,
		MdbConfig.Database)
	db, err := sql.Open("mysql", config)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(MdbConfig.maxIdle)
	db.SetConnMaxLifetime(MdbConfig.maxLifetime)
	db.SetMaxOpenConns(MdbConfig.maxOpen)
	err = db.Ping()
	if err != nil {
		return err
	}

	DB = db
	return nil
}
