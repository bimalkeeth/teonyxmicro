package protoapi

import com "teonyxmicro/mastersvc/common"

type MasterService struct{}

var ErrorResponse com.IErrorJson

func init() {
	ErrorResponse = com.NewError()
}
