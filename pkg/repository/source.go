package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"testing-system-api/models"
)

func NewTestingSystemDatabase(config models.ConfigService, environment models.Environment) *sqlx.DB {
	fmt.Println("start database connected")
	database, err := NewPostgresDB(&PostgresDBConfig{
		Host:     config.TestingSystemDB.Host,
		Port:     config.TestingSystemDB.Port,
		Username: config.TestingSystemDB.Username,
		Password: environment.DBPassword,
		DBName:   config.TestingSystemDB.DBName,
		SSLMode:  config.TestingSystemDB.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize business db: %s", err.Error())
	}
	fmt.Println("database connected")
	return database
}
