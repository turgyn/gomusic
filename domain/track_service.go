package domain

import (
	"errors"
	"gomusic/api"
	"gomusic/model"
)

type TrackService struct {
	artistService ArtistService
	albumService  AlbumService
}

func (ts TrackService) CreateTrack(dto api.CreateTrackDto) (model.TrackEntity, error) {
	if !ts.artistService.isExistsById(dto.ArtistId) {
		return model.TrackEntity{}, errors.New("unknown artist id")
	}

	track := model.TrackEntity{
		Id:       -1,
		File:     dto.File,
		Name:     dto.Name,
		ArtistId: dto.ArtistId,
		AlbumId:  dto.AlbumId,
		Duration: calculateDuration(dto.File), Lyrics: dto.Lyrics}

	// todo: save and update track (after connection with db)
	return track, nil
}

func calculateDuration(file string) int {
	// Todo
	return 200
}
