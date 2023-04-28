package main

import (
	"jwt-go/router"
)

const PORT = ":8086"

func main() {
	router.ReadyServer().Run(PORT)
}
