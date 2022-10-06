package api

import "gomusic/domain"

type TrackController struct {
	trackService domain.TrackService
}

func (tc TrackController) createTrack(dto CreateTrackDto) {
	tc.trackService.CreateTrack(dto)
}
