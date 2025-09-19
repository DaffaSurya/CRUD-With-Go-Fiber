package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"


// 	_ "github.com/lib/pq"
// )



// func ConnectDB() *sql.DB {
// 	connStr := fmt.Sprintf(
// 		"host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	var DB *sql.DB
// 	var err error
// 	DB, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal("Error connecting to database:", err)
// 	}

// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal("Database not reachable:", err)
// 	}

// 	fmt.Println("Connected to PostgreSQL 🚀")
// 	return DB
// }

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=12345678 dbname=mahasiswa_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database unreachable: ", err)
	}

	fmt.Println("Database connected ✅")
	return db
}