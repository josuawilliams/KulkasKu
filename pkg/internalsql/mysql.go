package internalsql

import (
	"database/sql"
	"fmt"
	"kulkasku/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *model.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, "Asia%2FJakarta")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error Connecting To Database")
	}

	fmt.Println("Database Connected")
	return db, nil
}