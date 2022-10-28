package repository

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
	"gomusic/model"
)

type AlbumRepository struct {
}

func (ar AlbumRepository) isExistsById(id int) bool {
	db, err := sql.Open("sqlite3", "test.db")


	// Todo: If db does not exists create it
    if err != nil {
        // log.Fatal(err)
		return false
    }

    defer db.Close()

    var album model.AlbumEntity

    row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
    err = row.Scan(&album)

    if err != nil {
        // log.Fatal(err)
		return false
    }

    return true
}
