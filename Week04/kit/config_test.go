package kit

import "testing"

func TestLoadConfig(t *testing.T) {
	db := new(Config)

	err := LoadConfig(db)
	t.Error(err)
	t.Error(db)
}
