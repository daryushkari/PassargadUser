package postgresql

import (
	"PassargadUser/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB = nil
)

func Get(cnf *config.EnvConfig) (err error, sDB *gorm.DB) {
	if db == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			cnf.Database.Host,
			cnf.Database.User,
			cnf.Database.Password,
			cnf.Database.DBName,
			cnf.Database.Port)
		log.Println(dsn)
		pdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err, nil
		}
		db = pdb
	}
	return nil, db
}
