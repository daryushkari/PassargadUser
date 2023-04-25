package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = nil
)

func Get(path string) (err error, sDB *gorm.DB) {
	if db == nil {
		sDB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
		if err != nil {
			return err, nil
		}
		db = sDB
	}
	return nil, db
}
