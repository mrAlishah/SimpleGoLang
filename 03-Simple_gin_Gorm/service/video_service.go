// Package service implements video-api servicec
package service

import (
	"simple_gin_gorm/entity"
	"simple_gin_gorm/repository"
)

type VideoService interface {
	FindAll() []entity.Video
	Get(video entity.Video) entity.Video

	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
}

type videoService struct {
	repo repository.VideoRepository
}

func NewVideoService(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repo: videoRepository,
	}
}

func (svc *videoService) FindAll() []entity.Video {
	return svc.repo.FindAll()
}

func (svc *videoService) Get(video entity.Video) entity.Video {
	return svc.repo.Get(video)
}

func (svc *videoService) Save(video entity.Video) {
	svc.repo.Save(video)
}

func (svc *videoService) Update(video entity.Video) {
	svc.repo.Update(video)
}

func (svc *videoService) Delete(video entity.Video) {
	svc.repo.Delete(video)
}
