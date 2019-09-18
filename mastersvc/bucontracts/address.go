package bucontracts

type AddressBO struct {
	Address       string
	Street        string
	Suburb        string
	StateId       uint
	CountryId     uint
	AddressTypeId int
	Location      string
	AddressType   AddressTypeBO
	State         StateBO
	Country       CountryBO
}
