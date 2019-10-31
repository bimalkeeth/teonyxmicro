package manager

import bu "teonyxmicro/mastersvc/bucontracts"

type IMasterService interface {
	CreateCompany(bo bu.CompanyBO)
}
