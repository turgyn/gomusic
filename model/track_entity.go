package model

type TrackEntity struct {
	Id       int
	File     string
	Name     string
	ArtistId int
	AlbumId  int
	Duration int // track duration in seconds
	Lyrics   string
}
