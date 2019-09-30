package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	"teonyxmicro/mastersvc/entities"
)

type IStates interface {
	CreateState(db *gorm.DB, bo bu.StateBO) (bool, error)
	UpdateState(db *gorm.DB, bo bu.StateBO) (bool, error)
	DeleteState(db *gorm.DB, id uint) (bool, error)
	GetStateById(db *gorm.DB, id uint) (bu.StateBO, error)
	GetStateByCountryId(db *gorm.DB, id uint) ([]bu.StateBO, error)
	GetStateByName(db *gorm.DB, name string) (bu.StateBO, error)
	GetAll(db *gorm.DB) ([]bu.StateBO, error)
}

type State struct{}

func NewState() IStates {
	return &State{}
}

func (s *State) CreateState(db *gorm.DB, bo bu.StateBO) (bool, error) {

	db.Create(&entities.TableState{Name: bo.Name, CountryId: bo.CountryId})
	return true, nil
}
func (s *State) UpdateState(db *gorm.DB, bo bu.StateBO) (bool, error) {

	state := entities.TableState{}
	db.First(&state, bo.Id)
	if state.ID == 0 {
		return false, errors.New("state not found")
	}
	state.CountryId = bo.CountryId
	state.Name = bo.Name

	db.Save(&state)
	return true, nil
}
func (s *State) DeleteState(db *gorm.DB, id uint) (bool, error) {

	found := entities.TableState{}
	db.First(&found, id)
	if found.ID == 0 {
		return false, errors.New("contact type not found")
	}
	db.Delete(&found)
	return true, nil
}
func (s *State) GetStateById(db *gorm.DB, id uint) (bu.StateBO, error) {
	result := entities.TableState{}
	db.First(&result, id)
	resultBO := bu.StateBO{}
	if result.ID == 0 {
		return resultBO, errors.New("state not found")
	}
	return resultBO, nil
}
func (s *State) GetStateByCountryId(db *gorm.DB, id uint) ([]bu.StateBO, error) {
	var resultsEntities []entities.TableState
	var results []bu.StateBO
	var country entities.TableCountry

	db.Where(&entities.TableState{CountryId: id}).Find(&resultsEntities).Related(&country)

	for _, item := range resultsEntities {
		results = append(results, bu.StateBO{CountryId: item.CountryId, Name: item.Name, Id: item.ID})
	}
	return results, nil
}
func (s *State) GetStateByName(db *gorm.DB, name string) (bu.StateBO, error) {
	state := entities.TableState{}
	db.Where(&entities.TableState{Name: name}).First(&state)
	if state.ID == 0 {
		return bu.StateBO{}, errors.New("record not found")
	}
	return bu.StateBO{Name: state.Name, CountryId: state.CountryId, Id: state.ID}, nil
}
func (s *State) GetAll(db *gorm.DB) ([]bu.StateBO, error) {

	var states []entities.TableState
	var stateResults []bu.StateBO
	db.Find(&states)
	for _, item := range states {
		stateResults = append(stateResults, bu.StateBO{Name: item.Name, CountryId: item.CountryId, Id: item.ID})
	}
	return stateResults, nil
}
