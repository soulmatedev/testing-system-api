package repository

import (
	"github.com/jmoiron/sqlx"
	"testing-system-api/models"
)

type Sources struct {
	TestingSystemDB *sqlx.DB
}

type Auth interface {
}

type Account interface {
	Get(email string) (*models.Account, error)
}

type Repository struct {
	Account
	Auth
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Account: NewAccountPostgres(sources.TestingSystemDB),
		Auth:    NewAuthPostgres(sources.TestingSystemDB),
	}
}
