package main

import "teonyxmicro/mastersvc/mapper"
import "log"

func main() {

	mapp := mapper.New()
	err := mapp.GenerateSchema()
	if err != nil {
		log.Fatal("Error")
	}

}
