package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create vehicle make table
//--------------------------------------
func MapVehicleMakeTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleMake{}) {
		db.CreateTable(&ent.TableVehicleMake{})
		db.Model(&ent.TableVehicleMake{}).AddUniqueIndex("ux_vehiclemake_make", "make")
		db.Model(&ent.TableVehicleMake{}).AddForeignKey("countryid", "table_contacttype(id)", "RESTRICT", "RESTRICT")

	}
}
