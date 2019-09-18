package bucontracts

type CountryBO struct {
	Id          uint
	CountryName string
	RegionId    uint
	Region      RegionBO
	States      []RegionBO
}
