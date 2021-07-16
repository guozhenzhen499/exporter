package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {
	// single file
	//file, err := c.FormFile("file")
	//if err!= nil {
	//	log.Fatal(err)
	//}
	//log.Println(file.Filename)
	//
	//// Upload the file to specific dst.
	//err=c.SaveUploadedFile(file, "D:\\APROJECT\\GO-Project\\src\\FileManage\\file\\"+file.Filename)
	//if err!=nil {
	//	log.Println(err)
	//}
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	//multiple file
	form, err := c.MultipartForm()
	if err != nil {
		log.Fatal(err)
	}

	files := form.File["file"]

	for _, file := range files {
		log.Println(file.Filename)
		f, err := os.Stat("D:\\APROJECT\\GO-Project\\src\\FileManage\\file\\" + file.Filename)
		if err != nil {
			log.Fatal(err)
		}

		if f == nil {
			//fmt.Println("file exist")
			err = c.SaveUploadedFile(file, "D:\\APROJECT\\GO-Project\\src\\FileManage\\file\\"+file.Filename)
			if err != nil {
				log.Println(err)
			}

		} else {
			num := time.Now().UnixNano()
			err = c.SaveUploadedFile(file, "D:\\APROJECT\\GO-Project\\src\\FileManage\\file\\"+strconv.Itoa(int(num))+file.Filename)
			if err != nil {
				log.Println(err)
			}
		}

	}
	c.String(http.StatusOK, fmt.Sprintf("'%d' uploaded!", len(files)))
}
