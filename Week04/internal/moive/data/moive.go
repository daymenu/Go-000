package data

// Moive moive
type Moive struct {
	name string
}

// Search search
func (m *Moive) Search() ([]*Moive, error) {
	res := []*Moive{}
	db, err := GetDB()
	if err != nil {

	}
	rows, err := db.Query("SELECT * FROM moive WHERE name = '%s'", m.name)
	if err != nil {

	}
	defer rows.Close()

	return res, rows.Err()
}
