package ds

import (
	"log"

	"links/internal/pkg/rand"
)

type Link struct {
	Link string `json:"link"`
}

// todo записать в конфиг
const count = 8

var M = make(map[string]string)

// todo проверить есть ли  у нас уже в мапке эта большая ссылка или нет
func (l *Link) Save() (string, error) {
	random := rand.StringRunes(count)
	M[random] = l.Link
	log.Print("Small link ", random)
	return random, nil
}
