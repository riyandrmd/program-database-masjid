package main

import (
	"masjiddotexe/routers"

	"masjiddotexe/connection"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectToDb()
	rh := routers.HandlerStr{
		DB: db,
		R:  r,
	}

	rh.Routers()
	r.Run()
}
