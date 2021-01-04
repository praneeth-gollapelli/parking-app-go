package main

import (
	"fmt"
	"log"
	"mux"
	"net/http"
	"task/base"
	"task/controllers"
	"task/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := mux.NewRouter()
	var config model.Config
	config = base.LoadConfig(config)

	fs := http.FileServer((http.Dir("static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Load Index
	r.HandleFunc("/", Index)

	//Laod NewParkingLotController
	p := controllers.NewParkingLotController(config)
	r.HandleFunc("/show", p.FindByGroup)
	r.HandleFunc("/new", p.LoadTemplate)
	r.HandleFunc("/create/parkinglot", p.Create)

	//Load NewParkingSlotReservationConteroller
	p1 := controllers.NewParkingSlotReservationConteroller(config)
	r.HandleFunc("/park", p1.LoadTemplate)
	r.HandleFunc("/park/vehicle", p1.Create)
	r.HandleFunc("/find", p1.Find)
	http.Handle("/", r)

	log.Printf("Server started on: http://localhost%s", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	str := base.BindTemplate("templates/main.html", nil)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, str)
}
