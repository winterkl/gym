package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

const (
	_errorCodeConstraint       = "23503"
	_errorCodeUniqueConstraint = "23505"
)

type ErrorsCode struct {
	CodeConstraint       string
	CodeUniqueConstraint string
}

type Postgres struct {
	*bun.DB
	Errors ErrorsCode
}

func New(user, password, host, port, dbName, sslMode string) (*Postgres, error) {

	psql := &Postgres{
		Errors: ErrorsCode{
			CodeConstraint:       _errorCodeConstraint,
			CodeUniqueConstraint: _errorCodeUniqueConstraint,
		},
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbName, sslMode)
	config, err := pgx.ParseConfig(url)
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true

	sqlDB := stdlib.OpenDB(*config)
	psql.DB = bun.NewDB(sqlDB, pgdialect.New())

	return psql, nil
}
