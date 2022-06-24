package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Cabri struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var db []Cabri = make([]Cabri, 0)

func parseStringToInt(id string) int {
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

// get un cabri par son id
func getCabribyID(w http.ResponseWriter, req *http.Request) {
	id := parseStringToInt(req.URL.Query()["id"][0])
	fmt.Println(id)
	for _, cabri := range db {
		if cabri.Id == id {
			json.NewEncoder(w).Encode(cabri)
			return
		}
	}
	json.NewEncoder(w).Encode("No Cabri found")
}

// create un cabri
func createCabri(w http.ResponseWriter, req *http.Request) {
	var cabri Cabri
	payload := req.Body
	json.NewDecoder(payload).Decode(&cabri)
	cabri.Id = len(db) + 1
	db = append(db, cabri)
	json.NewEncoder(w).Encode(cabri)
}

// delete le cabri
func deleteCabri(w http.ResponseWriter, req *http.Request) {
	id := parseStringToInt(req.URL.Query()["id"][0])
	for i, cabri := range db {
		if cabri.Id == id {
			db = append(db[:i], db[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(db)
}

// update data du cabri
func updateCabri(w http.ResponseWriter, req *http.Request) {
	id := parseStringToInt(req.URL.Query()["id"][0])
	fmt.Println("id = ", id)
	var cabri Cabri
	payload := req.Body
	json.NewDecoder(payload).Decode(&cabri)
	fmt.Println("new name", cabri.Name)
	for i, v := range db {
		if v.Id == id {
			db = append(db[:i], db[i+1:]...)
			db = append(db, cabri)
			fmt.Println("here")
			break
		}
	}
	json.NewEncoder(w).Encode(db)
}

// main func
func main() {
	var port = ":8080"
	// routes
	http.HandleFunc("/create", createCabri)
	http.HandleFunc("/get", getCabribyID)
	http.HandleFunc("/update", updateCabri)
	http.HandleFunc("/delete", deleteCabri)
	log.Println("Server started on port", port)
	http.ListenAndServe(port, nil)
}
