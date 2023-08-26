package entities

import (
	"errors"
	"time"
)

type UrlShortEntityInterface interface {
	NewUrlShortEntity(fullUrl string, hash string, expiration time.Time, timestamp time.Time)
}

type UrlShortEntity struct {
	Long    string
	Hash       string
	Expiration time.Time
	Timestamp  time.Time
}

func NewUrlShortEntity(fullUrl string, hash string, expiration time.Time, timestamp time.Time) (*UrlShortEntity, error) {
	url := &UrlShortEntity{
		Long:    fullUrl,
		Hash:       hash,
		Expiration: expiration,
		Timestamp:  timestamp,
	}

	err := url.Validate()
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (u *UrlShortEntity) Validate() error {
	if u.Long == "" {
		return errors.New("full url is required")
	} else if u.Hash == "" {
		return errors.New("Hash is required")
	}

	return nil
}
