package main

import (
	"fmt"
	"net/http"
	"time"


	storage "github.com/alokyadav/buildings/storage"
	service "github.com/alokyadav/buildings/service"
	"github.com/gorilla/mux"
	"github.com/google/jsonapi"
)

const (
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
)

func main() {
	jsonapi.Instrumentation = func(r *jsonapi.Runtime, eventType jsonapi.Event, callGUID string, dur time.Duration) {
		metricPrefix := r.Value("instrument").(string)

		if eventType == jsonapi.UnmarshalStart {
			fmt.Printf("%s: id, %s, started at %v\n", metricPrefix+".jsonapi_unmarshal_time", callGUID, time.Now())
		}

		if eventType == jsonapi.UnmarshalStop {
			fmt.Printf("%s: id, %s, stopped at, %v , and took %v to unmarshal payload\n", metricPrefix+".jsonapi_unmarshal_time", callGUID, time.Now(), dur)
		}

		if eventType == jsonapi.MarshalStart {
			fmt.Printf("%s: id, %s, started at %v\n", metricPrefix+".jsonapi_marshal_time", callGUID, time.Now())
		}

		if eventType == jsonapi.MarshalStop {
			fmt.Printf("%s: id, %s, stopped at, %v , and took %v to marshal payload\n", metricPrefix+".jsonapi_marshal_time", callGUID, time.Now(), dur)
		}
	}

	storage := storage.NewBuildingStorage()
	service := service.NewBuildingService(storage)

	r := mux.NewRouter()
	r.HandleFunc("/buildings", service.ListBuildings).Methods("GET")
	r.HandleFunc("/buildings", service.AddBuilding).Methods("POST")
	r.HandleFunc("/buildings/{id}", service.GetBuilding).Methods("GET")
	r.HandleFunc("/buildings/{id}", service.RemoveBuilding).Methods("DELETE")
	
	http.ListenAndServe(":8000", r)
	
}




