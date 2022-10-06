package api

type CreateTrackDto struct {
	File     string
	Name     string
	ArtistId int
	AlbumId  int
	Lyrics   string
}
