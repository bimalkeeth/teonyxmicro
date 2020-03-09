package protoapi

import (
	"context"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import bu "teonyxmicro/mastersvc/bucontracts"
import flt "teonyxmicro/mastersvc/manager/fleets"
import timestamp "github.com/golang/protobuf/ptypes"

func (m *MasterService) CreateFleet(ctx context.Context, req *pro.RequestFleet, out *pro.ResponseCreateSuccess) error {
	fltManager := flt.New()

	dateRegistered, err := timestamp.Timestamp(req.Fleet.DateRegistered)
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}
	result, errs := fltManager.CreateFleet(bu.FleetBO{
		Name:                 req.Fleet.Name,
		CountryId:            uint(req.Fleet.CountryId),
		FleetCode:            req.Fleet.FleetCode,
		SurName:              req.Fleet.SurName,
		OtherName:            req.Fleet.OtherName,
		RegistrationDuration: float64(req.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(errs)
	out.Id = uint64(result.Id)
	return nil
}

func (m *MasterService) UpdateFleet(ctx context.Context, req *pro.RequestFleet, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	dateRegistered, err := timestamp.Timestamp(req.Fleet.DateRegistered)
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}
	result, errs := fltManager.UpdateFleet(bu.FleetBO{
		Id:                   uint(req.Fleet.Id),
		Name:                 req.Fleet.Name,
		CountryId:            uint(req.Fleet.CountryId),
		FleetCode:            req.Fleet.FleetCode,
		SurName:              req.Fleet.SurName,
		OtherName:            req.Fleet.OtherName,
		RegistrationDuration: float64(req.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(errs)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleet(ctx context.Context, req *pro.RequestDelete, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.DeleteFleet(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetFleetById(ctx context.Context, req *pro.RequestKey, out *pro.ResponseFleet) error {
	fltManager := flt.New()
	result, err := fltManager.GetFleetById(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	updatedat, _ := timestamp.TimestampProto(result.UpdatedAt)
	dateRegistered, _ := timestamp.TimestampProto(result.DateRegistered)

	out.Fleet = append(out.Fleet, &pro.FleetProto{
		Id:                   uint64(result.Id),
		UpdatedAt:            updatedat,
		FleetCode:            result.FleetCode,
		Name:                 result.Name,
		SurName:              result.SurName,
		OtherName:            result.OtherName,
		DateRegistered:       dateRegistered,
		RegistrationDuration: float32(result.RegistrationDuration),
		FleetContactId:       uint64(result.FleetContactId),
		FleetLocationId:      uint64(result.FleetLocationId),
		CountryId:            uint64(result.CountryId),
	})
	return nil
}

func (m *MasterService) CreateFleetContact(ctx context.Context, req *pro.RequestFleetContact, out *pro.ResponseCreateSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.CreateFleetContact(uint(req.FleetId), uint(req.ContactId), req.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateFleetContact(ctx context.Context, req *pro.RequestFleetContact, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.UpdateFleetContact(uint(req.Id), uint(req.FleetId), uint(req.ContactId), req.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleetContact(ctx context.Context, req *pro.RequestDelete, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.DeleteFleetContact(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetContactByFleetId(ctx context.Context, req *pro.RequestKey, out *pro.ResponseFleetContact) error {
	fltManager := flt.New()
	result, err := fltManager.GetContactByFleetId(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, con := range result {

		updatedat, _ := timestamp.TimestampProto(con.Fleet.UpdatedAt)
		dateRegistered, _ := timestamp.TimestampProto(con.Fleet.DateRegistered)
		contact := &pro.FleetContactProto{
			Id:        uint64(con.Id),
			FleetId:   uint64(con.FleetId),
			ContactId: uint64(con.ContactId),
			Primary:   con.Primary,
		}
		contact.Fleet = &pro.FleetProto{
			Id:             uint64(con.Fleet.Id),
			FleetCode:      con.Fleet.FleetCode,
			Name:           con.Fleet.Name,
			SurName:        con.Fleet.SurName,
			OtherName:      con.Fleet.OtherName,
			DateRegistered: dateRegistered,
			UpdatedAt:      updatedat,
		}
		contact.Contact = &pro.ContactProto{
			Id:            uint64(con.Contact.Id),
			Contact:       con.Contact.Contact,
			ContactTypeId: uint64(con.Contact.ContactTypeId),
		}
		out.FleetContact = append(out.FleetContact, contact)
	}
	return nil
}

func (m *MasterService) CreateFleetLocation(ctx context.Context, req *pro.RequestFleetLocation, out *pro.ResponseCreateSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.CreateFleetLocation(uint(req.FleetId), uint(req.AddressId), req.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateFleetLocation(ctx context.Context, in *pro.RequestFleetLocation) (*pro.ResponseSuccess, error) {
	fltManager := flt.New()
	response := &pro.ResponseSuccess{}
	result, err := fltManager.UpdateFleetLocation(uint(in.Id), uint(in.FleetId), uint(in.AddressId), in.Primary)
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteFleetLocation(ctx context.Context, req *pro.RequestDelete, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.DeleteFleetLocation(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetLocationByFleetId(ctx context.Context, req *pro.RequestKey, out *pro.ResponseFleetLocation) error {
	fltManager := flt.New()
	result, err := fltManager.GetLocationByFleetId(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, loc := range result {
		item := &pro.FleetLocationProto{
			Id:        uint64(loc.Id),
			FleetId:   uint64(loc.FleetId),
			AddressId: uint64(loc.AddressId),
			Primary:   loc.Primary,
		}
		updatedat, _ := timestamp.TimestampProto(loc.Address.UpdatedAt)
		item.Address = &pro.AddressProto{
			Id:            uint64(loc.Address.Id),
			Address:       loc.Address.Address,
			Street:        loc.Address.Street,
			Suburb:        loc.Address.Suburb,
			StateId:       uint64(loc.Address.StateId),
			CountryId:     uint64(loc.Address.CountryId),
			AddressTypeId: uint64(loc.Address.AddressTypeId),
			Location:      loc.Address.Location,
			UpdatedAt:     updatedat,
		}
		fltUpdatedAt, _ := timestamp.TimestampProto(loc.Fleet.UpdatedAt)
		dateRegisterAt, _ := timestamp.TimestampProto(loc.Fleet.DateRegistered)
		item.Fleet = &pro.FleetProto{
			Id:             uint64(loc.Fleet.Id),
			CountryId:      uint64(loc.Fleet.CountryId),
			DateRegistered: dateRegisterAt,
			OtherName:      loc.Fleet.OtherName,
			SurName:        loc.Fleet.SurName,
			Name:           loc.Fleet.Name,
			UpdatedAt:      fltUpdatedAt,
		}
		out.FleetLocation = append(out.FleetLocation, item)
	}
	return nil
}
