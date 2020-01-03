package protoapi

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import mst "teonyxmicro/mastersvc/manager/masters"
import bu "teonyxmicro/mastersvc/bucontracts"

//------------------------------------------
//Create Company
//------------------------------------------
func (m *MasterService) CreateCompany(ctx context.Context, in *pro.RequestCompany) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Name:      in.Company.Name,
		ContactId: uint(in.Company.ContactId),
		AddressId: uint(in.Company.AddressId),
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

//--------------------------------------
//Update Company
//--------------------------------------
func (m *MasterService) UpdateCompany(ctx context.Context, in *pro.RequestCompany) (*pro.ResponseSuccess, error) {

	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Id:        uint(in.Company.Id),
		Name:      in.Company.Name,
		ContactId: uint(in.Company.ContactId),
		AddressId: uint(in.Company.AddressId),
	}
	success, err := masterManager.UpdateCompany(companyRequest)
	response := &pro.ResponseSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Success = success
	return response, nil
}

//--------------------------------------
//Delete Company
//--------------------------------------
func (m *MasterService) DeleteCompany(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.DeleteCompany(uint(in.Id))
	response := &pro.ResponseSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Success = result
	return response, nil
}

//---------------------------------------
//Create Address Type
//=======================================
func (m *MasterService) CreateAddressType(ctx context.Context, in *pro.RequestAddressType) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()

	result, err := masterManager.CreateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})
	response := &pro.ResponseCreateSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Id = uint64(result)
	return response, nil
}

//---------------------------------------
//Update Address Type
//---------------------------------------
func (m *MasterService) UpdateAddressType(ctx context.Context, in *pro.RequestAddressType) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.UpdateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})
	response := &pro.ResponseSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Success = result
	return response, nil
}

//---------------------------------------
//Delete Address Type
//---------------------------------------
func (m *MasterService) DeleteAddressType(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.DeleteAddressType(uint(in.Id))
	response := &pro.ResponseSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Success = result
	return response, nil
}

//---------------------------------------
//Get Address Type by Given Id
//---------------------------------------
func (m *MasterService) GetAddressTypeById(ctx context.Context, in *pro.RequestKey) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeById(uint(in.Id))
	response := &pro.ResponseAddressType{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	response.AddressType = addressType
	return response, nil
}

//--------------------------------------
//Get Address Type By Name
//--------------------------------------
func (m *MasterService) GetAddressTypeByName(ctx context.Context, in *pro.RequestByName) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeByName(in.Name)
	response := &pro.ResponseAddressType{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	response.AddressType = addressType
	return response, nil
}

//----------------------------------
//Get All Address Types
//----------------------------------
func (m *MasterService) GetAllAddressTypes(ctx context.Context, in *empty.Empty) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypes()
	response := &pro.ResponseAddressType{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   uint64(item.Id),
			Name: item.Name,
		})
	}
	response.AddressType = addressType
	return response, nil
}

//--------------------------------
//Get all address type by name part
//--------------------------------
func (m *MasterService) GetAllAddressTypeNames(ctx context.Context, in *pro.RequestByName) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypeNames(in.Name)
	response := &pro.ResponseAddressType{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   uint64(item.Id),
			Name: item.Name,
		})
	}
	response.AddressType = addressType
	return response, nil
}

//----------------------------------
//Create Region
//----------------------------------
func (m *MasterService) CreateRegion(ctx context.Context, in *pro.RequestRegion) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.CreateRegion(bu.RegionBO{Region: in.Region.Region})
	response := &pro.ResponseCreateSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Id = uint64(result)
	return response, nil
}

//-----------------------------------
//Update Region
//-----------------------------------
func (m *MasterService) UpdateRegion(ctx context.Context, in *pro.RequestRegion) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.UpdateRegion(bu.RegionBO{Id: uint(in.Region.Id), Region: in.Region.Region})
	response := &pro.ResponseSuccess{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		response.ErrorMessage = jsonError
		response.IsError = true
	}
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteRegion(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {

	response := &pro.ResponseSuccess{}
	return response, nil
}
