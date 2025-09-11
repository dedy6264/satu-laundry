package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"fmt"
	"strings"
	"time"
)

type customerPostgresRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerPostgresRepository{db: db}
}

func (r *customerPostgresRepository) Create(customer *entities.Customer) error {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Build INSERT clause based on available columns
	columnList, valuePlaceholders, valueArgs := r.buildInsertClause(customer, columns)
	
	// Build the query
	query := fmt.Sprintf("INSERT INTO pelanggan (%s) VALUES (%s) RETURNING id_pelanggan", 
		strings.Join(columnList, ", "), 
		strings.Join(valuePlaceholders, ", "))
	
	// Check if id_pelanggan column exists for RETURNING clause
	hasIDColumn := false
	for _, col := range columns {
		if col == "id_pelanggan" {
			hasIDColumn = true
			break
		}
	}
	
	if !hasIDColumn {
		// Remove RETURNING clause if id_pelanggan doesn't exist
		query = fmt.Sprintf("INSERT INTO pelanggan (%s) VALUES (%s)", 
			strings.Join(columnList, ", "), 
			strings.Join(valuePlaceholders, ", "))
		
		_, err = r.db.Exec(query, valueArgs...)
		return err
	}
	
	return r.db.QueryRow(query, valueArgs...).Scan(&customer.ID)
}

// Helper function to build INSERT clause based on available columns
func (r *customerPostgresRepository) buildInsertClause(customer *entities.Customer, columns []string) ([]string, []string, []interface{}) {
	var columnList []string
	var valuePlaceholders []string
	var valueArgs []interface{}
	
	// Create a map of available columns for quick lookup
	availableColumns := make(map[string]bool)
	for _, col := range columns {
		availableColumns[col] = true
	}
	
	// Add fields that exist in the table
	if availableColumns["id_outlet"] {
		columnList = append(columnList, "id_outlet")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.OutletID)
	}
	
	if availableColumns["nama"] {
		columnList = append(columnList, "nama")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Name)
	} else if availableColumns["name"] {
		// Alternative column name
		columnList = append(columnList, "name")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Name)
	}
	
	if availableColumns["email"] {
		columnList = append(columnList, "email")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Email)
	}
	
	if availableColumns["telepon"] {
		columnList = append(columnList, "telepon")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Phone)
	} else if availableColumns["phone"] {
		// Alternative column name
		columnList = append(columnList, "phone")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Phone)
	}
	
	if availableColumns["alamat"] {
		columnList = append(columnList, "alamat")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Address)
	} else if availableColumns["address"] {
		// Alternative column name
		columnList = append(columnList, "address")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Address)
	}
	
	// Always add timestamps if they exist
	if availableColumns["created_at"] {
		columnList = append(columnList, "created_at")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, time.Now())
	}
	
	if availableColumns["updated_at"] {
		columnList = append(columnList, "updated_at")
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("$%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, time.Now())
	}
	
	return columnList, valuePlaceholders, valueArgs
}

func (r *customerPostgresRepository) FindByID(id int) (*entities.Customer, error) {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return nil, fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Build SELECT clause based on available columns
	selectClause, actualColumns, _ := r.buildSelectClause(columns)
	
	// Check if id_pelanggan column exists
	hasIDColumn := false
	for _, col := range columns {
		if col == "id_pelanggan" {
			hasIDColumn = true
			break
		}
	}
	
	var query string
	var args []interface{}
	
	if hasIDColumn {
		// Use query with ID filter
		query = fmt.Sprintf("SELECT %s FROM pelanggan WHERE id_pelanggan = $1", selectClause)
		args = append(args, id)
	} else {
		// Fallback - this is problematic but we'll try to get the first matching row
		query = fmt.Sprintf("SELECT %s FROM pelanggan", selectClause)
		// We can't really filter without knowing the column names, so we'll return an error
		return nil, fmt.Errorf("id_pelanggan column not found in table")
	}
	
	// Execute the query
	row := r.db.QueryRow(query, args...)

	customer, err := r.scanCustomerRow(row, actualColumns)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan customer: %w", err)
	}

	return customer, nil
}

