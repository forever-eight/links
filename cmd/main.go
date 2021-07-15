package main

import (
	"links/Repository"
	"links/internal/app/serv"
)

func main() {
	Repository.InitDB("111111111")
	serv.MainServ()
}
