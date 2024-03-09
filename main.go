package main

import (
	"database/sql"
	"fmt"
	"os"
	"practice/controllers"
	"practice/database"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV CONFIGURATION
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println(`Failed Load File Environment`)
	} else {
		fmt.Println(`Success Load File Environment`)
	}

	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println(`DB Connection Failed`)
		panic(err)
	} else {
		fmt.Println(`DB Connection Success`)
	}

	database.DbConnectionFunc(DB)

	defer DB.Close()

	// Router GIN
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run("localhost:8080")
}
