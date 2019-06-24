package model

import (
	"time"
	//"github.com/google/jsonapi"
	//"fmt"
)

//Building - 
type Building struct {
	ID            string    `jsonapi:"primary,buildings"`
	Name          string    `jsonapi:"attr,name"`
	OwnerName     string    `jsonapi:"attr,owner_name"`
	Floors        []*Floor  `jsonapi:"relation,floors"`
	Address       *Address   `jsonapi:"relation,address"`
	CreatedAt     time.Time `jsonapi:"attr,created_at"`

}


//Floor -
type Floor struct {
	ID          string `jsonapi:"primary,floors"`
	Number		string    `jsonapi:"attr,number"`
}

//Address -  
type Address struct {
	ID      string  `jsonapi:"primary,address"`
	Line1	string  `jsonapi:"attr,line_1"`
	Line2	string  `jsonapi:"attr,line_2"`
	City    string  `jsonapi:"attr,city"`
	State	string  `jsonapi:"attr,state"`
	Country string  `jsonapi:"attr,country"`
	PinCode string  `jsonapi:"attr,pin_code"`    
}



/*
// JSONAPILinks implements the Linkable interface for a building
func (building Building) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("https://example.com/buildings/%d", building.ID),
	}
}

// JSONAPIRelationshipLinks implements the RelationshipLinkable interface for a building
func (building Building) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	if relation == "floors" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("https://example.com/buildings/%d/posts", building.ID),
		}
	}
	if relation == "address" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("https://example.com/buildings/%d/address", building.ID),
		}
	}
	return nil
}

// JSONAPIMeta implements the Metable interface for a building
func (building Building) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"detail": "extra details regarding the building",
	}
}

// JSONAPIRelationshipMeta implements the RelationshipMetable interface for a building
func (building Building) JSONAPIRelationshipMeta(relation string) *jsonapi.Meta {
	if relation == "floors" {
		return &jsonapi.Meta{
			"detail": "floors meta information",
		}
	}
	if relation == "address" {
		return &jsonapi.Meta{
			"detail": "address meta information",
		}
	}
	return nil
}
*/