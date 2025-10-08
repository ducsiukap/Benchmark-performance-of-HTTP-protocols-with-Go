package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

const (
	dbUser  = "admin"
	dbPassw = "vduczz#13304"
	host    = "ltmdbtopic.ch8mi6e66rp2.ap-southeast-2.rds.amazonaws.com"
	port    = "3306"
	dbname  = "ltmdbtopic"
)

// auto run before main()
func init() {
	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassw, host, port, dbname)

	var err error
	// connect
	conn, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}

	// try
	err = conn.Ping()
	if err != nil {
		log.Fatalf("MySQL ping Error: %v", err)
	}

	log.Println("connected!")
}

// định nghĩa product
type Product struct {
	id    int     `json:"id"`
	name  string  `json:"name"`
	price float64 `json:"price"`
}
