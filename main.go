package main

import (
	"assignment3/model"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	statusValue := GetStatusValue()

	var t, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, statusValue)
}

func randomIntGen(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}

func GetStatusValue() model.StatusValue {
	var statusValue model.StatusValue

	statusValue.WaterValue = randomIntGen(1, 100)
	time.Sleep(1)
	statusValue.WindValue = randomIntGen(1, 100)

	// Water Status
	if statusValue.WaterValue <= 5 {
		statusValue.WaterStatus = "Aman"
	} else if statusValue.WaterValue >= 6 && statusValue.WaterValue <= 8 {
		statusValue.WaterStatus = "Siaga"
	} else if statusValue.WaterValue > 8 {
		statusValue.WaterStatus = "Bahaya"
	} else {
		statusValue.WaterStatus = "Error"
	}

	// Wind Status
	if statusValue.WindValue <= 6 {
		statusValue.WindStatus = "Aman"
	} else if statusValue.WindValue >= 7 && statusValue.WindValue <= 15 {
		statusValue.WindStatus = "Siaga"
	} else if statusValue.WindValue > 15 {
		statusValue.WindStatus = "Bahaya"
	} else {
		statusValue.WindStatus = "Error"
	}

	return statusValue
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", getStatusHandler)

	fmt.Println("running at localhost:8888")
	http.ListenAndServe(":8888", router)
}
