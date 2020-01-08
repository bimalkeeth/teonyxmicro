package protoapi

import (
	"context"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import bu "teonyxmicro/mastersvc/bucontracts"
import flt "teonyxmicro/mastersvc/manager/fleets"
import timestamp "github.com/golang/protobuf/ptypes"

func (m *MasterService) CreateFleet(ctx context.Context, in *pro.RequestFleet) (*pro.ResponseCreateSuccess, error) {
	fltManager := flt.New()
	response := &pro.ResponseCreateSuccess{}
	dateRegistered, err := timestamp.Timestamp(in.Fleet.DateRegistered)
	if err != nil {
		response.Errors = ErrorResponse.GetCreateErrorJson(err)
		return response, nil
	}
	result, err := fltManager.CreateFleet(bu.FleetBO{
		Name:                 in.Fleet.Name,
		CountryId:            uint(in.Fleet.CountryId),
		FleetCode:            in.Fleet.FleetCode,
		SurName:              in.Fleet.SurName,
		OtherName:            in.Fleet.OtherName,
		RegistrationDuration: float64(in.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result.Id)
	return response, nil
}

func (m *MasterService) UpdateFleet(ctx context.Context, in *pro.RequestFleet) (*pro.ResponseSuccess, error) {
	fltManager := flt.New()
	response := &pro.ResponseSuccess{}
	dateRegistered, err := timestamp.Timestamp(in.Fleet.DateRegistered)
	if err != nil {
		response.Errors = ErrorResponse.GetCreateErrorJson(err)
		return response, nil
	}
	result, errs := fltManager.UpdateFleet(bu.FleetBO{
		Id:                   uint(in.Fleet.Id),
		Name:                 in.Fleet.Name,
		CountryId:            uint(in.Fleet.CountryId),
		FleetCode:            in.Fleet.FleetCode,
		SurName:              in.Fleet.SurName,
		OtherName:            in.Fleet.OtherName,
		RegistrationDuration: float64(in.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	response.Errors = ErrorResponse.GetCreateErrorJson(errs)
	response.Success = result
	return response, nil
}

func (m *MasterService) DeleteFleet(ctx context.Context, in *pro.RequestDelete) (*pro.ResponseSuccess, error) {
	fltManager := flt.New()
	response := &pro.ResponseSuccess{}
	result, err := fltManager.DeleteFleet(uint(in.Id))
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Success = result
	return response, nil
}
