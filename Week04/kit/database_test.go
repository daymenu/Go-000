package kit

import (
	"testing"
)

func TestDB(t *testing.T) {
	config := new(Config)

	err := LoadConfig(config)
	if err != nil {
		t.Error(err)
	}
	db, err := NewDB(&config.Mysql)
	if err != nil {
		t.Error(err)
	}
	err = db.Ping()

	if err != nil {
		t.Error(err)
	}
}
