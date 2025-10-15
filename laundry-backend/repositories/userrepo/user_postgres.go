package userrepo

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/repositories"
)

type userPostgresRepo struct {
	repo repositories.Repositories
}

func NewUserRepo(repo repositories.Repositories) userPostgresRepo {
	return userPostgresRepo{
		repo: repo,
	}
}

func (r userPostgresRepo) FindByEmail(email string) (entities.User, error) {
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE email = $1`
	row := r.repo.Db.QueryRow(query, email)

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
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (r userPostgresRepo) Create(user entities.User) error {
	query := `INSERT INTO users (email, password, name, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	return r.repo.Db.QueryRow(query, user.Email, user.Password, user.Name, user.Role).Scan(&user.ID)
}
