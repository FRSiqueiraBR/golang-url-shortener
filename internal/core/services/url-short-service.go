package services

import (
	"time"

	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/domain"
	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/entities"
	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/ports"
)

type UrlShortService struct {
	repo ports.UrlShortRepository
	hash ports.HashService
}

func NewUrlShortService(repo ports.UrlShortRepository, hash ports.HashService) *UrlShortService {
	return &UrlShortService{
		repo: repo,
		hash: hash,
	}
}

func (service *UrlShortService) Save(url string, ip string, expiration time.Time) (*domain.UrlShort, error) {
	hash, err := service.hash.Create(url, ip)
	if err != nil {
		return &domain.UrlShort{}, err
	}

	urlEntity, err := entities.NewUrlShortEntity(url, hash, expiration, time.Now())
	if err != nil {
		return &domain.UrlShort{}, err
	}

	service.repo.Save(urlEntity)

	urlShort, err := domain.NewUrlShort(
		urlEntity.Long,
		urlEntity.Hash,
		urlEntity.Expiration,
		urlEntity.Timestamp)
	if err != nil {
		return &domain.UrlShort{}, err
	}

	return urlShort, nil
}

func (service *UrlShortService) FindAll() (*[]domain.UrlShort, error) {
	entities, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var domains []domain.UrlShort

	for _, ent := range *entities {
		urlShort, err := domain.NewUrlShort(ent.Long, ent.Hash, ent.Expiration, ent.Timestamp)
		if err != nil {
			return nil, err
		}

		domains = append(domains, *urlShort)
	}

	return &domains, nil
}

func (service *UrlShortService) FindByHash(hash string) (*string, error) {
	url, err := service.repo.FindByHash(hash)
	if err != nil {
		return nil, err
	}

	return url, nil
}
