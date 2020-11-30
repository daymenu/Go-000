package dao

import (
	"database/sql"
	"errors"
	"log"
	"testing"
)

func TestFindMoiveByName(t *testing.T) {
	_, err := FindMoiveByName("yil")
	if !errors.Is(err, sql.ErrNoRows) {
		log.Println("没有mysql的行")
	}
	log.Println(errors.Unwrap(err))
	t.Fail()
}
