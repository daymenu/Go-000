package dao

import (
	"fmt"

	"github.com/daymenu/Go-000/Week02/work/model"

	// 初始化数据库
	_ "github.com/go-sql-driver/mysql"
)

var dbName = "moive"

// FindMoiveByName 根据影片名字查找影片
func FindMoiveByName(name string) (*model.Moive, error) {
	db, err := NewDB(dbName)
	if err != nil {
		return nil, err
	}

	moiveSQL := "SELECT * FROM moive WHERE `name`= ?"
	moive := &model.Moive{}
	if err := db.QueryRow(moiveSQL, name).Scan(&moive.Id, &moive.Name, &moive.ToStar); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return moive, nil
}

// FindMostLoveMoives 查找最喜欢的电影
func FindMostLoveMoives(number int) ([]*model.Moive, error) {
	db, err := NewDB(dbName)
	if err != nil {
		return nil, err
	}

	moiveSQL := fmt.Sprintf("SELECT * FROM moive ORDER BY love DESC LIMIT %d", number)
	rows, err := db.Query(moiveSQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	loveMoives := make([]*model.Moive, 3)
	for rows.Next() {
		moive := &model.Moive{}
		rows.Scan(&moive.Id, &moive.Name, &moive.ToStar)
		loveMoives = append(loveMoives, moive)
	}

	return loveMoives, nil
}
