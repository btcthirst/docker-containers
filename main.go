package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	var err error
	// init db
	time.Sleep(3 * time.Second) // wait for db to start
	db, err = initDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// init port
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "4000"
	}
	addr := fmt.Sprintf(":%s", port)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("", health)

	e.Logger.Fatal(e.Start(addr))
}

func health(c echo.Context) error {
	now := time.Now()

	return c.JSON(http.StatusOK, now)
}

func initDB() (*sql.DB, error) {
	var (
		host, port, user, password, dbname string
		ok                                 bool
	)

	if host, ok = os.LookupEnv("PG_HOST"); !ok {
		host = "db"
	}
	if port, ok = os.LookupEnv("POSTGRES_PORT"); !ok {
		port = "5432"
	}
	if user, ok = os.LookupEnv("POSTGRES_USER"); !ok {
		user = "postgres"
	}
	if password, ok = os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		password = "password"
	}
	if dbname, ok = os.LookupEnv("POSTGRES_DB"); !ok {
		dbname = "mydb"
	}
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlConn)
	conn, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
