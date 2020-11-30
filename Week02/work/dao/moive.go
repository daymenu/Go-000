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
		// 1. 将错误吞掉
		// return nil, nil
		// 2. 不wrap直接返回原始错误
		// return nil, err
		// 3. 使用wrap后返回, 最后选择使用第三种
		return nil, fmt.Errorf("没有找到 《%s》 这个影片: %w", name, err)
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
