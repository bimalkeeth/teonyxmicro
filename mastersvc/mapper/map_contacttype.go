package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Contact Type table
//--------------------------------------
func MapContactTypeTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableAddressType{}) {

		db.CreateTable(&ent.TableContactType{})
		db.Model(&ent.TableContactType{}).AddUniqueIndex("ux_addresstype_contacttype", "contacttype")
		db.Model(&ent.TableContactType{}).Association("Contacts")
	}
}