package Repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"links/internal/app/ds"
	"links/internal/pkg/rand"
)

type Repository struct {
	db *sql.DB
}

const count = 8

func InitDB() (*Repository, error) {
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:3306)/links")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	fmt.Println("OK")
	return &Repository{
		db: db,
	}, nil

}

// Установка данных
func (r *Repository) Add(l string) (string, error) {
	random := rand.StringRunes(count)
	insert, err := r.db.Query("INSERT INTO `link`(`small`, `big`) VALUES (?, ?) ", random, l)
	log.Print("Small link ", random)
	if err != nil {
		return "", err
	}
	defer insert.Close()
	return random, nil
}

// Выборка данных
func (r *Repository) Find(req string) string {
	res, err := r.db.Query("SELECT `big` FROM `link` WHERE `small` = ? ", req)
	if err != nil {
		panic(err)
	}
	var l ds.Link
	for res.Next() {
		err = res.Scan(&l.Link)
		if err != nil {
			panic(err)
		}
	}
	return l.Link
}
