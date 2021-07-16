package model

import (
	"FileManage/pkg/db"
	"strconv"
)

func ChangeTaskStatus(lastId int64) error {
	sql:="update task set status=1 where id="+strconv.FormatInt(lastId,10)
	_,err:=db.DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func AddTask(task string) (int64,error) {
	sql := "insert into task(task) values(" + task + ")"
	res, err := db.DB.Exec(sql)
	if err != nil {
		return 0,err
	}
	lastId,err:=res.LastInsertId()
	if err != nil {
		return 0,err
	}
	return lastId,nil
}
