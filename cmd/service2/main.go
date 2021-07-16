package main

import (
	"FileManage/internal/server2/consumer"
	"FileManage/pkg/db"
	"log"
)

func main() {
	err := db.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	consumer.Extract("one")

	select {

	}
}