package models

import (
	"encoding/json"
	"net/http"
)

type Lot struct {
	ID             string
	Name           string
	SmallCarSpots  int
	MediumCarSpots int
	Rates          string
	Full           bool
	Address        *Address
}

type Address struct {
	Street   string
	City     string
	Province string
	Country  string
}

var lots []Lot

func GetLots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lots)
}

func CreateLot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteLot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func FindLot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ParkCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UnparkCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateLot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CreateDummyLots() {
	lots = append(lots, Lot{ID: "1", Name: "Express", SmallCarSpots: 100, MediumCarSpots: 0, Rates: "$5 per hour", Full: false, Address: &Address{Street: "10 Niagara St", City: "Toronto", Province: "ON", Country: "Canada"}})

	lots = append(lots, Lot{ID: "2", Name: "Daily", SmallCarSpots: 150, MediumCarSpots: 0, Rates: "$50 per day", Full: false, Address: &Address{Street: " 20 Niagara St", City: "Toronto", Province: "ON", Country: "Canada"}})

	lots = append(lots, Lot{ID: "3", Name: "Express", SmallCarSpots: 200, MediumCarSpots: 100, Rates: "No charge for first 15 minutes. 5hrs or less: Small car - $10 per hour, Medium car - $20 per hour, more than 5hrs - daily maximum, saturday - half price", Full: false, Address: &Address{Street: "30 Niagara St", City: "Toronto", Province: "ON", Country: "Canada"}})

}
