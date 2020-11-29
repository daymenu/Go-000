package service

import (
	"github.com/daymenu/Go-000/Week02/work/dao"
	"github.com/daymenu/Go-000/Week02/work/model"
)

// FindMoiveByName 根据影片名搜索影片
func FindMoiveByName(name string) (*model.Moive, error) {
	moive, err := dao.FindMoiveByName(name)
	if err != nil {
		return nil, err
	}
	return moive, nil
}

// FindMostLoveThreeMoives 查找最喜爱的3部影片
func FindMostLoveThreeMoives() ([]*model.Moive, error) {
	loveMoives, err := dao.FindMostLoveMoives(3)
	if err != nil {
		return nil, err
	}
	return loveMoives, err
}
