package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Operator Contact table
//--------------------------------------
func MapVehicleTrackRegTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleTrackReg{}) {
		db.CreateTable(&ent.TableVehicleTrackReg{})
		db.Model(&ent.TableVehicleTrackReg{}).AddForeignKey("vehicleid", "table_vehicles(id)", "RESTRICT", "RESTRICT")
	}
}
