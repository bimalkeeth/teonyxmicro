package bucontracts

type AddressBO struct {
	Id            uint
	Address       string
	Street        string
	Suburb        string
	StateId       uint
	CountryId     uint
	AddressTypeId uint
	Location      string
	AddressType   AddressTypeBO
	State         StateBO
	Country       CountryBO
}
