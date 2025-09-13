package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"

	"golang.org/x/crypto/bcrypt"
)

type employeeAccessPostgresRepository struct {
	db *sql.DB
}

func NewEmployeeAccessRepository(db *sql.DB) EmployeeAccessRepository {
	return &employeeAccessPostgresRepository{
		db: db,
	}
}

func (r *employeeAccessPostgresRepository) Create(access *entities.EmployeeAccess) error {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(access.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO employee_access (id_pegawai, username, password, role, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id_access`

	err = r.db.QueryRow(query, access.EmployeeID, access.Username, hashedPassword, access.Role, access.IsActive).
		Scan(&access.ID)
	if err != nil {
		return err
	}

	access.Password = "" // Clear password from struct for security
	return nil
}

func (r *employeeAccessPostgresRepository) FindByID(id int) (*entities.EmployeeAccess, error) {
	query := `
		SELECT ea.id_access, ea.id_pegawai, ea.username, ea.role, ea.is_active, ea.last_login, 
		       ea.created_at, ea.updated_at,
		       p.nama_lengkap, p.email, p.id_outlet
		FROM employee_access ea
		JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
		WHERE ea.id_access = $1`

	var access entities.EmployeeAccess
	var lastLogin sql.NullTime
	err := r.db.QueryRow(query, id).Scan(
		&access.ID,
		&access.EmployeeID,
		&access.Username,
		&access.Role,
		&access.IsActive,
		&lastLogin,
		&access.CreatedAt,
		&access.UpdatedAt,
		&access.EmployeeName,
		&access.EmployeeEmail,
		&access.OutletID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if lastLogin.Valid {
		access.LastLogin = &lastLogin.Time
	}

	return &access, nil
}

func (r *employeeAccessPostgresRepository) FindByUsername(username string) (*entities.EmployeeAccess, error) {
	query := `
		SELECT ea.id_access, ea.id_pegawai, ea.username, ea.password, ea.role, ea.is_active, 
		       ea.last_login, ea.created_at, ea.updated_at,
		       p.nama_lengkap, p.email, p.id_outlet
		FROM employee_access ea
		JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
		WHERE ea.username = $1 AND ea.is_active = true`

	var access entities.EmployeeAccess
	var lastLogin sql.NullTime
	var password string
	err := r.db.QueryRow(query, username).Scan(
		&access.ID,
		&access.EmployeeID,
		&access.Username,
		&password,
		&access.Role,
		&access.IsActive,
		&lastLogin,
		&access.CreatedAt,
		&access.UpdatedAt,
		&access.EmployeeName,
		&access.EmployeeEmail,
		&access.OutletID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if lastLogin.Valid {
		access.LastLogin = &lastLogin.Time
	}

	// Set the password for authentication purposes
	access.Password = password

	return &access, nil
}

func (r *employeeAccessPostgresRepository) FindAll() ([]entities.EmployeeAccess, error) {
	query := `
		SELECT ea.id_access, ea.id_pegawai, ea.username, ea.role, ea.is_active, ea.last_login, 
		       ea.created_at, ea.updated_at,
		       p.nama_lengkap, p.email, p.id_outlet
		FROM employee_access ea
		JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
		ORDER BY ea.created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses []entities.EmployeeAccess
	for rows.Next() {
		var access entities.EmployeeAccess
		var lastLogin sql.NullTime
		err := rows.Scan(
			&access.ID,
			&access.EmployeeID,
			&access.Username,
			&access.Role,
			&access.IsActive,
			&lastLogin,
			&access.CreatedAt,
			&access.UpdatedAt,
			&access.EmployeeName,
			&access.EmployeeEmail,
			&access.OutletID,
		)
		if err != nil {
			return nil, err
		}

		if lastLogin.Valid {
			access.LastLogin = &lastLogin.Time
		}

		accesses = append(accesses, access)
	}

	return accesses, nil
}

func (r *employeeAccessPostgresRepository) FindAllWithPagination(limit, offset int) ([]entities.EmployeeAccess, int, error) {
	// Count query
	countQuery := `
		SELECT COUNT(*) 
		FROM employee_access ea
		JOIN pegawai p ON ea.id_pegawai = p.id_pegawai`

	var totalCount int
	err := r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	dataQuery := `
		SELECT ea.id_access, ea.id_pegawai, ea.username, ea.role, ea.is_active, ea.last_login, 
		       ea.created_at, ea.updated_at,
		       p.nama_lengkap, p.email, p.id_outlet
		FROM employee_access ea
		JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
		ORDER BY ea.created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(dataQuery, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var accesses []entities.EmployeeAccess
	for rows.Next() {
		var access entities.EmployeeAccess
		var lastLogin sql.NullTime
		err := rows.Scan(
			&access.ID,
			&access.EmployeeID,
			&access.Username,
			&access.Role,
			&access.IsActive,
			&lastLogin,
			&access.CreatedAt,
			&access.UpdatedAt,
			&access.EmployeeName,
			&access.EmployeeEmail,
			&access.OutletID,
		)
		if err != nil {
			return nil, 0, err
		}

		if lastLogin.Valid {
			access.LastLogin = &lastLogin.Time
		}

		accesses = append(accesses, access)
	}

	return accesses, totalCount, nil
}

func (r *employeeAccessPostgresRepository) Update(access *entities.EmployeeAccess) error {
	query := `
		UPDATE employee_access 
		SET username = $1, role = $2, is_active = $3, updated_at = NOW()
		WHERE id_access = $4`

	_, err := r.db.Exec(query, access.Username, access.Role, access.IsActive, access.ID)
	return err
}

func (r *employeeAccessPostgresRepository) UpdatePassword(id int, password string) error {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		UPDATE employee_access 
		SET password = $1, updated_at = NOW()
		WHERE id_access = $2`

	_, err = r.db.Exec(query, hashedPassword, id)
	return err
}

func (r *employeeAccessPostgresRepository) UpdateLastLogin(id int) error {
	query := `
		UPDATE employee_access 
		SET last_login = NOW()
		WHERE id_access = $1`

	_, err := r.db.Exec(query, id)
	return err
}

func (r *employeeAccessPostgresRepository) Delete(id int) error {
	query := `DELETE FROM employee_access WHERE id_access = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *employeeAccessPostgresRepository) FindByOutletID(outletID int) ([]entities.EmployeeAccess, error) {
	query := `
		SELECT ea.id_access, ea.id_pegawai, ea.username, ea.role, ea.is_active, ea.last_login,
		       ea.created_at, ea.updated_at,
		       p.nama_lengkap, p.email
		FROM employee_access ea
		JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
		WHERE p.id_outlet = $1
		ORDER BY p.nama_lengkap`

	rows, err := r.db.Query(query, outletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses []entities.EmployeeAccess
	for rows.Next() {
		var access entities.EmployeeAccess
		var lastLogin sql.NullTime
		err := rows.Scan(
			&access.ID,
			&access.EmployeeID,
			&access.Username,
			&access.Role,
			&access.IsActive,
			&lastLogin,
			&access.CreatedAt,
			&access.UpdatedAt,
			&access.EmployeeName,
			&access.EmployeeEmail,
		)
		if err != nil {
			return nil, err
		}

		if lastLogin.Valid {
			access.LastLogin = &lastLogin.Time
		}

		accesses = append(accesses, access)
	}

	return accesses, nil
}

// AuthenticateEmployee checks if the provided username and password are valid
func (r *employeeAccessPostgresRepository) AuthenticateEmployee(username, password string) (*entities.EmployeeAccess, error) {
	employeeAccess, err := r.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if employeeAccess == nil {
		return nil, nil // No employee found with this username
	}

	// Compare the provided password with the hashed password
	aa, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("::::::", string(aa), "::::", employeeAccess.Password, ":::::", password)

	err = bcrypt.CompareHashAndPassword([]byte(employeeAccess.Password), []byte(password))
	if err != nil {
		return nil, nil // Password doesn't match
	}

	// Clear password from struct for security
	employeeAccess.Password = ""

	// Update last login time
	err = r.UpdateLastLogin(employeeAccess.ID)
	if err != nil {
		// Log error but don't fail the authentication
		// In a real application, you would log this properly
	}

	return employeeAccess, nil
}
