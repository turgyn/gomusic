package domain

import "gomusic/repository"

type AlbumService struct {
	albumRepository repository.AlbumRepository
}

func (as AlbumService) isExistsById(id int) bool {
	return albumRepository.isExistsById(id)
}
