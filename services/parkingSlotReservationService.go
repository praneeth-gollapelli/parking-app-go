package services

import (
	"errors"
	"strings"
	"task/base"
	"task/base/db"
	"task/constants"
	"task/model"
	"time"

	"golang.org/x/exp/errors/fmt"
)

const (
	Query            = "JOIN parking_lots on id = parking_slots.parking_lot_id"
	parkingSlotTable = "parking_slots"
)

// parkingSlotReservationService ...
type ParkingSlotReservationService struct {
	parkingReservationTable db.Table
	parkingSlotTable        db.Table
	config                  model.Config
}

// NewparkingSlotReservationService ...
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
func (s *ParkingSlotReservationService) Create(input model.ParkingSlotReservation) error {
	if !s.isValidInput(input) {
		return errors.New(constants.InsuffInput)
	}
	if s.isSlotAllotted(input) {
		return errors.New(constants.AlreadyAllotted)
	}
	if s.isValidVehicle(input) {
		return errors.New(constants.AlreadyAllotted)
	}
	input.VehicleNumber = strings.ToUpper(strings.TrimSpace(input.VehicleNumber))
	input.ParkingSlotID = 1
	input.InTime = time.Now().String()
	input.OutTime = ""
	s.parkingReservationTable.Insert(&input)
	//update parking lot
	s.updateParkingSlotStatus(input, true)
	return nil
}

//Create ...
func (s *ParkingSlotReservationService) Exit(input model.ParkingSlotReservation) error {
	if input.ID == 0 || input.Status == "" {
		return errors.New(constants.InsuffInput)
	}
	input.OutTime = time.Now().String()
	s.parkingReservationTable.Update(&input, "id = ?", input.ID)
	s.updateParkingSlotStatus(input, false)
	return nil
}

func (s *ParkingSlotReservationService) updateParkingSlotStatus(input model.ParkingSlotReservation, status bool) {
	statusStruct := model.ParkingSlot{}
	statusStruct.Occupied = status
	s.parkingReservationTable.Update(&statusStruct, "id = ?", input.ParkingSlotID)
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
	if resp.Occupied {
		return true
	}
	return false
}

func (s *ParkingSlotReservationService) isValidVehicle(input model.ParkingSlotReservation) bool {
	parkingSlot := s.FindVehicle(input.VehicleNumber)
	if parkingSlot["status"] == "IN" {
		return true
	}
	return false
}

//FindVehicle ..
func (s *ParkingSlotReservationService) FindVehicle(vehicleNumber string) map[string]interface{} {
	var resp model.ParkingSlotReservation
	s.parkingReservationTable.Find(&resp, "vehicle_number =  ? AND status = ?", vehicleNumber, "PARKED")
	respMap := base.StructToMap(resp)
	if len(respMap) == 0 {
		return map[string]interface{}{}
	}
	results := []map[string]interface{}{}
	s.parkingSlotTable.Join(parkingSlotTable, Query, &results)
	fmt.Println("resp", respMap, "results", results)

	if len(results) == 0 {
		return respMap
	}
	respMap["parkingLot"] = results[0]
	return respMap
}
