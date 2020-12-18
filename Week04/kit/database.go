package kit

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// 初始化数据库
	_ "github.com/go-sql-driver/mysql"
)

type dbs struct {
	dbs map[string]*sql.DB
	sync.RWMutex
}

var mDBs = dbs{dbs: make(map[string]*sql.DB)}

// NewDB 创建一个数据库连接
func NewDB(c *DBConfig) (*sql.DB, error) {
	mDBs.RLock()

	if db, ok := mDBs.dbs[c.DBName]; ok {
		mDBs.RUnlock()
		return db, nil
	}
	mDBs.RUnlock()
	mDBs.Lock()
	defer mDBs.Unlock()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	log.Println("connect mysql ...")
	if err != nil {
		log.Printf("connect mysql[%s] is failed, err:%v", dataSourceName, err)
	}
	mDBs.dbs[c.DBName] = db
	return db, err
}
