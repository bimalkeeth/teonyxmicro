package mapper

import (
	"log"
	con "teonyxmicro/mastersvc/connection"
)

type IEntityMapper interface {
	GenerateSchema() error
}

type SchemaGenerator struct{}

//-------------------------------------
// Create instance
//-------------------------------------
func New() SchemaGenerator {

	return SchemaGenerator{}
}

//--------------------------------------
// Create mapping
//--------------------------------------
func (t SchemaGenerator) GenerateSchema() error {
	db := con.New()
	database, err := db.Open()
	if err != nil {
		log.Fatal("error in connection")
	}
	MapAddressTypeTable(database)
	MapContactTypeTable(database)
	MapRegionTable(database)
	MapCountryTable(database)
	MapStatesTable(database)

	return nil
}
