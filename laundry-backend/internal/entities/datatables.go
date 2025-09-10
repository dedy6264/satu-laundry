package entities

// DataTablesRequest represents the request structure for DataTables
type DataTablesRequest struct {
	Draw    int                `json:"draw"`
	Start   int                `json:"start"`
	Length  int                `json:"length"`
	Search  DataTablesSearch   `json:"search"`
	Columns []DataTablesColumn `json:"columns"`
	Order   []DataTablesOrder  `json:"order"`
}

// DataTablesSearch represents the search structure for DataTables
type DataTablesSearch struct {
	Value string `json:"value"`
	Regex bool   `json:"regex"`
}

// DataTablesColumn represents a column structure for DataTables
type DataTablesColumn struct {
	Data       string           `json:"data"`
	Name       string           `json:"name"`
	Searchable bool             `json:"searchable"`
	Orderable  bool             `json:"orderable"`
	Search     DataTablesSearch `json:"search"`
}

// DataTablesOrder represents the order structure for DataTables
type DataTablesOrder struct {
	Column int    `json:"column"`
	Dir    string `json:"dir"` // "asc" or "desc"
}

// DataTablesResponse represents the response structure for DataTables
type DataTablesResponse struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}
