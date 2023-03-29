package main

import (
	"sale-system/src/repository"
	"sale-system/src/router"
)

func main() {
	repository.ConnectSQL()
	router.Router()
}