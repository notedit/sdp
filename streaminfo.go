package sdp

import (
	"strings"
)

type StreamInfo struct {
	id     string
	tracks []*TrackInfo
}

func NewStreamInfo(streamID string) *StreamInfo {

	return &StreamInfo{
		id:     streamID,
		tracks: make([]*TrackInfo, 0),
	}
}

func (s *StreamInfo) Clone() *StreamInfo {
	stream := &StreamInfo{
		id:     s.id,
		tracks: make([]*TrackInfo, 0),
	}

	for k, v := range s.tracks {
		stream.tracks[k] = v.Clone()
	}
	return stream
}

func (s *StreamInfo) GetID() string {

	return s.id
}

func (s *StreamInfo) AddTrack(track *TrackInfo) {

	s.tracks = append(s.tracks, track)
}

func (s *StreamInfo) RemoveTrack(track *TrackInfo) {
	tracks := make([]*TrackInfo,0)
	for _,temtrack := range s.tracks {
		if temtrack.GetID() != track.GetID() {
			tracks = append(tracks,temtrack)
		}
	}
	s.tracks = tracks
}

func (s *StreamInfo) RemoveTrackById(trackId string) {
	tracks := make([]*TrackInfo,0)
	for _,temtrack := range s.tracks {
		if temtrack.GetID() != trackId {
			tracks = append(tracks,temtrack)
		}
	}
	s.tracks = tracks
}

func (s *StreamInfo) GetFirstTrack(mediaType string) *TrackInfo {

	for _, trak := range s.tracks {
		if strings.ToLower(trak.GetMediaType()) == strings.ToLower(mediaType) {

			return trak
		}
	}
	return nil
}

func (s *StreamInfo) GetTracks() []*TrackInfo {

	return s.tracks
}

func (s *StreamInfo) RemoveAllTracks() {

	s.tracks = make([]*TrackInfo,0)
}

func (s *StreamInfo) GetTrack(trackID string) *TrackInfo {

	for _,track := range s.tracks {
		if track.GetID() == trackID {
			return track
		}
	}

	return nil
}
