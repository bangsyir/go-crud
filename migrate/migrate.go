package main

import (
	"github.com/bangsyir/go-crud/initializers"
	"github.com/bangsyir/go-crud/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

// open terminal in this path and run :
// go run migrate/migrate
