package protoapi

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import mst "teonyxmicro/mastersvc/manager/masters"
import bu "teonyxmicro/mastersvc/bucontracts"
import con "teonyxmicro/mastersvc/manager/contacts"
import timestamp "github.com/golang/protobuf/ptypes"

func (m *MasterService) CreateCompany(ctx context.Context, in *pro.RequestCompany) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Name:      in.Company.Name,
		ContactId: uint(in.Company.ContactId),
		AddressId: uint(in.Company.AddressId),
	}
	id, err := masterManager.CreateCompany(companyRequest)
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)

	response.Id = uint64(id)
	return response, nil
}

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
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = success
	return response, nil
}

func (m *MasterService) DeleteCompany(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.DeleteCompany(uint(in.Id))
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) CreateAddressType(ctx context.Context, in *pro.RequestAddressType) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()

	result, err := masterManager.CreateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	return response, nil
}

func (m *MasterService) UpdateAddressType(ctx context.Context, in *pro.RequestAddressType) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.UpdateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteAddressType(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.DeleteAddressType(uint(in.Id))
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) GetAddressTypeById(ctx context.Context, in *pro.RequestKey) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeById(uint(in.Id))
	response := &pro.ResponseAddressType{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	response.AddressType = addressType
	return response, nil
}

func (m *MasterService) GetAddressTypeByName(ctx context.Context, in *pro.RequestByName) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeByName(in.Name)
	response := &pro.ResponseAddressType{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	response.AddressType = addressType
	return response, nil
}

