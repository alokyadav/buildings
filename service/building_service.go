package service

import (
	"net/http"
	"github.com/google/jsonapi"
	"github.com/alokyadav/buildings/model"
	"github.com/alokyadav/buildings/storage"
	"github.com/alokyadav/buildings/logger"
	"github.com/gorilla/mux"
	"log"
)

const (
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
)

// NewBuildingService 
func NewBuildingService(store *storage.BuildingStorage) *BuildingService {
	return &BuildingService{store}
}

type BuildingService struct{
    buildingStorage *storage.BuildingStorage
}

//AddBuilding -  Add a new building
func (srv *BuildingService) AddBuilding(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("")
	building := new(model.Building)
	w.Header().Set("Content-Type", "application/json")
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, building); err != nil {
		logger.Log.Printf("Error %+v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	
	log.Printf("Building %+v", r.Body)
    id := srv.buildingStorage.Insert(building)
	building.ID = id

	w.WriteHeader(http.StatusCreated)

	if err := jsonapiRuntime.MarshalPayload(w, building); err != nil {
		logger.Log.Printf("Error %+v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
//GetBuilding -  get a building by Id
func (srv *BuildingService) GetBuilding(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("")
	params := mux.Vars(r)
	building,err := srv.buildingStorage.GetOne(params["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logger.Log.Printf("Error %+v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)

	if err := jsonapiRuntime.MarshalPayload(w, building); err != nil {
		logger.Log.Printf("Error %+v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//ListBuildings - list all the buildings
func (srv *BuildingService) ListBuildings(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("buidlings.list")

	// but, for now
	buildings := srv.buildingStorage.GetAll();

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := jsonapiRuntime.MarshalPayload(w, buildings); err != nil {
		logger.Log.Printf("Error %+v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//RemoveBuilding - remove a building
func (srv *BuildingService) RemoveBuilding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := srv.buildingStorage.Delete(params["id"])

	if err != nil {
		logger.Log.Printf("Error %+v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	
}

