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

func (m *MasterService) CreateCompany(ctx context.Context, req *pro.RequestCompany, res *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Name:      req.Company.Name,
		ContactId: uint(req.Company.ContactId),
		AddressId: uint(req.Company.AddressId),
	}
	id, err := masterManager.CreateCompany(companyRequest)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(id)
	return nil
}

func (m *MasterService) UpdateCompany(ctx context.Context, req *pro.RequestCompany, res *pro.ResponseSuccess) error {

	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Id:        uint(req.Company.Id),
		Name:      req.Company.Name,
		ContactId: uint(req.Company.ContactId),
		AddressId: uint(req.Company.AddressId),
	}
	success, err := masterManager.UpdateCompany(companyRequest)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = success
	return nil
}

func (m *MasterService) DeleteCompany(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.DeleteCompany(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) CreateAddressType(ctx context.Context, req *pro.RequestAddressType, res *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()

	result, err := masterManager.CreateAddressType(bu.AddressTypeBO{
		Name: req.AddressType.Name,
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateAddressType(ctx context.Context, req *pro.RequestAddressType, res *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.UpdateAddressType(bu.AddressTypeBO{
		Name: req.AddressType.Name,
	})

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteAddressType(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.DeleteAddressType(uint(req.Id))

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) GetAddressTypeById(ctx context.Context, req *pro.RequestKey, res *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeById(uint(req.Id))

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	res.AddressType = addressType
	return nil
}

func (m *MasterService) GetAddressTypeByName(ctx context.Context, req *pro.RequestByName, res *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeByName(req.Name)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	res.AddressType = addressType
	return nil
}

func (m *MasterService) GetAllAddressTypes(ctx context.Context, req *empty.Empty, res *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypes()
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   uint64(item.Id),
			Name: item.Name,
		})
	}
	res.AddressType = addressType
	return nil
}

func (m *MasterService) GetAllAddressTypeNames(ctx context.Context, req *pro.RequestByName, res *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypeNames(req.Name)

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   uint64(item.Id),
			Name: item.Name,
		})
	}
	res.AddressType = addressType
	return nil
}

func (m *MasterService) CreateRegion(ctx context.Context, req *pro.RequestRegion, res *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.CreateRegion(bu.RegionBO{Region: req.Region.Region})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateRegion(ctx context.Context, req *pro.RequestRegion, res *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.UpdateRegion(bu.RegionBO{Id: uint(req.Region.Id), Region: req.Region.Region})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteRegion(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {

	masterManager := mst.New()
	result, err := masterManager.DeleteRegion(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) GetAllRegion(ctx context.Context, req *empty.Empty, res *pro.ResponseRegion) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllRegion()

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		res.Region = append(res.Region, &pro.RegionProto{Region: item.Region,
			RegionName: item.RegionName,
			Id:         uint64(item.Id),
		})
	}
	return nil
}

func (m *MasterService) GetRegionById(ctx context.Context, req *pro.RequestKey, res *pro.ResponseRegion) error {
	masterManager := mst.New()
	result, err := masterManager.GetRegionById(uint(req.Id))

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Region = append(res.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: uint64(result.Id)})
	return nil
}

func (m *MasterService) GetRegionByName(ctx context.Context, req *pro.RequestByName, res *pro.ResponseRegion) error {
	masterManager := mst.New()
	result, err := masterManager.GetRegionByName(req.Name)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Region = append(res.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: uint64(result.Id)})
	return nil

}

