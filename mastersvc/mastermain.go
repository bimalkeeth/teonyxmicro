package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"teonyxmicro/mastersvc/mapper"
)

func main() {

	mapp := mapper.New()
	err := mapp.GenerateSchema()
	if err != nil {
		log.Fatal("Error")
	}
}
