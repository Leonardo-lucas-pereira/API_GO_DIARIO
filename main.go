package main

import (
	"fmt"

	"api/database"
	"api/server"
)

func main() {
	fmt.Println("Iniciando servidor....")
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
