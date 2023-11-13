package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateCustomer(*CustomerResponse) error
	GetCustomerById(string) (*CustomerResponse, error)
	DeleteCustomer(string) (*CustomerResponse, error)
	UpdateCustomer(string) (*CustomerResponse, error)
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

func (s *PostgresDB) CreateCustomer(*CustomerResponse) error {
	return nil
}

func (s *PostgresDB) GetCustomerById(SystemID string) (*CustomerResponse, error) {
	return nil, nil
}
