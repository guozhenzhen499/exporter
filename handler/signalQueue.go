package handler

import (
	"FileManage/pkg/nsq"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SingnalQueue(c *gin.Context) {
	sex:=c.Request.URL.Query().Get("sex")
	err:=nsq.Producer.Publish("sex",[]byte(sex))
	if err!= nil {
		fmt.Println(err.Error())
		c.String(http.StatusOK,fmt.Sprintf("导出reader表失败"))
	}
	//cache.Cache.RPush("condition",sex)
	c.String(http.StatusOK, fmt.Sprintln( "后台正在导出reader表"))
}
