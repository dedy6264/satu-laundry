package usecases

import (
	"errors"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
	"laundry-backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo repositories.UserRepository
}

func NewAuthUsecase(userRepo repositories.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

func (u *authUsecase) Login(request entities.LoginRequest) (*entities.LoginResponse, error) {
	// Validasi input
	if request.Email == "" || request.Password == "" {
		return nil, errors.New("email dan password harus diisi")
	}

	// Cari user berdasarkan email
	user, err := u.userRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("email atau password salah")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	// Hilangkan password dari response
	user.Password = ""

	return &entities.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}