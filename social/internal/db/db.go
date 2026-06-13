package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/lib/pq"
)

func checkAndCreateDB(addr string) error {
	var dbName string
	var defaultAddr string

	if strings.HasPrefix(addr, "postgres://") || strings.HasPrefix(addr, "postgresql://") {
		u, err := url.Parse(addr)
		if err != nil {
			return err
		}
		dbName = strings.TrimPrefix(u.Path, "/")
		if dbName == "" {
			dbName = "postgres"
		}
		u.Path = "/postgres"
		defaultAddr = u.String()
	} else {
		// Key-value DSN format
		fields := strings.Fields(addr)
		var newFields []string
		dbName = "postgres" // default if not specified
		for _, f := range fields {
			parts := strings.SplitN(f, "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				val := parts[1]
				if key == "dbname" {
					dbName = val
					continue
				}
			}
			newFields = append(newFields, f)
		}
		newFields = append(newFields, "dbname=postgres")
		defaultAddr = strings.Join(newFields, " ")
	}

	// Connect to default system DB 'postgres' to perform check/creation
	defaultDB, err := sql.Open("postgres", defaultAddr)
	if err != nil {
		return err
	}
	defer defaultDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	err = defaultDB.QueryRowContext(ctx, query, dbName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		// Use pq.QuoteIdentifier to prevent SQL injection in database name
		_, err = defaultDB.ExecContext(ctx, "CREATE DATABASE "+pq.QuoteIdentifier(dbName))
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDB(addr string, maxOpenConns int, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	if err := checkAndCreateDB(addr); err != nil {
		return nil, fmt.Errorf("failed to check/create database: %w", err)
	}

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	parsedMaxIdleTime, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(parsedMaxIdleTime)

	return db, nil
}
