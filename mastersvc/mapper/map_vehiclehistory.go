package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Operator Contact table
//--------------------------------------
func MapVehicleHistoryTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleHistory{}) {
		db.CreateTable(&ent.TableVehicleHistory{})
		db.Model(&ent.TableVehicleHistory{}).AddForeignKey("vehicleid", "table_vehicles(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicleHistory{}).AddForeignKey("fromstatusid", "table_vehiclestatus(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicleHistory{}).AddForeignKey("tostatusid", "tostatusid(id)", "RESTRICT", "RESTRICT")
	}
}
