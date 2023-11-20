package main

import (
	"database/sql"

	"github.com/gsasso/go-api/internal/app/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateCustomer(*types.CustomerResponse) error
	GetCustomerById(string) (*types.CustomerResponse, error)
	DeleteCustomer(string) (*types.CustomerResponse, error)
	UpdateCustomer(string) (*types.CustomerResponse, error)
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=postgres dbname=postgres password=goapi sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (s *PostgresDB) Init() error {
	//Create database, script?
	return nil
}

func (s *PostgresDB) CreateCustomer(*types.CustomerResponse) error {
	return nil
}

func (s *PostgresDB) GetCustomerById(SystemID string) (*types.CustomerResponse, error) {
	return nil, nil
}
