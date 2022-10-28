package domain

import "gomusic/repository"

type ArtistService struct {
	artistRepository repository.ArtistRepository
}

func (as ArtistService) isExistsById(id int) bool {
	return artistRepository.isExistsById(id)
}
