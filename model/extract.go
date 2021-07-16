package model

import (
	"FileManage/pkg/db"
	"database/sql"
)

func ReadDataFromDB(querySql string) (*sql.Rows, error) {
	rows,err:= db.DB.Query(querySql)
	if err!= nil {
		return nil,err
	}
	return rows,nil
}