func (r *customerPostgresRepository) FindByOutletID(outletID int) ([]entities.Customer, error) {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return nil, fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Build SELECT clause based on available columns
	selectClause, actualColumns, _ := r.buildSelectClause(columns)
	
	// Check if id_outlet column exists
	hasOutletColumn := false
	for _, col := range columns {
		if col == "id_outlet" {
			hasOutletColumn = true
			break
		}
	}
	
	var query string
	var args []interface{}
	
	if hasOutletColumn {
		// Use query with outlet filter
		query = fmt.Sprintf("SELECT %s FROM pelanggan WHERE id_outlet = $1", selectClause)
		args = append(args, outletID)
	} else {
		// Fallback to getting all customers
		query = fmt.Sprintf("SELECT %s FROM pelanggan", selectClause)
	}
	
	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query customers by outlet: %w", err)
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		customer, err := r.scanCustomerRow(rows, actualColumns)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer row: %w", err)
		}
		customers = append(customers, *customer)
	}

	return customers, nil
}

func (r *customerPostgresRepository) FindAll() ([]entities.Customer, error) {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return nil, fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Build SELECT clause based on available columns
	selectClause, actualColumns, _ := r.buildSelectClause(columns)
	
	// Build the query
	query := fmt.Sprintf("SELECT %s FROM pelanggan", selectClause)
	
	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all customers: %w", err)
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		customer, err := r.scanCustomerRow(rows, actualColumns)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer row: %w", err)
		}
		customers = append(customers, *customer)
	}

	return customers, nil
}

func (r *customerPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Customer, int, int, error) {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Build SELECT clause based on available columns
	selectClause, actualColumns, fieldMap := r.buildSelectClause(columns)
	
	// Build the query
	baseQuery := fmt.Sprintf("SELECT %s FROM pelanggan", selectClause)
	countQuery := "SELECT COUNT(*) FROM pelanggan"
	
	var args []interface{}
	argIndex := 1
	
	// Add search condition if provided (only for columns that exist)
	if search != "" {
		search = strings.ToLower(search)
		whereClause, searchArgs, newIndex := r.buildSearchClause(columns, search, argIndex)
		if whereClause != "" {
			baseQuery += " WHERE " + whereClause
			countQuery += " WHERE " + whereClause
			args = append(args, searchArgs...)
			argIndex = newIndex
		}
	}
	
	// Add ordering (only for columns that exist)
	dbOrderBy := fieldMap[orderBy]
	if dbOrderBy == "" {
		// Default to first available column for ordering
		if len(actualColumns) > 0 {
			dbOrderBy = actualColumns[0]
		} else {
			dbOrderBy = "id_pelanggan" // fallback
		}
	}
	
	// Validate order direction
	if orderDir != "asc" && orderDir != "desc" {
		orderDir = "asc"
	}
	
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", dbOrderBy, strings.ToUpper(orderDir))
	
	// Add pagination
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)
	
	// Execute the data query
	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to query customers: %w", err)
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		customer, err := r.scanCustomerRow(rows, actualColumns)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to scan customer row: %w", err)
		}
		customers = append(customers, *customer)
	}
	
	// Execute the count query
	var recordsTotal, recordsFiltered int
	err = r.db.QueryRow(countQuery, args[:len(args)-2]...).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count customers: %w", err)
	}
	
	// If search is applied, we need to get the filtered count
	if search != "" && len(args[:len(args)-2]) > 0 {
		searchArgs := args[:len(args)-2] // Remove limit and offset args
		err = r.db.QueryRow(countQuery, searchArgs...).Scan(&recordsFiltered)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to count filtered customers: %w", err)
		}
	} else {
		recordsFiltered = recordsTotal
	}
	
	return customers, recordsTotal, recordsFiltered, nil
}

