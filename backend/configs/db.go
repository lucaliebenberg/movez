package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("SQL_DB_USER"),
        Passwd: os.Getenv("SQL_DB_PASS"),
        Net:    os.Getenv("SQL_NET"),
        Addr:   os.Getenv("SQL_DB_URL"),
        DBName: os.Getenv("SQL_DB_NAME"),
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}