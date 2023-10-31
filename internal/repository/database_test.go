package repository

import (
	"testing"
)

// Тестовые параметры подключения к БД (замените на свои)
var testDBConfig = Config{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "mysecretpassword",
	DBName:   "banner_rotation",
}

func TestDatabaseConnection(t *testing.T) {
	db, err := NewDatabase(testDBConfig)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
}
