package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func main() {
	db_host := "192.168.101.3"
	db_port := "3306"
	db_user := os.Getenv("ISUBATA_DB_USER")
	if db_user == "" {
		db_user = "root"
	}
	db_password := os.Getenv("ISUBATA_DB_PASSWORD")
	if db_password != "" {
		db_password = ":" + db_password
	}

	dsn := fmt.Sprintf("%s%s@tcp(%s:%s)/isubata?parseTime=true&loc=Local&charset=utf8mb4",
		db_user, db_password, db_host, db_port)

	log.Printf("Connecting to db: %q", dsn)
	db, _ = sqlx.Connect("mysql", dsn)

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Printf("Succeeded to connect db.")

	var name string
	var data []byte
	rows, err := db.Query(`SELECT name, data FROM image`)
	if err != sql.ErrNoRows {
		log.Fatalf("select error: %v", err)
	}
	for rows.Next() {
		err = rows.Scan(&name, &data)
		if err != nil {
			log.Fatalf("scan error: %v", err)
		}

		output := fmt.Sprintf("/home/isucon/isubata/webapp/public/icons/%s", name)
		file, err := os.Create(output)
		if err != nil {
			// Openエラー処理
		}
		defer file.Close()
		file.Write(([]byte)(data))
	}
}
