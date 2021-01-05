package model

import "time"

//ParkingLot - Model has fileds like Basic attribues like capacity and Name of the lot
type ParkingLot struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	FourwheelerSlots int    `json:"fourwheeler_slots"`
	TwowheelerSlots  int    `json:"twowheeler_slots"`
	Comments         string `json:"comments"`
}

//ParkingSlot ... Model has fileds like slot number type of slot(4 wheeler or 2 Wheeler) with respect to lotID
type ParkingSlot struct {
	ID           int    `json:"id"`
	Number       int    `json:"number"`
	ParkingLotID int    `json:"parkingLotID"`
	Type         string `json:"type"`
	Occupied     string `json:"occupied"`
}

//ParkingSlotReservation .. Model has fileds like historic data of parking slot with respect to Vehicle and Customer
type ParkingSlotReservation struct {
	ID            int    `json:"id"`
	VehicleNumber string `json:"vehicleNumber"`
	CutomerName   string `json:"name"`
	ContacNumber  string `json:"contacNumber"`
	Type          string `json:"type"`
	ParkingSlotID int    `json:"parkingSlotID"`
	ParkingSlotNo int    `json:"parkingSlotNo"`
	Status        string `json:"status"`
	InTime        string `json:"inTime"`
	OutTime       string `json:"outTime"`
}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	EmailID string    `json:"email"`
	InTime  time.Time `json:"inTime"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
