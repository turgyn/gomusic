package repository

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
	"gomusic/model"
)

type ArtistRepository struct {
}

func (ar ArtistRepository) isExistsById(id int) bool {
	db, err := sql.Open("sqlite3", "test.db")


	// Todo: If db does not exists create it
    if err != nil {
        // log.Fatal(err)
		return false
    }

    defer db.Close()

    var artist model.ArtistEntity

    row := db.QueryRow("SELECT * FROM artists WHERE id = ?", id)
    err = row.Scan(&artist)

    if err != nil {
        // log.Fatal(err)
		return false
    }

    return true
}
