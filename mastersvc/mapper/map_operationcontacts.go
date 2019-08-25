package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Fleet Location table
//--------------------------------------
func MapOperationContactsTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableOperatorContacts{}) {

		db.CreateTable(&ent.TableOperatorContacts{})
		db.Model(&ent.TableOperatorContacts{}).AddUniqueIndex("ux_operatorcontacts_operationidcontactid", "operatorid", "contactid")
		db.Model(&ent.TableOperatorContacts{}).AddForeignKey("contactid", "table_contacts(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableOperatorContacts{}).AddForeignKey("operatorid", "table_vehicleoperators(id)", "RESTRICT", "RESTRICT")
	}
}
