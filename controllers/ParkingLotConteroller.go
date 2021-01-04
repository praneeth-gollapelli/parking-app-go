package controllers

import (
	"encoding/json"
	"net/http"
	"task/base"
	"task/model"
	"task/services"

	"golang.org/x/exp/errors/fmt"
)

// ParkingLotController ...
type ParkingLotController struct {
	ParkingLotService services.ParkingLotService
}

// NewParkingLotController ...
func NewParkingLotController(config model.Config) *ParkingLotController {
	return &ParkingLotController{
		ParkingLotService: services.NewParkingLotService(config),
	}
}

//Create ...
func (pl *ParkingLotController) Create(w http.ResponseWriter, r *http.Request) {
	var input model.ParkingLot
	inputMap := base.PareBody(r)
	inputMap["fourwheeler_slots"] = base.StrToInt(inputMap["fourwheeler_slots"])
	inputMap["twowheeler_slots"] = base.StrToInt(inputMap["twowheeler_slots"])
	by, _ := json.Marshal(inputMap)

	json.Unmarshal(by, &input)
	err := pl.ParkingLotService.Create(input)
	resp := model.Response{}
	resp.Status = "success"
	if err != nil {
		resp.Status = err.Error()
	}
	str := base.BindTemplate("templates/NewParkingLot.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}

//LoadTemplate
func (pl *ParkingLotController) LoadTemplate(w http.ResponseWriter, r *http.Request) {
	str := base.BindTemplate("templates/NewParkingLot.html", model.Response{})
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}

//FindByGroup ...
func (pl *ParkingLotController) FindByGroup(w http.ResponseWriter, r *http.Request) {

	resp := pl.ParkingLotService.FindByGroup()

	str := base.BindTemplate("templates/parkingLotList.html", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}
