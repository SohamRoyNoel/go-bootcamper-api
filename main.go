package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"bootcamp.com/server/routes"
	"bootcamp.com/server/configs"
)

func main() {
	err := godotenv.Load(".env.dev");
	if err != nil {
		log.Fatal(configs.ENV_LOAD_ERROR)
	}
	port := os.Getenv("PORT");
	getRouterInstance := routes.RouterCombinerFunction();

	getRouterInstance.Run(":" + port);
}