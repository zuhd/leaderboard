package migrate

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

type Migrator struct {
	m              *migrate.Migrate
	migrationsPath string
}

func NewMigrator(db *sqlx.DB, migrationsPath string) (*Migrator, error) {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "mysql", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return &Migrator{m: m, migrationsPath: migrationsPath}, nil
}

func (m *Migrator) MigrateUp() error {
	if err := m.m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to up migrations: %w", err)
	}

	return nil
}

func (m *Migrator) MigrateDown() error {
	if err := m.m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to down migrations: %w", err)
	}

	return nil
}

func (m *Migrator) MigrateDownAndUp() error {
	if err := m.m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations: %w", err)
	}

	if err := m.m.Up(); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
