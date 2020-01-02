package protoapi

import (
	"context"
	"encoding/json"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import mst "teonyxmicro/mastersvc/manager/masters"
import bu "teonyxmicro/mastersvc/bucontracts"

type MasterService struct{}

//------------------------------------------
//Create Company
//------------------------------------------
func (m *MasterService) CreateCompany(ctx context.Context, req *pro.RequestCompany) (rsp *pro.ResponseCreateSuccess, err error) {
	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Name:      req.Company.Name,
		ContactId: uint(req.Company.ContactId),
		AddressId: uint(req.Company.AddressId),
	}
	id, err := masterManager.CreateCompany(companyRequest)
	response := &pro.ResponseCreateSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Id = uint64(id)
	return response, nil
}
