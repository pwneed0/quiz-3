package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"mini-project/database"
	"mini-project/routers"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load env")
	} else {
		fmt.Println("succes load file env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed!")
		panic(err)
	} else {
		fmt.Println("DB Connetion Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	routers.StartServer()
}
