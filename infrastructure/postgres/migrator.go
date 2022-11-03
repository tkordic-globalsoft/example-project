package postgres

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	postgres2 "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"path/filepath"
	"runtime"
)

type Migrator struct {
	DB *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{DB: db}
}

// CreateTable only for testing
func (m *Migrator) CreateTable(i interface{}) error {
	hasTable := m.DB.Migrator().HasTable(i)
	if hasTable {
		return nil
	}

	err := m.DB.Migrator().CreateTable(i)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTable only for testing
func (m *Migrator) DeleteTable(i interface{}) error {
	hasTable := m.DB.Migrator().HasTable(i)
	if !hasTable {
		return nil
	}

	err := m.DB.Migrator().DropTable(i)
	if err != nil {
		return err
	}

	return nil
}

// AlterColumn alter value's `field` column' type based on schema definition
// only for testing
func (m *Migrator) AlterColumn(i interface{}, columnName string) error {
	err := m.DB.Migrator().AlterColumn(i, columnName)
	if err != nil {
		return err
	}
	return nil
}

// ApplyMigrations new migrations are applied every time
// application starts
func (m *Migrator) ApplyMigrations() error {
	db, err := m.DB.DB()
	if err != nil {
		return err
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("Unable to locate migration files")
	}

	dir := filepath.Dir(currentFile)
	migrationsDir := filepath.Join(dir, "../../infrastructure/postgres/migrations")
	driver, err := postgres2.WithInstance(db, &postgres2.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://"+migrationsDir, "postgres", driver)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
