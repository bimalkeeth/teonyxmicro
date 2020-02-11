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
			Address: &pro.AddressProto{
				Id:            uint64(loc.Address.Id),
				Address:       loc.Address.Address,
				Street:        loc.Address.Street,
				Suburb:        loc.Address.Suburb,
				StateId:       uint64(loc.Address.StateId),
				CountryId:     uint64(loc.Address.CountryId),
				AddressTypeId: uint64(loc.Address.AddressTypeId),
				Location:      loc.Address.Location,
				State: &pro.StateProto{
					Id:        uint64(loc.Address.State.Id),
					Name:      loc.Address.State.Name,
					CountryId: uint64(loc.Address.State.CountryId),
				},
			},
		}
		data.Locations = append(data.Locations, location)
	}
	for _, op := range result.Operators {
		opr := &pro.VehicleOperatorBoundProto{
			Id:         uint64(op.Id),
			OperatorId: uint64(op.OperatorId),
			VehicleId:  uint64(op.VehicleId),
			Active:     op.Active,
			Operator: &pro.OperatorProto{
				Id:         uint64(op.Operator.Id),
				Name:       op.Operator.Name,
				SurName:    op.Operator.SurName,
				Active:     op.Operator.Active,
				DrivingLic: op.Operator.DrivingLic,
			},
			Vehicle: nil,
		}
		data.Operators = append(data.Operators, opr)
	}

	for _, trc := range result.Registrations {
		expiryDate, _ := timestamp.TimestampProto(trc.ExpiredDate)
		updateDate, _ := timestamp.TimestampProto(trc.UpdatedAt)
		track := &pro.VehicleTrackRegProto{
			Id:           uint64(trc.Id),
			RegisterDate: trc.RegisterDate.String(),
			Duration:     int32(trc.Duration),
			ExpiredDate:  expiryDate,
			Active:       trc.Active,
			VehicleId:    uint64(trc.VehicleId),
			UpdatedAt:    updateDate,
		}
		data.Registrations = append(data.Registrations, track)
	}
	res.Vehicles = append(res.Vehicles, data)
	return nil

}

func (m *MasterService) GetVehicleByRegistration(ctx context.Context, in *pro.RequestByName, res *pro.ResponseVehicle) error {

	vehManager := vh.New()
	result, err := vehManager.GetVehicleByRegistration(in.Name)
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
			Address: &pro.AddressProto{
				Id:            uint64(loc.Address.Id),
				Address:       loc.Address.Address,
				Street:        loc.Address.Street,
				Suburb:        loc.Address.Suburb,
				StateId:       uint64(loc.Address.StateId),
				CountryId:     uint64(loc.Address.CountryId),
				AddressTypeId: uint64(loc.Address.AddressTypeId),
				Location:      loc.Address.Location,
				State: &pro.StateProto{
					Id:        uint64(loc.Address.State.Id),
					Name:      loc.Address.State.Name,
					CountryId: uint64(loc.Address.State.CountryId),
				},
			},
		}
		data.Locations = append(data.Locations, location)
	}
	for _, op := range result.Operators {
		opr := &pro.VehicleOperatorBoundProto{
			Id:         uint64(op.Id),
			OperatorId: uint64(op.OperatorId),
			VehicleId:  uint64(op.VehicleId),
			Active:     op.Active,
			Operator: &pro.OperatorProto{
				Id:         uint64(op.Operator.Id),
				Name:       op.Operator.Name,
				SurName:    op.Operator.SurName,
				Active:     op.Operator.Active,
				DrivingLic: op.Operator.DrivingLic,
			},
			Vehicle: nil,
		}
		data.Operators = append(data.Operators, opr)
	}

	for _, trc := range result.Registrations {
		expiryDate, _ := timestamp.TimestampProto(trc.ExpiredDate)
		updateDate, _ := timestamp.TimestampProto(trc.UpdatedAt)
		track := &pro.VehicleTrackRegProto{
			Id:           uint64(trc.Id),
			RegisterDate: trc.RegisterDate.String(),
			Duration:     int32(trc.Duration),
			ExpiredDate:  expiryDate,
			Active:       trc.Active,
			VehicleId:    uint64(trc.VehicleId),
			UpdatedAt:    updateDate,
		}
		data.Registrations = append(data.Registrations, track)
	}
	res.Vehicles = append(res.Vehicles, data)

	return nil
}
