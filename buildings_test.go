package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/mux"

	"github.com/alokyadav/buildings/storage"
	"github.com/alokyadav/buildings/service"
	"github.com/alokyadav/buildings/model"

	"github.com/google/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Buildings", func() {
	var rec *httptest.ResponseRecorder

	var addBuilding = func() {
		storage := storage.NewBuildingStorage()
		service := service.NewBuildingService(storage)

		r := mux.NewRouter()
		r.HandleFunc("/buildings", service.AddBuilding).Methods("POST")
		rec = httptest.NewRecorder()
        building := &model.Building {Name : "wework", OwnerName : "Wework" ,Floors : []*model.Floor {&model.Floor{Number : "0"}}, Address : &model.Address {Line1: "Street 1", Line2: "Street 2"}}
		requestBody := bytes.NewBuffer(nil)
        jsonapi.MarshalOnePayloadEmbedded(requestBody,building)
		req, err := http.NewRequest("POST", "/buildings", requestBody)
		Expect(err).ToNot(HaveOccurred())
		r.ServeHTTP(rec, req)
		Expect(rec.Code).To(Equal(http.StatusCreated))
	}

	It("Adds a new Building ", func() {
		addBuilding()
	})

	var getBuilding = func() {
		building := &model.Building {Name : "wework", OwnerName : "Wework"}
		storage := storage.NewBuildingStorage()
		service := service.NewBuildingService(storage)
		storage.Insert(building)
		r := mux.NewRouter()
		r.HandleFunc("/buildings/{id}", service.GetBuilding).Methods("GET")
		rec = httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/buildings/1", nil)
		Expect(err).ToNot(HaveOccurred())
		r.ServeHTTP(rec, req)
		Expect(rec.Code).To(Equal(http.StatusOK))
		Expect(rec.Body.String()).To(MatchJSON(`
		{
			"data": {
				"type": "buildings",
				"id": "1",
				"attributes": {
				"name": "wework",
				"owner_name": "Wework"
				},
				"relationships": {
				"address": {
					"data": null
				},
				"floors": {
					"data": []
				}
				}
			}
		}`))

	}

	It("Get building by Id", func() {
		getBuilding()
	})


	var listBuildings = func() {
		building1 := &model.Building {Name : "Test 1", OwnerName : "Wework"}
		storage := storage.NewBuildingStorage()
		service := service.NewBuildingService(storage)
		storage.Insert(building1)
		r := mux.NewRouter()
		r.HandleFunc("/buildings", service.ListBuildings).Methods("GET")
		rec = httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/buildings", nil)
		Expect(err).ToNot(HaveOccurred())
		r.ServeHTTP(rec, req)
		Expect(rec.Code).To(Equal(http.StatusOK))
		Expect(rec.Body.String()).To(MatchJSON(`
		{
			"data": [
				{
					"type": "buildings",
					"id": "1",
					"attributes": {
						"name": "Test 1",
						"owner_name": "Wework"
					},
					"relationships": {
						"address": {
							"data": null
						},
						"floors": {
							"data": []
						}
					}
				}
			]
		}`))

	}

	It("Get building list", func() {
		listBuildings()
	})

	var removeBuilding = func() {
		building1 := &model.Building {Name : "Test 1", OwnerName : "Wework"}
		building2 := &model.Building {Name : "Test 2", OwnerName : "Wework"}
		storage := storage.NewBuildingStorage()
		service := service.NewBuildingService(storage)
		storage.Insert(building1)
		storage.Insert(building2)
		r := mux.NewRouter()
		r.HandleFunc("/buildings/{id}", service.RemoveBuilding).Methods("DELETE")
		rec = httptest.NewRecorder()

		req, err := http.NewRequest("DELETE", "/buildings/1", nil)
		Expect(err).ToNot(HaveOccurred())
		r.ServeHTTP(rec, req)
		Expect(err).ToNot(HaveOccurred())
		r.ServeHTTP(rec, req)
		Expect(rec.Code).To(Equal(http.StatusNoContent))	

	}

	It("Rermove building from DB", func() {
		removeBuilding()
	})


})
