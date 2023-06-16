package db

import (
	"database/sql"
	"fmt"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/repository"
	_ "github.com/lib/pq"
)

type Handler struct {
	Db             *sql.DB
	UserRepository repository.UserRepository
}

func InitDB(s string) (*sql.DB, Handler) {
	db, err := sql.Open("postgres", s)
	if err != nil {
		fmt.Println("Failed to connect to DB", err)
	}

	return db, Handler{
		Db:             db,
		UserRepository: repository.NewUserRepository(db),
	}
}
