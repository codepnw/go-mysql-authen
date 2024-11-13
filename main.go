package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const envFile = "dev.env"

func main() {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Run(":" + os.Getenv("APP_PORT"))
}
