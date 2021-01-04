package controllers

import (
	"encoding/json"
	"net/http"

	"parking-app-go/base"
	"parking-app-go/constants"
	"parking-app-go/model"
	"parking-app-go/services"

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
	lotID := base.StrToInt(inputMap["id"])
	inputMap["status"] = "IN"
	by, _ := json.Marshal(inputMap)
	json.Unmarshal(by, &input)
	err := pl.ParkingSlotReservationService.Create(input, lotID)
	resp := model.Response{}
	resp.Status = constants.StatusSuccess
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

//Find
func (pl *ParkingSlotReservationConteroller) Find(w http.ResponseWriter, r *http.Request) {
	inputMap := base.PareBody(r)
	parkingSlot := map[string]interface{}{}
	resp := model.Response{}
	resp.Status = ""
	if inputMap["vehicleNumber"] != nil {
		parkingSlot = pl.ParkingSlotReservationService.FindVehicle(inputMap["vehicleNumber"].(string))
		if parkingSlot["vehicleNumber"] == "" {
			resp.Status = "Not found !"
		} else {
			resp.Status = constants.StatusSuccess
		}
	}
	resp.Data = parkingSlot
	str := base.BindTemplate("templates/findAndExit.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}

//Find
func (pl *ParkingSlotReservationConteroller) Exit(w http.ResponseWriter, r *http.Request) {
	inputMap := base.PareBody(r)
	resp := model.Response{}
	resp.Status = ""
	if inputMap["id"] != nil {
		err := pl.ParkingSlotReservationService.Exit(inputMap)
		resp.Status = "Exit failed"
		if err == nil {
			resp.Status = "Exit success"
		}
	}
	inputMap["type"] = ""
	resp.Data = inputMap
	str := base.BindTemplate("templates/findAndExit.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}
