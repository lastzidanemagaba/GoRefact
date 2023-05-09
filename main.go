package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/lastzidanemagaba/GoRefact/app"
)

var (
	Loc *time.Location
	Now time.Time
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var err error
	Loc, err = time.LoadLocation(getenv("TIME_ZONE"))
	if err != nil {
		panic(err)
	}

	Now = time.Now().In(Loc)
}

func main() {
	app.Run()
}

func getenv(key string) string {
	value, ok := godotenv.Loaded()[key]
	if !ok {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return value
}
