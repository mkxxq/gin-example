package models

import (
	"github.com/jmoiron/sqlx"
)
type User struct {
	ID int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}


func(obj *User) InsertUser(tx *sqlx.DB)(int64 ,error){
	sql := "INSERT INTO t_user(name) VALUES(?)"
	res, err := tx.Exec(sql, obj.Name)
	if err != nil{
		return 0, err
	}
	return res.RowsAffected()
}