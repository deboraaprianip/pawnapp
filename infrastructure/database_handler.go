package infrastructure

import (
	"database/sql"
	"log"
	"pawn-clean/config"
)

func NewSqlHandler(db *config.Db) *sql.DB {
	dsn := db.DbUser + `:` + db.DbPass + `@tcp(` + db.DbHost + `:` + db.DbPort + `)/` + db.DbName + `?parseTime=1&loc=Asia%2FJakarta`
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return dbConn
}
