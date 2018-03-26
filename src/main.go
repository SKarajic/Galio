package main

import (
	"github.com/joho/godotenv"
	"./galio/regions"
	"./galio"
	"log"
	"fmt"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("API_KEY")
	Galio := galio.New(key, regions.EUW)

	fmt.Println(Galio.GetChampions())
}
