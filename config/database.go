package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	// format: username:password@tcp(host:port)/dbname
	dsn := "root:password@tcp(127.0.0.1:3306)/lc3"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Gagal connect ke DB:", err)
	}

	log.Println("✅ Connected to MySQL")
	return db
}
