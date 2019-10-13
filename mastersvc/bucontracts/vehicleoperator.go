package bucontracts

type VehicleOperatorBO struct {
	Id        uint
	Name      string
	SurName   string
	Locations []*OperatorLocationBO
	Contacts  []*OperatorContactsBO
}
