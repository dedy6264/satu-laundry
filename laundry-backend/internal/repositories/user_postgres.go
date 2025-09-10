package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
)

type userPostgresRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) FindByEmail(email string) (*entities.User, error) {
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user entities.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userPostgresRepository) Create(user *entities.User) error {
	query := `INSERT INTO users (email, password, name, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	return r.db.QueryRow(query, user.Email, user.Password, user.Name, user.Role).Scan(&user.ID)
}