func (m *MasterService) CreateState(ctx context.Context, req *pro.RequestState, res *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.CreateState(bu.StateBO{Name: req.State.Name,
		CountryId: uint(req.State.CountryId)})

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateState(ctx context.Context, req *pro.RequestState, res *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.UpdateState(bu.StateBO{Name: req.State.Name,
		CountryId: uint(req.State.CountryId), Id: uint(req.State.Id)})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteState(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.DeleteState(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) GetStateById(ctx context.Context, req *pro.RequestKey, res *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetStateById(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.State = append(res.State, &pro.StateProto{Id: uint64(result.Id),
		CountryId: uint64(result.CountryId), Name: result.Name})
	return nil
}

func (m *MasterService) GetStateByCountryId(ctx context.Context, req *pro.RequestKey, res *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetStateByCountryId(uint(req.Id))

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		res.State = append(res.State, &pro.StateProto{
			Id:        uint64(item.Id),
			CountryId: uint64(item.CountryId),
			Name:      item.Name,
		})
	}
	return nil
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

func (m *MasterService) GetAllStates(ctx context.Context, req *empty.Empty, res *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllStates()
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		res.State = append(res.State, &pro.StateProto{
			Id:        uint64(item.Id),
			CountryId: uint64(item.CountryId),
			Name:      item.Name,
		})
	}
	return nil
}

func (m *MasterService) CreateContact(ctx context.Context, req *pro.RequestContact, res *pro.ResponseCreateSuccess) error {
	conManager := con.New()
	result, err := conManager.CreateContact(bu.ContactBO{Contact: req.Contact.Contact,
		ContactTypeId: uint(req.Contact.ContactTypeId),
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateContact(ctx context.Context, req *pro.RequestContact, res *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.UpdateContact(bu.ContactBO{Id: uint(req.Contact.Id),
		Contact:       req.Contact.Contact,
		ContactTypeId: uint(req.Contact.ContactTypeId),
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteContact(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.DeleteContact(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) ContactById(ctx context.Context, req *pro.RequestKey, res *pro.ResponseContact) error {
	conManager := con.New()
	result, err := conManager.ContactById(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, err := timestamp.TimestampProto(result.UpdatedAt)
	if err != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}
	res.Contact = &pro.ContactProto{
		Id:            uint64(result.Id),
		Contact:       result.Contact,
		ContactTypeId: uint64(result.ContactTypeId),
		UpdatedAt:     timeUpdate,
	}
	return nil
}

func (m *MasterService) CreateAddress(ctx context.Context, req *pro.RequestAddress, res *pro.ResponseCreateSuccess) error {
	conManager := con.New()
	result, err := conManager.CreateAddress(bu.AddressBO{
		CountryId:     uint(req.Address.CountryId),
		AddressTypeId: uint(req.Address.AddressTypeId),
		Location:      req.Address.Location,
		Address:       req.Address.Address,
		StateId:       uint(req.Address.StateId),
		Street:        req.Address.Street,
		Suburb:        req.Address.Suburb,
	})

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	res.Success = true
	return err
}

func (m *MasterService) UpdateAddress(ctx context.Context, req *pro.RequestAddress, res *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.UpdateAddress(bu.AddressBO{
		Id:            uint(req.Address.Id),
		CountryId:     uint(req.Address.CountryId),
		AddressTypeId: uint(req.Address.AddressTypeId),
		Location:      req.Address.Location,
		Address:       req.Address.Address,
		StateId:       uint(req.Address.StateId),
		Street:        req.Address.Street,
		Suburb:        req.Address.Suburb,
	})

	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	res.Success = true
	return err
}

func (m *MasterService) DeleteAddress(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.DeleteAddress(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return err
}

func (m *MasterService) GetAddressById(ctx context.Context, req *pro.RequestKey, res *pro.ResponseAddress) error {
	conManager := con.New()
	result, err := conManager.GetAddressById(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, errs := timestamp.TimestampProto(result.UpdatedAt)
	if errs != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}

	res.Address = append(res.Address, &pro.AddressProto{
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
	return nil
}

func (m *MasterService) GetAddressByName(ctx context.Context, req *pro.RequestByName, res *pro.ResponseAddress) error {
	conManager := con.New()
	result, err := conManager.GetAddressByName(req.Name)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, item := range result {
		timeUpdate, errs := timestamp.TimestampProto(item.UpdatedAt)
		if errs != nil {
			timeUpdate, _ = timestamp.TimestampProto(time.Now())
		}
		res.Address = append(res.Address, &pro.AddressProto{
			Address:       item.Address,
			Id:            uint64(item.Id),
			Suburb:        item.Suburb,
			Street:        item.Street,
			StateId:       uint64(item.StateId),
			Location:      item.Location,
			AddressTypeId: uint64(item.AddressTypeId),
			CountryId:     uint64(item.CountryId),
			UpdatedAt:     timeUpdate,
		})
	}
	return nil
}
