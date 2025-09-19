package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userAccessPostgresRepository struct {
	db *sql.DB
}

func NewUserAccessRepository(db *sql.DB) UserAccessRepository {
	return &userAccessPostgresRepository{
		db: db,
	}
}

func (r *userAccessPostgresRepository) Create(access *entities.UserAccess) error {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(access.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO user_access (username, password, role, is_active, reference_level, reference_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id_access`

	err = r.db.QueryRow(query, access.Username, hashedPassword, access.Role, access.IsActive, access.ReferenceLevel, access.ReferenceID).
		Scan(&access.ID)
	if err != nil {
		return err
	}

	access.Password = "" // Clear password from struct for security
	return nil
}

func (r *userAccessPostgresRepository) FindByID(id int) (*entities.UserAccess, error) {
	query := `
		SELECT id_access, username, role, is_active, last_login, 
		       COALESCE(reference_level,''), COALESCE(reference_id,0), created_at, updated_at
		FROM user_access
		WHERE id_access = $1`

	var access entities.UserAccess
	var lastLogin sql.NullTime
	err := r.db.QueryRow(query, id).Scan(
		&access.ID,
		&access.Username,
		&access.Role,
		&access.IsActive,
		&lastLogin,
		&access.ReferenceLevel,
		&access.ReferenceID,
		&access.CreatedAt,
		&access.UpdatedAt,
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

func (r *userAccessPostgresRepository) FindByUsername(username string) (*entities.UserAccess, error) {
	query := `
		SELECT id_access, username, password, role, is_active, 
		       last_login, COALESCE(reference_level,''), COALESCE(reference_id,0), created_at, updated_at
		FROM user_access
		WHERE username = $1 AND is_active = true`

	var access entities.UserAccess
	var lastLogin sql.NullTime
	var password string
	err := r.db.QueryRow(query, username).Scan(
		&access.ID,
		&access.Username,
		&password,
		&access.Role,
		&access.IsActive,
		&lastLogin,
		&access.ReferenceLevel,
		&access.ReferenceID,
		&access.CreatedAt,
		&access.UpdatedAt,
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

func (r *userAccessPostgresRepository) FindAll() ([]entities.UserAccess, error) {
	query := `
		SELECT id_access, username, role, is_active, last_login, 
		       COALESCE(reference_level,''), COALESCE(reference_id,0), created_at, updated_at
		FROM user_access
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses []entities.UserAccess
	for rows.Next() {
		var access entities.UserAccess
		var lastLogin sql.NullTime
		err := rows.Scan(
			&access.ID,
			&access.Username,
			&access.Role,
			&access.IsActive,
			&lastLogin,
			&access.ReferenceLevel,
			&access.ReferenceID,
			&access.CreatedAt,
			&access.UpdatedAt,
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

func (r *userAccessPostgresRepository) FindAllWithPagination(limit, offset int) ([]entities.UserAccess, int, error) {
	// Count query
	countQuery := `SELECT COUNT(*) FROM user_access`

	var totalCount int
	err := r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	dataQuery := `
		SELECT id_access, username, role, is_active, last_login, 
		       COALESCE(reference_level,''), COALESCE(reference_id,0), created_at, updated_at
		FROM user_access
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(dataQuery, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var accesses []entities.UserAccess
	for rows.Next() {
		var access entities.UserAccess
		var lastLogin sql.NullTime
		err := rows.Scan(
			&access.ID,
			&access.Username,
			&access.Role,
			&access.IsActive,
			&lastLogin,
			&access.ReferenceLevel,
			&access.ReferenceID,
			&access.CreatedAt,
			&access.UpdatedAt,
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

func (r *userAccessPostgresRepository) Update(access *entities.UserAccess) error {
	query := `
		UPDATE user_access 
		SET username = $1, role = $2, is_active = $3, reference_level = $4, reference_id = $5, updated_at = NOW()
		WHERE id_access = $6`

	_, err := r.db.Exec(query, access.Username, access.Role, access.IsActive, access.ReferenceLevel, access.ReferenceID, access.ID)
	return err
}

func (r *userAccessPostgresRepository) UpdatePassword(id int, password string) error {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		UPDATE user_access 
		SET password = $1, updated_at = NOW()
		WHERE id_access = $2`

	_, err = r.db.Exec(query, hashedPassword, id)
	return err
}

func (r *userAccessPostgresRepository) UpdateLastLogin(id int) error {
	query := `
		UPDATE user_access 
		SET last_login = NOW()
		WHERE id_access = $1`

	_, err := r.db.Exec(query, id)
	return err
}

func (r *userAccessPostgresRepository) Delete(id int) error {
	query := `DELETE FROM user_access WHERE id_access = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// AuthenticateUser checks if the provided username and password are valid
func (r *userAccessPostgresRepository) AuthenticateUser(username, password string) (*entities.UserAccess, error) {
	userAccess, err := r.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if userAccess == nil {
		return nil, nil // No user found with this username
	}

	// Compare the provided password with the hashed password
	aa, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("::::::", string(aa), "::::", userAccess.Password, ":::::", password)

	err = bcrypt.CompareHashAndPassword([]byte(userAccess.Password), []byte(password))
	if err != nil {
		log.Println("Password doesn't match")
		return nil, nil // Password doesn't match
	}

	// Clear password from struct for security
	userAccess.Password = ""

	// Update last login time
	err = r.UpdateLastLogin(userAccess.ID)
	if err != nil {
		log.Println("update last Login failed")

		// Log error but don't fail the authentication
		// In a real application, you would log this properly
	}

	return userAccess, nil
}
