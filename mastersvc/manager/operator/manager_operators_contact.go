package operator

import (
	bs "teonyxmicro/mastersvc/business"
)

//----------------------------------------------------
//Create operation contact
//----------------------------------------------------
func (o *OprManager) CreateOperatorContact(contactId uint, operatorId uint, primary bool) (uint, error) {
	op := OpFac.New(bs.COperationContact).(bs.OperatorContact)
	OpFac.Conn.Begin()
	res, err := op.CreateOperatorContact(contactId, operatorId, primary)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//Update operation contact
//----------------------------------------------------
func (o *OprManager) UpdateOperatorContact(id uint, contactId uint, operatorId uint, primary bool) (bool, error) {
	op := OpFac.New(bs.COperationContact).(bs.OperatorContact)
	OpFac.Conn.Begin()
	res, err := op.UpdateOperatorContact(id, contactId, operatorId, primary)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}
