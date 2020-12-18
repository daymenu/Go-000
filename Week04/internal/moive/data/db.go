package data

import (
	"database/sql"

	"github.com/daymenu/Go-000/Week04/kit"
)

// GetDB 获取数据库实例
func GetDB() (*sql.DB, error) {
	config := kit.Config{}
	err := kit.LoadConfig(&config)
	if err != nil {
		return nil, err
	}
	return kit.NewDB(&config.Mysql)
}

// Dataer dataer
type Dataer interface {
	Search()
}
