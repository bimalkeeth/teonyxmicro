package mapper

import "github.com/jinzhu/gorm"
import ent "teonyxmicro/mastersvc/entities"

//--------------------------------------
// Create Operator Contact table
//--------------------------------------
func MapRegionTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableRegion{}) {

		db.CreateTable(&ent.TableRegion{})
		db.Model(&ent.TableState{}).AddUniqueIndex("ux_states_name", "name")
		db.Model(&ent.TableState{}).Association("Countries")
	}
}
