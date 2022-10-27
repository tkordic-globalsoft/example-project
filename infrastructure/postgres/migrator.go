package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Migrator struct {
	DB *gorm.DB
}

type posts struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Title   string
	Content string
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{DB: db}
}

func (m *Migrator) CreateTable(i interface{}) error {
	hasTable := m.DB.Migrator().HasTable(i)
	if hasTable {
		println("Table already exists")
		return nil
	}

	err := m.DB.Migrator().CreateTable(i)
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrator) DeleteTable(i interface{}) error {
	hasTable := m.DB.Migrator().HasTable(i)
	if !hasTable {
		println("Table does not exist")
		return nil
	}

	err := m.DB.Migrator().DropTable(i)
	if err != nil {
		return err
	}

	return nil
}

// AlterColumn alter value's `field` column' type based on schema definition
func (m *Migrator) AlterColumn(i interface{}, columnName string) error {
	err := m.DB.Migrator().AlterColumn(i, columnName)
	if err != nil {
		return err
	}
	return nil
}
