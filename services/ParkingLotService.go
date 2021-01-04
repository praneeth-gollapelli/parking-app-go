package services

import (
	"errors"
	"log"
	"task/base/db"
	"task/constants"
	"task/model"

	"golang.org/x/exp/errors/fmt"
)

const (
	joinQuery = "JOIN parking_slots on parking_lot_id = parking_lots.id and occupied=false"
	table     = "parking_lots"
)

// ParkingLotService ...
type ParkingLotService struct {
	parkingLotTable  db.Table
	parkingSlotTable db.Table
	config           model.Config
}

// NewParkingLotService ...
func NewParkingLotService(config model.Config) ParkingLotService {
	linstance := db.Instance(config, "local").TableInstance(&model.ParkingLot{})
	sinstance := db.Instance(config, "local").TableInstance(&model.ParkingSlot{})
	log.Println("1111111111")
	return ParkingLotService{
		parkingLotTable:  linstance,
		parkingSlotTable: sinstance,
		config:           config,
	}
}

//Create ...
func (s *ParkingLotService) Create(input model.ParkingLot) error {
	// t := model.ParkingLot{}
	// t.FourwheelerSlots = 440
	// s.parkingLotTable.Update(&t, "id = ?", 2)
	// fmt.Println(t)
	if !s.isValidInput(input) {
		return errors.New(constants.InsuffInput)
	}
	if s.isNameExists(input) {
		return errors.New(constants.DataExists)
	}
	s.parkingLotTable.Insert(&input)
	fmt.Println(input)
	go s.processParkingSlots(input)
	return nil
}

func (s *ParkingLotService) processParkingSlots(input model.ParkingLot) {
	if input.FourwheelerSlots > 0 {
		s.createPrkingSlots(input, "FOUR_WHEELER")
	}
	if input.TwowheelerSlots > 0 {
		s.createPrkingSlots(input, "TWO_WHEELER")
	}
}

func (s *ParkingLotService) createPrkingSlots(input model.ParkingLot, slotType string) {
	var slotNum, noOfSLots int
	switch slotType {
	case "FOUR_WHEELER":
		slotNum = 1
		noOfSLots = input.FourwheelerSlots
	case "TWO_WHEELER":
		slotNum = input.FourwheelerSlots + 1
		noOfSLots = input.TwowheelerSlots
	}
	var pslots []model.ParkingSlot
	for i := noOfSLots; i > 0; i-- {
		p := model.ParkingSlot{}
		p.Number = slotNum
		p.Type = slotType
		p.Occupied = false
		p.ParkingLotID = input.ID
		pslots = append(pslots, p)
	}
	s.parkingSlotTable.Insert(&pslots)
}

func (s *ParkingLotService) isValidInput(input model.ParkingLot) bool {
	if input.Name == "" || (input.FourwheelerSlots == 0 && input.TwowheelerSlots == 0) {
		return false
	}
	return true
}

func (s *ParkingLotService) isNameExists(input model.ParkingLot) bool {
	var resp []model.ParkingLot
	s.parkingLotTable.Find(&resp, "name = ?", input.Name)
	if len(resp) > 0 {
		return true
	}
	return false
}

//FindByID ...
func (s *ParkingLotService) FindByID(input map[string]interface{}) []model.ParkingLot {
	var resp []model.ParkingLot
	s.parkingLotTable.Find(&resp, "id = ?", input["id"])

	return resp
}

//FindByID ...
func (s *ParkingLotService) FindByGroup() []map[string]interface{} {
	parkinglots := []model.ParkingLot{}
	s.parkingLotTable.Find(&parkinglots, nil, nil)
	results := []map[string]interface{}{}
	s.parkingLotTable.Join(table, joinQuery, &results)
	resp := enrichResponse(results, parkinglots)
	return resp
}

func enrichResponse(input []map[string]interface{}, parkinglots []model.ParkingLot) []map[string]interface{} {
	response := []map[string]interface{}{}
	for _, p := range parkinglots {
		r := map[string]interface{}{}
		r["available_fourwheeler_slots"], r["available_twowheeler_slots"] = findAvailableSlotsCount(p.ID, input)
		r["id"] = p.ID
		r["name"] = p.Name
		r["fourwheeler_slots"] = p.FourwheelerSlots
		r["tworwheeler_slots"] = p.TwowheelerSlots
		response = append(response, r)
	}
	return response
}

func findAvailableSlotsCount(id int, parkinglots []map[string]interface{}) (int, int) {
	var fourwheelerSlots, twowheelerSlots int
	for _, p := range parkinglots {
		if int64(id) != p["parking_lot_id"].(int64) {
			continue
		}
		switch p["type"].(string) {
		case "FOUR_WHEELER":
			fourwheelerSlots++
		case "TWO_WHEELER":
			twowheelerSlots++
		}
	}
	return fourwheelerSlots, twowheelerSlots
}
