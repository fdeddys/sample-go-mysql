package main

import (
	"fmt"

	_ "com.ddabadi.connection/database"
	"com.ddabadi.connection/repository"
)

func main() {
	fmt.Println("start")

	repository.GetUser()
	repository.GetUser()
	repository.GetUser()
}
