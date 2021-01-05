# parking-app-go

## purpose

A  webserver which serves a UI for managing a parking lot
   1. Allocate a parking lot to a vehicle [2 wheelers, 4 wheelers]
   2. Capacity to find allotted parking lot for a user [User forgot his parking stop]
   3. Ability to tell free space available in parking lot

## usage and implementation

User can able to view home page once open the application using URL ex: http://localhost:8080
User can able to view 3 buttons

1. Add parking lot : Once clicks on that user will have an option of create parking lot with capacity of 4 and 2 wheelers once user clicks on save button, system will create a parking lot and respected parking slots based on capacity

## ParkingLot - Model has fileds like Basic attribues like capacity and Name of the lot

```
type ParkingLot struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	FourwheelerSlots int    `json:"fourwheeler_slots"`
	TwowheelerSlots  int    `json:"twowheeler_slots"`
	Comments         string `json:"comments"`
}
```

## ParkingSlot ... Model has fileds like slot number type of slot(4 wheeler or 2 Wheeler) with respect to lotID
```
type ParkingSlot struct {
	ID           int    `json:"id"`
	Number       int    `json:"number"`
	ParkingLotID int    `json:"parkingLotID"`
	Type         string `json:"type"`
	Occupied     string `json:"occupied"`
}

```

2. Park vehicle : Once clicks on that, user can abel to view the free space available in parking lot and also user can see option of park a vehicle. Once user clicks on park 4 or 2 wheeler, user can navigate to Park Vehicle page to enter customer details. Once clicks on save the data will be saved and parking slot status will be updated

## ParkingSlotReservation .. Model has fileds like historic data of parking slot with respect to Vehicle and Customer
```
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
```

3. Find/Exit Vehicle: Once clicks on that button user can serach vehicle location and exit vehicle from parking.

## setup (to run locally)

Create a .env file in the project root with the following name=value pair. You will need mysql database username password (not shown here)

```
db_host=127.0.0.1
db_user=dbuser
db_password=password
db_port=3306
db_type=mysql
port=:8080
```

```
Build or Run Locally
```

Run below commands from project root folder: 

go get

go run main.go or go build

if you use go build then run using ./parking-app-go.exe

`````
Health check or test:
`````

Use  http://localhost:8080 URL in browser, the reponse should be a homepage with above explained buttons
alsoe you can see a log on console like `2021/01/05 02:46:50 Server started on: http://localhost:8080`

## improvements

1. We can have dashboard kind of page where user can see w.r.to parking lot
2. Report or history kind of page where user can serach and view the reports based on selected filter criteria
etc 
