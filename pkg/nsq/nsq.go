package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var Producer *nsq.Producer

type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

func InitProducer() error {
	p, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		return err
	}
	Producer = p

	//err = Producer.Publish("sex",[]byte("123"))
	//if err != nil {
	//	fmt.Printf("publish msg to nsq failed, err:%v\n", err)
	//	return err
	//}
	return nil
}


func NewConsumer(channel string) (*nsq.Consumer,error) {
	c,err:= nsq.NewConsumer("sex",channel,nsq.NewConfig())
	if err!=nil {
		return nil,err
	}

	//consumer := &MyHandler{
	//	Title: "zhenzhen",
	//}
	//c.AddHandler(consumer)

	//if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	//	return nil,err
	//}
	return c,nil
}


