package storage

import (
	"errors"
	"fmt"
	"github.com/alokyadav/buildings/model"

)


// NewUserStorage initializes the storage
func NewBuildingStorage() *BuildingStorage {
	return &BuildingStorage{make(map[string]*model.Building), 1}
}

// UserStorage stores all users
type BuildingStorage struct {
	buildings   map[string]*model.Building
	idCount int
}

// GetAll returns the building map (because we need the ID as key too)
func (s BuildingStorage) GetAll()  []*model.Building {
	buildings := []*model.Building{}
	for _,building := range s.buildings {
		buildings = append(buildings,building)
	}
	return buildings
}


// GetOne building
func (s BuildingStorage) GetOne(id string) (*model.Building, error) {
	building, ok := s.buildings[id]
	if ok {
		return building, nil
	}
	errMessage := fmt.Sprintf("Building for id %s not found", id)
	return nil, errors.New(errMessage)
}

// Insert a building
func (s *BuildingStorage) Insert(c *model.Building) string {
	id := fmt.Sprintf("%d", s.idCount)
	c.ID = id
	if c.Address != nil {
		c.Address.ID = id
	}
	for _,floor := range c.Floors {
			floor.ID = fmt.Sprintf("%d%s", s.idCount,floor.Number)
	}
	s.buildings[id] = c
	s.idCount++

	return id
}

// Delete one :(
func (s *BuildingStorage) Delete(id string) error {
	_, exists := s.buildings[id]
	if !exists {
		return fmt.Errorf("Building with id %s does not exist", id)
	}
	delete(s.buildings, id)

	return nil
}

// Update a Building
func (s *BuildingStorage) Update(c model.Building) error {
	_, exists := s.buildings[c.ID]
	if !exists {
		return fmt.Errorf("Building with id %s does not exist", c.ID)
	}
	s.buildings[c.ID] = &c

	return nil
}
