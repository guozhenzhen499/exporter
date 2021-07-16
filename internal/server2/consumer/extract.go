package consumer

import (
	"FileManage/model"
	n "FileManage/pkg/nsq"
	"github.com/nsqio/go-nsq"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	"time"
)

type Reader struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	IDCard     string `json:"id_card"`
	ReaderCard int    `json:"reader_card"`
	Status     int    `json:"status"`
}

type ExtractHandler struct {
}

func Extract(channel string) {
	c,err := n.NewConsumer(channel)
	if err!=nil {
		log.Fatal(err)
	}
	e := &ExtractHandler{}
	c.AddHandler(e)
	if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		log.Fatal(err)
	}
}

func(e *ExtractHandler) HandleMessage(msg *nsq.Message) error {

	m:=string(msg.Body)

	//l, err := cache.Cache.LLen("condition").Result()
	//if err != nil {
	//	return err
	//}
	//task, err := cache.Cache.LRange("condition", 0, l-1).Result()
	//for i := int64(0); i < l;i++ {
	//	cache.Cache.LPop("condition")
	//}
	//if err != nil {
	//	return err
	//}
	//for _, v := range task {
		lastId, err := model.AddTask(m)
		if err != nil {
			return err
		}
		sql := "select * from reader where sex=" + m
		rows, err := model.ReadDataFromDB(sql)
		if err != nil {
			return err
		}
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("sheet1")
		if err != nil {
			return err
		}
		title := []string{"ID", "Name", "Sex", "IDCard", "ReaderCard", "Status"}
		titleRow := sheet.AddRow()
		for _, v := range title {
			cell := titleRow.AddCell()
			cell.Value = v
		}

		reader := Reader{}
		for rows.Next() {
			_ = rows.Scan(&reader.Id, &reader.Name, &reader.Sex, &reader.IDCard, &reader.ReaderCard, &reader.Status)
			row := sheet.AddRow()
			row.WriteStruct(&reader, -1)
		}
		num := time.Now().UnixNano()
		err = file.Save("D:\\APROJECT\\GO-Project\\src\\FileManage\\file\\" + strconv.FormatInt(num, 10) + "reader.xlsx")
		if err != nil {
			return err
		}
		err = model.ChangeTaskStatus(lastId)
		if err != nil {
			return err
		}
    //}

	return nil
	//.xlsx返回前端
	//var buffer bytes.Buffer
	//_ = file.Write(&buffer)
	//content := bytes.NewReader(buffer.Bytes())
	//fileName := fmt.Sprintf("%s.xlsx","reader")
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	//c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)
}