// Helper function to get table columns
func (r *customerPostgresRepository) getTableColumns() ([]string, error) {
	// Try to get column information from information_schema
	query := `
		SELECT column_name 
		FROM information_schema.columns 
		WHERE table_name = 'pelanggan' 
		AND table_schema = 'public'
		ORDER BY ordinal_position
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var columns []string
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	
	return columns, nil
}

// Helper function to build SELECT clause based on available columns
func (r *customerPostgresRepository) buildSelectClause(columns []string) (string, []string, map[string]string) {
	// Field mapping for ordering
	fieldMap := map[string]string{
		"id_pelanggan": "id_pelanggan",
		"id_outlet":    "id_outlet",
		"nama":         "nama",
		"email":        "email",
		"telepon":      "telepon",
		"alamat":       "alamat",
		"created_at":   "created_at",
		"updated_at":   "updated_at",
	}
	
	var selectFields []string
	var actualColumns []string // The actual column names being selected
	availableColumns := make(map[string]bool)
	
	// Check which expected columns are actually available
	for _, column := range columns {
		availableColumns[column] = true
	}
	
	// Add columns that exist in the table and are expected
	if availableColumns["id_pelanggan"] {
		selectFields = append(selectFields, "id_pelanggan")
		actualColumns = append(actualColumns, "id_pelanggan")
	}
	
	if availableColumns["id_outlet"] {
		selectFields = append(selectFields, "id_outlet")
		actualColumns = append(actualColumns, "id_outlet")
	}
	
	// Handle name column (could be 'nama' or 'name')
	if availableColumns["nama"] {
		selectFields = append(selectFields, "nama")
		actualColumns = append(actualColumns, "nama")
	} else if availableColumns["name"] {
		selectFields = append(selectFields, "name")
		actualColumns = append(actualColumns, "name")
	}
	
	if availableColumns["email"] {
		selectFields = append(selectFields, "email")
		actualColumns = append(actualColumns, "email")
	}
	
	// Handle phone column (could be 'telepon' or 'phone')
	if availableColumns["telepon"] {
		selectFields = append(selectFields, "telepon")
		actualColumns = append(actualColumns, "telepon")
	} else if availableColumns["phone"] {
		selectFields = append(selectFields, "phone")
		actualColumns = append(actualColumns, "phone")
	}
	
	// Handle address column (could be 'alamat' or 'address')
	if availableColumns["alamat"] {
		selectFields = append(selectFields, "alamat")
		actualColumns = append(actualColumns, "alamat")
	} else if availableColumns["address"] {
		selectFields = append(selectFields, "address")
		actualColumns = append(actualColumns, "address")
	}
	
	if availableColumns["created_at"] {
		selectFields = append(selectFields, "created_at")
		actualColumns = append(actualColumns, "created_at")
	}
	
	if availableColumns["updated_at"] {
		selectFields = append(selectFields, "updated_at")
		actualColumns = append(actualColumns, "updated_at")
	}
	
	// If no expected columns found, just use all available columns
	if len(selectFields) == 0 {
		selectFields = columns
		actualColumns = columns
	}
	
	// If still empty, use a minimal set
	if len(selectFields) == 0 {
		selectFields = []string{"id_pelanggan"}
		actualColumns = []string{"id_pelanggan"}
	}
	
	return strings.Join(selectFields, ", "), actualColumns, fieldMap
}

// Helper function to build search clause
func (r *customerPostgresRepository) buildSearchClause(columns []string, search string, startIndex int) (string, []interface{}, int) {
	searchableColumns := []string{"nama", "name", "email", "telepon", "phone"}
	var conditions []string
	var args []interface{}
	argIndex := startIndex
	
	for _, col := range searchableColumns {
		// Check if column exists in table
		found := false
		for _, tableCol := range columns {
			if tableCol == col {
				found = true
				break
			}
		}
		
		if found {
			conditions = append(conditions, fmt.Sprintf("LOWER(%s) LIKE $%d", col, argIndex))
			args = append(args, "%"+search+"%")
			argIndex++
		}
	}
	
	if len(conditions) == 0 {
		return "", args, argIndex
	}
	
	return strings.Join(conditions, " OR "), args, argIndex
}

// Helper function to scan customer row (handles both *sql.Row and *sql.Rows)
func (r *customerPostgresRepository) scanCustomerRow(row interface{}, columns []string) (*entities.Customer, error) {
	// Create a map of column names to scan destinations
	customer := &entities.Customer{}
	
	// Create slice of interface{} for scanning - only for actual columns returned
	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	
	for i := range columns {
		valuePointers[i] = &values[i]
	}
	
	// Scan the row based on type
	switch v := row.(type) {
	case *sql.Row:
		// For single row
		if err := v.Scan(valuePointers...); err != nil {
			return nil, err
		}
	case *sql.Rows:
		// For multiple rows
		if err := v.Scan(valuePointers...); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported row type")
	}
	
	// Map scanned values to customer fields
	for i, column := range columns {
		if values[i] == nil {
			continue
		}
		
		switch column {
		case "id_pelanggan":
			if v, ok := values[i].(int64); ok {
				customer.ID = int(v)
			} else if v, ok := values[i].(int32); ok {
				customer.ID = int(v)
			} else if v, ok := values[i].(int); ok {
				customer.ID = v
			}
		case "id_outlet":
			if v, ok := values[i].(int64); ok {
				customer.OutletID = int(v)
			} else if v, ok := values[i].(int32); ok {
				customer.OutletID = int(v)
			} else if v, ok := values[i].(int); ok {
				customer.OutletID = v
			}
		case "nama", "name":
			if v, ok := values[i].(string); ok {
				customer.Name = v
			}
		case "email":
			if v, ok := values[i].(string); ok {
				customer.Email = v
			}
		case "telepon", "phone":
			if v, ok := values[i].(string); ok {
				customer.Phone = v
			}
		case "alamat", "address":
			if v, ok := values[i].(string); ok {
				customer.Address = v
			}
		case "created_at":
			if v, ok := values[i].(time.Time); ok {
				customer.CreatedAt = v
			}
		case "updated_at":
			if v, ok := values[i].(time.Time); ok {
				customer.UpdatedAt = v
			}
		}
	}
	
	return customer, nil
}

func (r *customerPostgresRepository) Update(customer *entities.Customer) error {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Build UPDATE clause based on available columns
	setClause, valueArgs := r.buildUpdateClause(customer, columns)
	
	// Check if id_pelanggan column exists
	hasIDColumn := false
	for _, col := range columns {
		if col == "id_pelanggan" {
			hasIDColumn = true
			break
		}
	}
	
	if !hasIDColumn {
		return fmt.Errorf("id_pelanggan column not found in table")
	}
	
	// Build the query
	query := fmt.Sprintf("UPDATE pelanggan SET %s WHERE id_pelanggan = $%d", 
		setClause, 
		len(valueArgs)+1)
	
	// Add the ID to the args
	valueArgs = append(valueArgs, customer.ID)
	
	_, err = r.db.Exec(query, valueArgs...)
	return err
}

// Helper function to build UPDATE clause based on available columns
func (r *customerPostgresRepository) buildUpdateClause(customer *entities.Customer, columns []string) (string, []interface{}) {
	var setParts []string
	var valueArgs []interface{}
	
	// Create a map of available columns for quick lookup
	availableColumns := make(map[string]bool)
	for _, col := range columns {
		availableColumns[col] = true
	}
	
	// Add fields that exist in the table
	if availableColumns["id_outlet"] {
		setParts = append(setParts, fmt.Sprintf("id_outlet = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.OutletID)
	}
	
	if availableColumns["nama"] {
		setParts = append(setParts, fmt.Sprintf("nama = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Name)
	} else if availableColumns["name"] {
		// Alternative column name
		setParts = append(setParts, fmt.Sprintf("name = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Name)
	}
	
	if availableColumns["email"] {
		setParts = append(setParts, fmt.Sprintf("email = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Email)
	}
	
	if availableColumns["telepon"] {
		setParts = append(setParts, fmt.Sprintf("telepon = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Phone)
	} else if availableColumns["phone"] {
		// Alternative column name
		setParts = append(setParts, fmt.Sprintf("phone = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Phone)
	}
	
	if availableColumns["alamat"] {
		setParts = append(setParts, fmt.Sprintf("alamat = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Address)
	} else if availableColumns["address"] {
		// Alternative column name
		setParts = append(setParts, fmt.Sprintf("address = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, customer.Address)
	}
	
	// Always update timestamp if it exists
	if availableColumns["updated_at"] {
		setParts = append(setParts, fmt.Sprintf("updated_at = $%d", len(valueArgs)+1))
		valueArgs = append(valueArgs, time.Now())
	}
	
	return strings.Join(setParts, ", "), valueArgs
}

func (r *customerPostgresRepository) Delete(id int) error {
	// First, let's try to determine what columns actually exist in the table
	columns, err := r.getTableColumns()
	if err != nil {
		return fmt.Errorf("failed to get table columns: %w", err)
	}
	
	// Check if id_pelanggan column exists
	hasIDColumn := false
	for _, col := range columns {
		if col == "id_pelanggan" {
			hasIDColumn = true
			break
		}
	}
	
	if !hasIDColumn {
		return fmt.Errorf("id_pelanggan column not found in table")
	}
	
	query := `DELETE FROM pelanggan WHERE id_pelanggan = $1`
	_, err = r.db.Exec(query, id)
	return err
}