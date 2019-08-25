package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Fleet table
//--------------------------------------
func MapFleetTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableFleet{}) {

		db.CreateTable(&ent.TableFleet{})
		db.Model(&ent.TableFleet{}).AddUniqueIndex("ux_fleet_fleetid", "fleetid")
		db.Model(&ent.TableFleet{}).Association("FleetContacts")
		db.Model(&ent.TableFleet{}).Association("FleetLocations")
	}
}
