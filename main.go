package main

import (
	"github.com/mariogarcia/kpia/api"
)

func main() {
	api := api.API{}
	api.InitDB("postgres://kpia:kpia@localhost/kpia")
	api.InitHandlers()
	api.Startup(":8080")
}
