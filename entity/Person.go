package entity

import "time"

type Person struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

func (Person) TableName() string {
	return "people"
}
