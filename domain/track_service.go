package domain

import (
	"errors"
	"gomusic/api"
	"gomusic/model"
	"gomusic/repository"
)

type TrackService struct {
	artistService ArtistService
	albumService  AlbumService
	trackRepository repository.TrackRepository
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

	return ts.trackRepository.createdTrack(track)
}

func calculateDuration(file string) int {
	// Todo
	return 200
}
