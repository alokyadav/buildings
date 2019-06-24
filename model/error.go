package model

import (
	"fmt"
)

type ErrorObject struct { 
	 Title  string  `jsonapi:"attr,title"`
	 Detail  string  `jsonapi:"attr,detail"`
 }

 func (e *ErrorObject) Error() string {
	return fmt.Sprintf("Error: %s %s\n", e.Title, e.Detail)
}