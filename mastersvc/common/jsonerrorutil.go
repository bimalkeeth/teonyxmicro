package common

type IErrorJson interface {
	GetCreateErrorJson(rsc *interface{}, err error)
}
type ErrorJson struct{}

func (e *ErrorJson) GetCreateErrorJson(rsc *interface{}, err error) {

}
