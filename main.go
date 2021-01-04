package main

import (
	"fmt"
	"log"
	"net/http"
	"parking-app-go/base"
	"parking-app-go/controllers"
	"parking-app-go/model"

	"github.com/gorilla/mux"
)

func main() {

	//Loading config
	var config model.Config
	config = base.LoadConfig(config)

	//loading static files
	loadStatic()

	r := mux.NewRouter()

	//Loading main/home file
	r.HandleFunc("/", index)

	//Laod NewParkingLotController
	loadParkingLotHandler(r, config)

	//Load NewParkingSlotReservationConteroller
	loadParkingSlotReservationHandler(r, config)

	http.Handle("/", r)

	log.Printf("Server started on: http://localhost%s", config.Port)

	log.Fatal(http.ListenAndServe(config.Port, nil))
}

//index ... Loads the main(home template)file
func index(w http.ResponseWriter, r *http.Request) {
	str := base.BindTemplate("templates/main.html", nil)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}

//loadParkingLotHandler
/*
loadParkingLotHandler have features to create parking lot
and display list of parking lots
*/
func loadParkingLotHandler(r *mux.Router, config model.Config) {
	cntrlr := controllers.NewParkingLotController(config)
	r.HandleFunc("/show", cntrlr.FindByGroup)
	r.HandleFunc("/new", cntrlr.LoadTemplate)
	r.HandleFunc("/create/parkinglot", cntrlr.Create)
}

//loadParkingSlotReservationHandler
/*
loadParkingSlotReservationHandler have functiomality to park vehicle,
Find vehicle and exit the vehilcle
*/
func loadParkingSlotReservationHandler(r *mux.Router, config model.Config) {
	cntrlr := controllers.NewParkingSlotReservationConteroller(config)
	r.HandleFunc("/park", cntrlr.LoadTemplate)
	r.HandleFunc("/park/vehicle", cntrlr.Create)
	r.HandleFunc("/find", cntrlr.Find)
	r.HandleFunc("/exit", cntrlr.Exit)
}

//loadStatic ... loads the static files like css and images
func loadStatic() {
	fs := http.FileServer((http.Dir("static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
