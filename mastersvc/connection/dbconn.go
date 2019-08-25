package connection

import (
	"github.com/jinzhu/gorm"
	"log"
)

type IDatabase interface {
	Open() (*gorm.DB, error)
}

type DB struct{}

func New() *DB {
	return &DB{}
}

//-----------------------------------
// open connection
//-----------------------------------
func (db DB) Open() (*gorm.DB, error) {
	pdb, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=root sslmode=disable")
	if err != nil {
		log.Fatal("error in connecting to the database")
		return nil, err
	}
	return pdb, nil
}
