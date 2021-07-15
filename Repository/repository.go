package Repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"links/internal/app/ds"
)

func InitDB(str string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/links")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("OK")

	// Установка данных
	/*insert, err := db.Query("INSERT INTO `link`(`small`, `big`) VALUES (222, 22) ")//, str, str)
	if err != nil {
		panic(err)
	}
	defer insert.Close()*/
	// Выборка данных

	res, err := db.Query("SELECT `big` FROM `link` WHERE `small` = ? ", str)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var l ds.Link
		err = res.Scan(&l.Link)
		if err != nil {
			panic(err)
		}
		fmt.Println(l.Link)
	}
}
