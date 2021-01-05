package services

import (
	"errors"
	"parking-app-go/base"
	"parking-app-go/base/db"
	"parking-app-go/constants"
	"parking-app-go/model"
	"strings"
	"time"
)

const (
	Query            = "JOIN parking_lots on parking_lots.id = parking_slots.parking_lot_id  AND parking_slots.id = ?"
	parkingSlotTable = "parking_slots"
)

// ParkingSlotReservationService ...
type ParkingSlotReservationService struct {
	parkingReservationTable db.Table
	parkingSlotTable        db.Table
	config                  model.Config
}

// NewParkingSlotReservationService ...
func NewParkingSlotReservationService(config model.Config) ParkingSlotReservationService {
	t1 := db.Instance(config, "local").TableInstance(&model.ParkingSlotReservation{})
	t2 := db.Instance(config, "local").TableInstance(&model.ParkingSlot{})
	return ParkingSlotReservationService{
		parkingReservationTable: t1,
		parkingSlotTable:        t2,
		config:                  config,
	}
}

//Create ...
func (s *ParkingSlotReservationService) Create(input model.ParkingSlotReservation, lotID int) error {
	if !s.isValidInput(input) {
		return errors.New(constants.InsuffInput)
	}
	if s.isValidVehicle(input) {
		return errors.New(constants.VehicleAlreadyAllotted)
	}
	input.VehicleNumber = strings.ToUpper(strings.Replace(input.VehicleNumber, " ", "", -1))
	pslot := s.getActiveParkingSlot(lotID, input.Type)
	input.ParkingSlotID = pslot.ID
	input.ParkingSlotNo = pslot.Number
	if input.ParkingSlotID == 0 {
		return errors.New(constants.AlreadyAllotted)
	}
	input.InTime = time.Now().String()
	input.OutTime = ""
	s.parkingReservationTable.Insert(&input)
	//update parking lot
	s.updateParkingSlotStatus(input.ParkingSlotID, "YES")
	return nil
}

//Exit ...
func (s *ParkingSlotReservationService) Exit(params map[string]interface{}) error {
	var input model.ParkingSlotReservation
	input.OutTime = time.Now().String()
	input.Status = "OUT"
	s.parkingReservationTable.Update(input, "id = ?", base.StrToInt(params["id"]))
	s.updateParkingSlotStatus(base.StrToInt(params["parkingSlotID"]), "NO")
	return nil
}

//FindVehicle ..
func (s *ParkingSlotReservationService) FindVehicle(vehicleNumber string) map[string]interface{} {
	var resp model.ParkingSlotReservation
	s.parkingReservationTable.Find(&resp, "vehicle_number =  ? AND status = ?", strings.ToUpper(strings.Replace(vehicleNumber, " ", "", -1)), "IN")
	respMap := base.StructToMap(resp)
	if respMap["vehicleNumber"] == "" {
		return respMap
	}
	results := map[string]interface{}{}
	s.parkingSlotTable.Join(parkingSlotTable, Query, &results, resp.ParkingSlotID)
	respMap["parkingLot"] = results
	return respMap
}

/*
Below are the helper funcs
*/

func (s *ParkingSlotReservationService) updateParkingSlotStatus(slotID int, status string) {
	statusStruct := model.ParkingSlot{}
	statusStruct.Occupied = status
	s.parkingSlotTable.Update(statusStruct, "id = ?", slotID)
}

//
func (s *ParkingSlotReservationService) isValidInput(input model.ParkingSlotReservation) bool {
	if input.VehicleNumber == "" || input.CutomerName == "" || input.ContacNumber == "" || input.Status == "" {
		return false
	}
	return true
}

func (s *ParkingSlotReservationService) isSlotAllotted(input model.ParkingSlotReservation) bool {
	var resp model.ParkingSlot
	s.parkingSlotTable.Find(&resp, "id = ?", input.ParkingSlotID)
	if resp.Occupied == "YES" {
		return true
	}
	return false
}

func (s *ParkingSlotReservationService) getActiveParkingSlot(input int, tType string) model.ParkingSlot {
	var resp model.ParkingSlot
	s.parkingSlotTable.Find(&resp, "parking_lot_id = ? and occupied = ? and type =?", input, "NO", tType)
	return resp
}

func (s *ParkingSlotReservationService) isValidVehicle(input model.ParkingSlotReservation) bool {
	parkingSlot := s.FindVehicle(input.VehicleNumber)
	if parkingSlot["status"] == "IN" {
		return true
	}
	return false
}
