package biz

import (
	"github.com/daymenu/Go-000/Week02/work/model"
	"github.com/daymenu/Go-000/Week02/work/service"
)

func FindMoive(name string) (*model.SearchMoive, error) {
	var searchMoive model.SearchMoive
	moive, err := service.FindMoiveByName(name)
	if err != nil {
		return nil, err
	}
	searchMoive.Moive = moive
	return &searchMoive, nil
}
