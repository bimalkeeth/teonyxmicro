package mapper

import (
	"log"
	con "teonyxmicro/mastersvc/connection"
	ent "teonyxmicro/mastersvc/entities"
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
	dbase, err := db.Open()
	if err != nil {
		log.Fatal("error in connection")
	}
	if !dbase.HasTable(&ent.TableAddressType{}) {

		dbase.CreateTable(&ent.TableAddressType{})
	}
	if !dbase.HasTable(&ent.TableRegion{}) {

		dbase.CreateTable(&ent.TableRegion{})
	}
	if !dbase.HasTable(&ent.TableState{}) {

		dbase.CreateTable(&ent.TableState{})
	}
	if !dbase.HasTable(&ent.TableCountry{}) {

		dbase.CreateTable(&ent.TableCountry{})
	}
	if !dbase.HasTable(&ent.TableAddress{}) {

		dbase.CreateTable(&ent.TableAddress{})
	}
	return nil
}
