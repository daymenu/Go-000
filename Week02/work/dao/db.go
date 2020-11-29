package dao

import (
	"database/sql"
	"log"
	"sync"
)

type moiveDBs struct {
	dbs map[string]*sql.DB
	sync.RWMutex
}

var mDBs = moiveDBs{dbs: make(map[string]*sql.DB)}

// NewDB 创建一个数据库连接
func NewDB(dbname string) (*sql.DB, error) {
	mDBs.RLock()
	if db, ok := mDBs.dbs[dbname]; ok {
		mDBs.RUnlock()
		return db, nil
	}
	mDBs.RUnlock()
	mDBs.Lock()
	defer mDBs.Unlock()
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/"+dbname)
	log.Println("connect mysql ...")
	if err != nil {
		log.Printf("connect mysql[%s] is failed", "127.0.0.1")
	}
	mDBs.dbs[dbname] = db
	return db, err
}
