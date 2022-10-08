package postgres

import (
	_ "github.com/lib/pq"
	// imports the driver, don't remove this comment, golint requires.
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const postgresDriverName = "postgres"

// New returns a postgres connection.
func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgresDriverName, datasource, opts...)
}
