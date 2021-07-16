package main

import (
	"FileManage/handler"
	"FileManage/pkg/cache"
	"FileManage/pkg/db"
	"FileManage/pkg/nsq"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	err := cache.InitCache()
	if err != nil {
		log.Fatal(err)
	}
	err = db.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", handler.Upload)
	router.GET("/extract", handler.SingnalQueue)
	router.StaticFS("/file", http.Dir("D:\\APROJECT\\GO-Project\\src\\FileManage\\file"))
	err = nsq.InitProducer()
	if err != nil {
		log.Fatal(err)
	}
	router.Run(":8081")
}