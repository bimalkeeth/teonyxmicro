package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import bu "teonyxmicro/mastersvc/bucontracts"
import vh "teonyxmicro/mastersvc/manager/vehicles"

func (m *MasterService) CreateVehicle(ctx context.Context, in *pro.RequestVehicle, res *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicle(bu.VehicleBO{
		ModelId:      uint(in.Vehicle.ModelId),
		MakeId:       uint(in.Vehicle.MakeId),
		Registration: in.Vehicle.Registration,
		FleetId:      uint(in.Vehicle.FleetId),
		StatusId:     uint(in.Vehicle.StatusId),
		OfficeName:   in.Vehicle.OfficeName,
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicle(ctx context.Context, in *pro.RequestVehicle, res *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicle(bu.VehicleBO{
		Id:           uint(in.Vehicle.Id),
		ModelId:      uint(in.Vehicle.ModelId),
		MakeId:       uint(in.Vehicle.MakeId),
		Registration: in.Vehicle.Registration,
		FleetId:      uint(in.Vehicle.FleetId),
		StatusId:     uint(in.Vehicle.StatusId),
		OfficeName:   in.Vehicle.OfficeName,
	})
	res.Success = result
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) DeleteVehicle(ctx context.Context, in *pro.RequestDelete, res *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicle(uint(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) GetVehicleById(ctx context.Context, in *pro.RequestKey, res *pro.ResponseVehicle) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleById(uint(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	updateVehicleAt, _ := timestamp.TimestampProto(result.UpdatedAt)
	data := &pro.VehicleProto{
		Id:           uint64(result.Id),
		OfficeName:   result.OfficeName,
		StatusId:     uint64(result.StatusId),
		FleetId:      uint64(result.FleetId),
		Registration: result.Registration,
		MakeId:       uint64(result.MakeId),
		ModelId:      uint64(result.ModelId),
		UpdatedAt:    updateVehicleAt,
		Status: &pro.VehicleStatusProto{
			Id:         uint64(result.Status.Id),
			StatusType: result.Status.StatusType,
			StatusName: result.Status.StatusName,
		},
	}

	for _, loc := range result.Locations {

		updateLocAt, _ := timestamp.TimestampProto(loc.UpdateAt)
		location := &pro.VehicleAddressProto{
			AddressId: uint64(loc.AddressId),
			VehicleId: uint64(loc.VehicleId),
			UpdateAt:  updateLocAt,
		}

		data.Locations = append(data.Locations, location)
	}

}
