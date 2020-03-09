package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
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
func (m *MasterService) GetVehiclesByFleetId(ctx context.Context, in *pro.RequestKey, res *pro.ResponseVehicle) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehiclesByFleetId(uint(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, item := range result {
		updateVehicleAt, _ := timestamp.TimestampProto(item.UpdatedAt)
		data := &pro.VehicleProto{
			Id:           uint64(item.Id),
			OfficeName:   item.OfficeName,
			StatusId:     uint64(item.StatusId),
			FleetId:      uint64(item.FleetId),
			Registration: item.Registration,
			MakeId:       uint64(item.MakeId),
			ModelId:      uint64(item.ModelId),
			UpdatedAt:    updateVehicleAt,
			Status: &pro.VehicleStatusProto{
				Id:         uint64(item.Status.Id),
				StatusType: item.Status.StatusType,
				StatusName: item.Status.StatusName,
			},
		}

		for _, loc := range item.Locations {

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

		for _, op := range item.Operators {
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

		for _, trc := range item.Registrations {
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
	}
	return nil
}

func (m *MasterService) CreateVehicleHistory(ctx context.Context, in *pro.RequestVehicleHistory, res *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	history := in.VehicleHistory
	changeDate, _ := timestamp.Timestamp(history.ChangeDate)
	result, err := vehManager.CreateVehicleHistory(bu.VehicleHistoryBO{
		VehicleId:    uint(history.VehicleId),
		ChangeDate:   changeDate,
		Description:  history.Description,
		FromStatusId: uint(history.FromStatusId),
		OfficerName:  history.OfficerName,
		ToStatusId:   uint(history.ToStatusId)})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleHistory(ctx context.Context, in *pro.RequestVehicleHistory, res *pro.ResponseSuccess) error {
	vehManager := vh.New()
	history := in.VehicleHistory
	changeDate, _ := timestamp.Timestamp(history.ChangeDate)
	result, err := vehManager.UpdateVehicleHistory(bu.VehicleHistoryBO{
		Id:           uint(history.Id),
		ToStatusId:   uint(history.ToStatusId),
		OfficerName:  history.OfficerName,
		FromStatusId: uint(history.FromStatusId),
		Description:  history.Description,
		ChangeDate:   changeDate,
		VehicleId:    uint(history.VehicleId),
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleHistory(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleHistory(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetVehicleHistoryByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleHistory) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleHistoryByVehicleId(uint(in.Id))

	for _, item := range result {
		changeDate, _ := timestamp.TimestampProto(item.ChangeDate)
		out.VehicleHistory = append(out.VehicleHistory, &pro.VehicleHistoryProto{
			ChangeDate:   changeDate,
			Description:  item.Description,
			FromStatusId: uint64(item.FromStatusId),
			OfficerName:  item.OfficerName,
			VehicleId:    uint64(item.VehicleId),
			ToStatusId:   uint64(item.ToStatusId),
			Id:           uint64(item.Id),
			FromStatus: &pro.VehicleStatusProto{
				Id:         uint64(item.FromStatus.Id),
				StatusType: item.FromStatus.StatusType,
				StatusName: item.FromStatus.StatusName,
			},
			ToStatus: &pro.VehicleStatusProto{
				Id:         uint64(item.ToStatus.Id),
				StatusType: item.ToStatus.StatusType,
				StatusName: item.ToStatus.StatusName,
			},
		})
	}
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) CreateVehicleLocation(ctx context.Context, in *pro.RequestVehicleAddress, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleLocation(bu.VehicleAddressBO{
		VehicleId: uint(in.VehicleAddress.VehicleId),
		AddressId: uint(in.VehicleAddress.AddressId),
		Primary:   in.VehicleAddress.Primary,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleLocation(ctx context.Context, in *pro.RequestVehicleAddress, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleLocation(bu.VehicleAddressBO{
		VehicleId: uint(in.VehicleAddress.VehicleId),
		AddressId: uint(in.VehicleAddress.AddressId),
		Primary:   in.VehicleAddress.Primary,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleLocation(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetVehicleLocationByVehicle(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleAddress) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleLocationByVehicle(uint(in.Id))

	for _, item := range result {

		vehicleDate, _ := timestamp.TimestampProto(item.UpdateAt)
		addUpdateAt, _ := timestamp.TimestampProto(item.Address.UpdatedAt)
		vehicleAddress := &pro.VehicleAddressProto{
			VehicleId: uint64(item.VehicleId),
			AddressId: uint64(item.AddressId),
			Primary:   item.Primary,
			UpdateAt:  vehicleDate,
			Address: &pro.AddressProto{
				Id:            uint64(item.Address.Id),
				Address:       item.Address.Address,
				Street:        item.Address.Street,
				Suburb:        item.Address.Suburb,
				StateId:       uint64(item.Address.StateId),
				CountryId:     uint64(item.Address.CountryId),
				AddressTypeId: uint64(item.Address.AddressTypeId),
				Location:      item.Address.Location,
				UpdatedAt:     addUpdateAt,
			},
		}
		out.Address = append(out.Address, vehicleAddress)
	}
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) CreateVehicleMake(ctx context.Context, in *pro.RequestVehicleMake, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleMake(bu.VehicleMakeBO{
		CountryId: uint(in.VehicleMake.CountryId),
		Make:      in.VehicleMake.Make,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleMake(ctx context.Context, in *pro.RequestVehicleMake, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleMake(bu.VehicleMakeBO{
		Id:        uint(in.VehicleMake.Id),
		CountryId: uint(in.VehicleMake.CountryId),
		Make:      in.VehicleMake.Make,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleMake(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleMake(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetAllVehicleMake(ctx context.Context, in *empty.Empty, out *pro.ResponseVehicleMake) error {
	vehManager := vh.New()
	result, err := vehManager.GetAllVehicleMake()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.VehicleMake = append(out.VehicleMake, &pro.VehicleMakeProto{
			Id:        uint64(item.Id),
			Make:      item.Make,
			CountryId: uint64(item.CountryId),
			UpdateAt:  item.UpdateAt.String(),
			Country: &pro.CountryProto{
				Id:          uint64(item.Country.Id),
				CountryName: item.Country.CountryName,
				RegionId:    uint64(item.Country.RegionId),
			},
		})
	}
	return nil
}

func (m *MasterService) GetVehicleMakeById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleMake) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleMakeById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.VehicleMake = append(out.VehicleMake, &pro.VehicleMakeProto{
		Id:        uint64(result.Id),
		Make:      result.Make,
		CountryId: uint64(result.CountryId),
		UpdateAt:  result.UpdateAt.String(),
		Country: &pro.CountryProto{
			Id:          uint64(result.Country.Id),
			CountryName: result.Country.CountryName,
			RegionId:    uint64(result.Country.RegionId),
		}})

	return nil
}
