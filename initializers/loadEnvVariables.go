package initializers

import (
  "log"
  "github.com/joho/godotenv"
)

func LoadEnvVariables(){
	  err := godotenv.Load()

  if err != nil {
    // In containers/production, vars come from the OS environment, not a .env file.
    log.Println("No .env file found, relying on environment variables")
  }
}