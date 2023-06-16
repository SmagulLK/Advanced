package repository

import (
	"database/sql"
	"fmt"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/models"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/utils"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	ValidateCredentials(email, password string) (*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(user *models.User) error {
	_, err := r.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow("SELECT email, password FROM users WHERE email = $1", email)
	user := models.User{}
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepo) ValidateCredentials(email, password string) (*models.User, error) {
	user, err := r.FindByEmail(email)
	if err != nil {
		fmt.Println("Error while fetching user from database")
		return nil, err
	}
	if user == nil {
		fmt.Println("User not found")

		return nil, nil // User not found
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		fmt.Println("Invalid password")
		return nil, nil // Invalid password
	}
	return user, nil
}