func (m *MasterService) GetAllAddressTypes(ctx context.Context, in *empty.Empty) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypes()
	response := &pro.ResponseAddressType{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
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

func (m *MasterService) GetAllAddressTypeNames(ctx context.Context, in *pro.RequestByName) (*pro.ResponseAddressType, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypeNames(in.Name)
	response := &pro.ResponseAddressType{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
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

func (m *MasterService) CreateRegion(ctx context.Context, in *pro.RequestRegion) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.CreateRegion(bu.RegionBO{Region: in.Region.Region})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	return response, nil
}

func (m *MasterService) UpdateRegion(ctx context.Context, in *pro.RequestRegion) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.UpdateRegion(bu.RegionBO{Id: uint(in.Region.Id), Region: in.Region.Region})
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteRegion(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {

	masterManager := mst.New()
	result, err := masterManager.DeleteRegion(uint(in.Id))
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) GetAllRegion(ctx context.Context, in *empty.Empty) (*pro.ResponseRegion, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAllRegion()
	response := &pro.ResponseRegion{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, item := range result {
		response.Region = append(response.Region, &pro.RegionProto{Region: item.Region,
			RegionName: item.RegionName,
			Id:         uint64(item.Id),
		})
	}
	return response, nil
}

func (m *MasterService) GetRegionById(ctx context.Context, in *pro.RequestKey) (*pro.ResponseRegion, error) {
	masterManager := mst.New()
	result, err := masterManager.GetRegionById(uint(in.Id))
	response := &pro.ResponseRegion{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Region = append(response.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: uint64(result.Id)})
	return response, nil
}

func (m *MasterService) GetRegionByName(ctx context.Context, in *pro.RequestByName) (*pro.ResponseRegion, error) {
	masterManager := mst.New()
	result, err := masterManager.GetRegionByName(in.Name)
	response := &pro.ResponseRegion{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Region = append(response.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: uint64(result.Id)})
	return response, nil

}

func (m *MasterService) CreateState(ctx context.Context, in *pro.RequestState) (*pro.ResponseCreateSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.CreateState(bu.StateBO{Name: in.State.Name,
		CountryId: uint(in.State.CountryId)})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	return response, nil
}

func (m *MasterService) UpdateState(ctx context.Context, in *pro.RequestState) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.UpdateState(bu.StateBO{Name: in.State.Name,
		CountryId: uint(in.State.CountryId), Id: uint(in.State.Id)})
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteState(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	masterManager := mst.New()
	result, err := masterManager.DeleteState(uint(in.Id))
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) GetStateById(ctx context.Context, in *pro.RequestKey) (*pro.ResponseState, error) {
	masterManager := mst.New()
	result, err := masterManager.GetStateById(uint(in.Id))
	response := &pro.ResponseState{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.State = append(response.State, &pro.StateProto{Id: uint64(result.Id),
		CountryId: uint64(result.CountryId), Name: result.Name})
	return response, nil
}

func (m *MasterService) GetStateByCountryId(ctx context.Context, in *pro.RequestKey) (*pro.ResponseState, error) {
	masterManager := mst.New()
	result, err := masterManager.GetStateByCountryId(uint(in.Id))
	response := &pro.ResponseState{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		response.State = append(response.State, &pro.StateProto{
			Id:        uint64(item.Id),
			CountryId: uint64(item.CountryId),
			Name:      item.Name,
		})
	}
	return response, nil
}

func (m *MasterService) GetStateByName(ctx context.Context, in *pro.RequestByName) (*pro.ResponseState, error) {
	masterManager := mst.New()
	result, err := masterManager.GetStateByName(in.Name)
	response := &pro.ResponseState{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.State = append(response.State, &pro.StateProto{Id: uint64(result.Id),
		CountryId: uint64(result.CountryId), Name: result.Name})
	return response, nil
}

func (m *MasterService) GetAllStates(ctx context.Context, in *empty.Empty) (*pro.ResponseState, error) {
	masterManager := mst.New()
	result, err := masterManager.GetAllStates()
	response := &pro.ResponseState{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		response.State = append(response.State, &pro.StateProto{
			Id:        uint64(item.Id),
			CountryId: uint64(item.CountryId),
			Name:      item.Name,
		})
	}
	return response, nil
}

func (m *MasterService) CreateContact(ctx context.Context, in *pro.RequestContact) (*pro.ResponseCreateSuccess, error) {
	conManager := con.New()
	result, err := conManager.CreateContact(bu.ContactBO{Contact: in.Contact.Contact,
		ContactTypeId: uint(in.Contact.ContactTypeId),
	})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	return response, nil
}

func (m *MasterService) UpdateContact(ctx context.Context, in *pro.RequestContact) (*pro.ResponseSuccess, error) {
	conManager := con.New()
	result, err := conManager.UpdateContact(bu.ContactBO{Id: uint(in.Contact.Id),
		Contact:       in.Contact.Contact,
		ContactTypeId: uint(in.Contact.ContactTypeId),
	})
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteContact(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	conManager := con.New()
	result, err := conManager.DeleteContact(uint(in.Id))
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) ContactById(ctx context.Context, in *pro.RequestKey) (*pro.ResponseContact, error) {
	conManager := con.New()
	result, err := conManager.ContactById(uint(in.Id))
	response := &pro.ResponseContact{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, err := timestamp.TimestampProto(result.UpdatedAt)
	if err != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}
	response.Contact = &pro.ContactProto{
		Id:            uint64(result.Id),
		Contact:       result.Contact,
		ContactTypeId: uint64(result.ContactTypeId),
		UpdatedAt:     timeUpdate,
	}
	return response, nil
}

func (m *MasterService) CreateAddress(ctx context.Context, in *pro.RequestAddress) (*pro.ResponseCreateSuccess, error) {
	conManager := con.New()
	result, err := conManager.CreateAddress(bu.AddressBO{
		CountryId:     uint(in.Address.CountryId),
		AddressTypeId: uint(in.Address.AddressTypeId),
		Location:      in.Address.Location,
		Address:       in.Address.Address,
		StateId:       uint(in.Address.StateId),
		Street:        in.Address.Street,
		Suburb:        in.Address.Suburb,
	})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	response.Success = true
	return response, err
}

func (m *MasterService) UpdateAddress(ctx context.Context, in *pro.RequestAddress) (*pro.ResponseSuccess, error) {
	conManager := con.New()
	result, err := conManager.UpdateAddress(bu.AddressBO{
		Id:            uint(in.Address.Id),
		CountryId:     uint(in.Address.CountryId),
		AddressTypeId: uint(in.Address.AddressTypeId),
		Location:      in.Address.Location,
		Address:       in.Address.Address,
		StateId:       uint(in.Address.StateId),
		Street:        in.Address.Street,
		Suburb:        in.Address.Suburb,
	})
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	response.Success = true
	return response, err
}

func (m *MasterService) DeleteAddress(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	conManager := con.New()
	result, err := conManager.DeleteAddress(uint(in.Id))
	response := &pro.ResponseSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, err
}

func (m *MasterService) GetAddressById(ctx context.Context, in *pro.RequestKey) (*pro.ResponseAddress, error) {
	conManager := con.New()
	result, err := conManager.GetAddressById(uint(in.Id))
	response := &pro.ResponseAddress{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, err := timestamp.TimestampProto(result.UpdatedAt)
	if err != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}

	response.Address = append(response.Address, &pro.AddressProto{
		Address:       result.Address,
		Id:            uint64(result.Id),
		Suburb:        result.Suburb,
		Street:        result.Street,
		StateId:       uint64(result.StateId),
		Location:      result.Location,
		AddressTypeId: uint64(result.AddressTypeId),
		CountryId:     uint64(result.CountryId),
		UpdatedAt:     timeUpdate,
	})
	return response, nil
}
