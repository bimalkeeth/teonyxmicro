package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create vehicle operator location table
//--------------------------------------
func MapVehicleTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicle{}) {

		db.CreateTable(&ent.TableVehicle{})
		db.Model(&ent.TableVehicle{}).AddUniqueIndex("ux_vehicle_registration", "registration")
		db.Model(&ent.TableVehicle{}).AddForeignKey("modelid", "table_vehiclemodel(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicle{}).AddForeignKey("makeid", "table_vehiclemake(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicle{}).AddForeignKey("fleetid", "table_fleet(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicle{}).AddForeignKey("statusid", "table_vehiclestatus(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicle{}).Association("Locations")
		db.Model(&ent.TableVehicle{}).Association("History")
		db.Model(&ent.TableVehicle{}).Association("Operators")
		db.Model(&ent.TableVehicle{}).Association("Registrations")
	}
}
