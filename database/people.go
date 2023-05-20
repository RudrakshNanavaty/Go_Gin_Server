package database

// contains the person struct
import "web-server/entities"

// people slice to seed record entities.Person data.
var People = []entities.Person{
	{
		ID:   "1",
		Name: "ABC",
	},
	{
		ID:   "2",
		Name: "DEF",
	},
	{
		ID:   "3",
		Name: "GHI",
	},
}
