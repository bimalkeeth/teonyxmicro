package protoapi

import (
	"context"
)
import pro "teonyxmicro/mastersvc/proto/builder"
import bu "teonyxmicro/mastersvc/bucontracts"
import opr "teonyxmicro/mastersvc/manager/operator"
import timestamp "github.com/golang/protobuf/ptypes"

func (m *MasterService) CreateOperator(ctx context.Context, req *pro.RequestOperator, res *pro.ResponseCreateSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.CreateOperator(bu.OperatorBO{
		Name:       req.Operator.Name,
		SurName:    req.Operator.SurName,
		DrivingLic: req.Operator.DrivingLic,
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateOperator(ctx context.Context, req *pro.RequestOperator, res *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.UpdateOperator(bu.OperatorBO{
		Id:         uint(req.Operator.Id),
		Name:       req.Operator.Name,
		SurName:    req.Operator.SurName,
		Active:     req.Operator.Active,
		DrivingLic: req.Operator.DrivingLic,
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteOperator(ctx context.Context, req *pro.RequestDelete, res *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.DeleteOperator(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) GetOperatorById(ctx context.Context, req *pro.RequestKey, res *pro.ResponseOperator) error {
	oprManager := opr.New()
	result, err := oprManager.GetOperatorById(uint(req.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	data := &pro.OperatorProto{}
	data.Id = uint64(result.Id)
	data.Active = result.Active
	data.DrivingLic = result.DrivingLic
	data.SurName = result.SurName
	data.Name = result.Name

	for _, con := range result.Contacts {
		updateCon, _ := timestamp.TimestampProto(con.Contact.UpdatedAt)
		contact := &pro.OperatorContactsProto{
			Id:        uint64(con.Id),
			Primary:   con.Primary,
			ContactId: uint64(con.ContactId),
			Contact: &pro.ContactProto{
				Id:            uint64(con.Contact.Id),
				Contact:       con.Contact.Contact,
				ContactTypeId: uint64(con.Contact.ContactTypeId),
				UpdatedAt:     updateCon,
			},
		}
		data.Contacts = append(data.Contacts, contact)
	}
	for _, add := range result.Locations {
		updateLoc, _ := timestamp.TimestampProto(add.UpdateAt)
		address := &pro.OperatorLocationProto{
			Id:         uint64(add.Id),
			AddressId:  uint64(add.AddressId),
			OperatorId: uint64(add.OperatorId),
			Primary:    add.Primary,
			UpdateAt:   updateLoc,
		}
		updateAdd, _ := timestamp.TimestampProto(add.Address.UpdatedAt)
		addr := &pro.AddressProto{}
		addr.Address = add.Address.Address
		addr.Id = uint64(add.Address.Id)
		addr.CountryId = uint64(add.Address.CountryId)
		addr.AddressTypeId = uint64(add.Address.AddressTypeId)
		addr.StateId = uint64(add.Address.StateId)
		addr.UpdatedAt = updateAdd
		addr.Location = add.Address.Location
		addr.Suburb = add.Address.Suburb
		addr.Street = add.Address.Street
		address.Address = addr
		data.Locations = append(data.Locations, address)
	}

	for _, vh := range result.Vehicles {
		updateVh, _ := timestamp.TimestampProto(vh.UpdatedAt)
		vehicle := &pro.VehicleProto{
			Id:           uint64(vh.Id),
			UpdatedAt:    updateVh,
			MakeId:       uint64(vh.MakeId),
			ModelId:      uint64(vh.ModelId),
			StatusId:     uint64(vh.StatusId),
			Registration: vh.Registration,
			FleetId:      uint64(vh.FleetId),
		}
		data.Vehicles = append(data.Vehicles, vehicle)
	}
	res.Operator = append(res.Operator, data)
	return nil

}
