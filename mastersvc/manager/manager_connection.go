package manager

import (
	"github.com/jinzhu/gorm"
	"teonyxmicro/mastersvc/connection"
)

//---------------------------------------
//Set database connection
//---------------------------------------
func Setconn() (*gorm.DB, error) {
	conn := connection.New()
	db, err := conn.Open()
	if err != nil {
		return db, err
	}
	return db, nil
}
