package main

import (
	"encoding/json"
	"fmt"
	"laundry-backend/internal/entities"
)

func main() {
	// Create a sample DataTables request
	request := entities.DataTablesRequest{
		Draw:   1,
		Start:  0,
		Length: 10,
		Search: entities.DataTablesSearch{
			Value: "",
			Regex: false,
		},
		Columns: []entities.DataTablesColumn{
			{Data: "id", Name: "", Searchable: true, Orderable: true, Search: entities.DataTablesSearch{Value: "", Regex: false}},
			{Data: "name", Name: "", Searchable: true, Orderable: true, Search: entities.DataTablesSearch{Value: "", Regex: false}},
		},
		Order: []entities.DataTablesOrder{
			{Column: 0, Dir: "asc"},
		},
	}

	// Convert to JSON to verify structure
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error marshaling request: %v\n", err)
		return
	}

	fmt.Printf("DataTables request structure: %s\n", data)
}