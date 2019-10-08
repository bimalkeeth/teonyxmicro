package bucontracts

type OperatorContactsBO struct {
	Id         uint
	ContactId  uint
	OperatorId uint
	Contact    ContactBO
}
