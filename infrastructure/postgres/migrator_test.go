package postgres

import (
	"example/core/domain/model"
	"testing"
)

func TestMigrator_CreateTable(t *testing.T) {
	dbConfig := &Config{
		Host:     "localhost",
		Port:     "5432",
		Password: "",
		User:     "ivanmartinovic",
		DBName:   "test",
		SSLMode:  "disable",
	}

	db, err := NewConnection(dbConfig)
	if err != nil {
		t.Errorf("Unable to establish connection %+v", err)
	}

	migrator := NewMigrator(db)

	err = migrator.CreateTable(model.Post{})
	if err != nil {
		t.Errorf("Create table error %+v", err)
		return
	}
}

func TestMigrator_DeleteTable(t *testing.T) {

}
