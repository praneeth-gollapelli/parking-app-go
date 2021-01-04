package controllers

import (
	"encoding/json"
	"net/http"
	"task/base"
	"task/model"
	"task/services"

	"golang.org/x/exp/errors/fmt"
)

// ParkingSlotReservationConteroller ...
type ParkingSlotReservationConteroller struct {
	ParkingSlotReservationService services.ParkingSlotReservationService
}

// NewParkingSlotReservationConteroller ...
func NewParkingSlotReservationConteroller(config model.Config) *ParkingSlotReservationConteroller {
	return &ParkingSlotReservationConteroller{
		ParkingSlotReservationService: services.NewParkingSlotReservationService(config),
	}
}

//Create ...
func (pl *ParkingSlotReservationConteroller) Create(w http.ResponseWriter, r *http.Request) {
	var input model.ParkingSlotReservation
	inputMap := base.PareBody(r)
	inputMap["status"] = "PARKED"
	by, _ := json.Marshal(inputMap)
	json.Unmarshal(by, &input)
	fmt.Println(string(by[:]), input)
	err := pl.ParkingSlotReservationService.Create(input)
	resp := model.Response{}
	resp.Status = "success"
	resp.Data = inputMap
	if err != nil {
		resp.Status = err.Error()
	}
	str := base.BindTemplate("templates/parkingReservation.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}

//LoadTemplate
func (pl *ParkingSlotReservationConteroller) LoadTemplate(w http.ResponseWriter, r *http.Request) {
	inputMap := base.PareBody(r)
	resp := model.Response{}
	resp.Status = ""
	resp.Data = inputMap
	str := base.BindTemplate("templates/parkingReservation.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}

//LoadTemplate
func (pl *ParkingSlotReservationConteroller) Find(w http.ResponseWriter, r *http.Request) {
	inputMap := base.PareBody(r)
	parkingSlot := map[string]interface{}{}
	resp := model.Response{}
	resp.Status = ""
	if inputMap["vehicleNumber"] != nil {
		parkingSlot = pl.ParkingSlotReservationService.FindVehicle(inputMap["vehicleNumber"].(string))
		resp.Status = "success"
	}
	resp.Data = parkingSlot
	str := base.BindTemplate("templates/findAndExit.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}
