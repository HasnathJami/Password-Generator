package main

import (
	"os"

	"github.com/HasnathJami/go-echo-redis/routers"
	"github.com/HasnathJami/go-echo-redis/utils"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
  utils.CheckSimpleError(err)

  e := routers.Router()

  //e.Use(middleware.Logger())
  //e.Use(middleware.Recover())
  
  e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}