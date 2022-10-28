package repository

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
	"gomusic/model"
)

type TrackRepository struct {
}

func (tr TrackRepository) createdTrack(track model.TrackEntity) (track model.TrackEntity, error) {
	db, err := sql.Open("sqlite3", "test.db")


	// Todo: If db does not exists create it
    if err != nil {
        // log.Fatal(err)
		return model.TrackEntity(), erorrs.New("Database does not exists")
    }

    defer db.Close()

	res := fmt.Sprintf(
		"INSERT INTO tracks(id, file, name, artist_id, album_id, duration, lyrics) VALUES (%d, %s, %s, %d, %d, %d, %s);",
		track.Id,
		track.File,
		track.Name,
		track.ArtistId,
		track.AlbumId,
		track.Duration,
		track.Lyrics
	)

    _, err = db.Exec(sts)

    if err != nil {
        return model.TrackEntity(), err
    }

    return track, nil
}
