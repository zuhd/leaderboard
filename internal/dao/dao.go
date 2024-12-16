package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"leaderboard/internal/migrate"
)

type DAO interface {
	rankDAO
	authDAO
	scoreDAO
}

type genericDAO struct {
	db *sqlx.DB
}

func NewDAO(dsn string) (*genericDAO, error) {
	var err error
	d := &genericDAO{}
	d.db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect sql: %w", err)
	}

	d.db.SetMaxOpenConns(20)
	d.db.SetMaxIdleConns(10)

	m, err := migrate.NewMigrator(d.db, "file://internal/migrate")
	if err != nil {
		return nil, fmt.Errorf("failed to init migrate: %w", err)
	}

	err = m.MigrateUp()
	if err != nil {
		err = m.MigrateDown()
		if err != nil {
			return nil, fmt.Errorf("failed to migrate: %w", err)
		}
	}

	return d, nil
}

func NewTestDAO(dsn string) (*genericDAO, error) {
	var err error
	d := &genericDAO{}
	d.db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect sql: %w", err)
	}

	d.db.SetMaxOpenConns(20)
	d.db.SetMaxIdleConns(10)

	m, err := migrate.NewMigrator(d.db, "./migrate")
	if err != nil {
		return nil, fmt.Errorf("failed to init migrate: %w", err)
	}

	err = m.MigrateDownAndUp()
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}
	return d, nil
}
