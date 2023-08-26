package ports

import (
	"time"

	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/entities"
	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/domain"
)

type UrlShortRepository interface {
	Save(url *entities.UrlShortEntity) error
	FindAll() (*[]entities.UrlShortEntity, error)
	FindByHash(hash string) (*string, error)
}

type UrlShortService interface {
	Save(url string, ip string, expiration time.Time) (*domain.UrlShort, error)
	FindAll() (*[]domain.UrlShort, error)
	FindByHash(hash string) (*string, error)
}