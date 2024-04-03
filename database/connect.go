package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/snowflakedb/gosnowflake"
)

type DBCredentials struct {
	Account  string
	User     string
	Password string
}

func Connect() (*sql.DB, error) {
	dbCredentials := DBCredentials{
		Account:  os.Getenv("SNOWFLAKE_ACCOUNT"),
		User:     os.Getenv("SNOWFLAKE_USER"),
		Password: os.Getenv("SNOWFLAKE_PASSWORD"),
	}

	dsn := fmt.Sprintf("%s:%s@%s", dbCredentials.User, dbCredentials.Password, dbCredentials.Account)
	db, err := sql.Open("snowflake", dsn)

	if err != nil {
		log.Fatal("Error connecting to Snowflake")
	}

	return db, nil
}

func Query(db *sql.DB, query string) ([]map[string]interface{}, error) {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error running query: " + query)
	}
	defer rows.Close()

	var results []map[string]interface{}
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			log.Fatal("Error scanning row" + err.Error())
		}

		rowMap := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			rowMap[colName] = *val
		}

		results = append(results, rowMap)
	}

	return results, nil
}
