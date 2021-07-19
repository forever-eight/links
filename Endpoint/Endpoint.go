package Endpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"links/Repository"
	"links/internal/app/ds"
)

type Endpoint struct {
	R *Repository.Repository
}

func (e *Endpoint) Fallback(eCtx echo.Context) error {
	// ловит запрос после /
	request := eCtx.Request().URL.Path[1:]
	log.Println(request)
	// перенаправляет запрос
	website := e.Redirect(request)
	if website == "" {
		return eCtx.String(http.StatusOK, "I can't find")
	}

	return eCtx.Redirect(http.StatusTemporaryRedirect, website)
}

func (e *Endpoint) Redirect(req string) string {
	smallLink := e.R.Find(req)
	return smallLink
}

func (ep *Endpoint) GetLink(c echo.Context) error {
	link := ds.Link{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&link)
	if err != nil {
		log.Print("Can't decode", err)
		return c.String(http.StatusBadRequest, "Can't decode")
	}
	log.Println("Your link:", link)
	// сохраняет в бд и делает маленькую ссылку
	cutlink, err := ep.R.Add(link.Link)
	//cutlink, err := link.Save()
	if err != nil {
		return err
	}
	// кодируем в json

	return c.JSON(http.StatusOK, ds.Link{
		Link: "http://localhost:1323/" + cutlink,
	})
}
