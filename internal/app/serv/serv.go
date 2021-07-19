package serv

import (
	"log"

	"github.com/labstack/echo"

	"links/Endpoint"
	"links/Repository"
)

func MainServ() {
	e := echo.New()
	rep, err := Repository.InitDB()
	if err != nil {
		//todo: err
		log.Println("InitDB err", err)
		return
	}

	ep := Endpoint.Endpoint{
		R: rep,
	}

	e.POST("/links", ep.GetLink)

	e.GET("/*", ep.Fallback)
	e.Logger.Fatal(e.Start(":1323"))
}
