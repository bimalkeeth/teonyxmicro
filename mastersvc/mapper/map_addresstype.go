package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Address Type table
//--------------------------------------
func MapAddressTypeTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableAddressType{}) {

		db.CreateTable(&ent.TableAddressType{})
		db.Model(&ent.TableAddressType{}).AddUniqueIndex("ux_addresstype_name", "name")

	}
}
