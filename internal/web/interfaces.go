package web

import "time"

type VideoMetadata struct {
	Id         string
	UploadedAt time.Time
}

type VideoMetadataService interface {
	Read(id string) (*VideoMetadata, error)
	List() ([]VideoMetadata, error)
	Create(videoId string, uploadedAt time.Time) error
}

type VideoContentService interface {
	Read(videoId string, filename string) ([]byte, error)
	Write(videoId string, filename string, data []byte) error
}
