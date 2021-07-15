package serv

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"links/internal/app/ds"
)

func GetLink(c echo.Context) error {
	link := ds.Link{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&link)
	if err != nil {
		log.Print("Can't decode", err)
		return c.String(http.StatusBadRequest, "Can't decode")
	}
	log.Println("Your link:", link)
	// сохраняет в мапку и делает маленькую ссылку
	cutlink, err := link.Save()
	if err != nil {
		return err
	}
	// кодируем в json

	return c.JSON(http.StatusOK, ds.Link{
		Link: "http://localhost:1323/" + cutlink,
	})
}

func MainServ() {
	e := echo.New()
	err := e.POST("/links", GetLink)
	if err != nil {
		log.Print("Can't get big link", err)
	}

	e.GET("/*", fallback)
	e.Logger.Fatal(e.Start(":1323"))
}

func fallback(eCtx echo.Context) error {
	// ловит запрос после /
	request := eCtx.Request().URL.Path[1:]
	log.Println(request)
	// перенаправляет запрос
	website := Redirect(request)
	if website == "" {
		return eCtx.String(http.StatusOK, "I can't find")
	}

	return eCtx.Redirect(http.StatusTemporaryRedirect, website)
}

func Redirect(req string) string {
	smallLink := ds.M[req]
	return smallLink
}
