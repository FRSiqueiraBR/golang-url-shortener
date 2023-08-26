package domain

import (
	"errors"
	"time"
)

type UrlShort struct {
	Url       string     `json:"url"`
	Hash       string    `json:"hash"`
	Expiration time.Time `json:"expiration"`
	Timestamp  time.Time `json:"timestamp"`
}

type UrlShortInteface interface {
	NewUrlShort(url string, hash string, expiration time.Time, timestamp time.Time)
}

func NewUrlShort(url string, hash string, expiration time.Time, timestamp time.Time) (*UrlShort, error) {
	short := &UrlShort{
		Url:       url,
		Hash:       hash,
		Expiration: expiration,
		Timestamp:  timestamp,
	}

	err := short.Validate()
	if err != nil {
		return nil, err
	}

	return short, nil
}

func (u *UrlShort) Validate() error {
	if u.Url == "" {
		return errors.New("full url is required")
	} else if u.Hash == "" {
		return errors.New("Hash is required")
	}

	return nil
}
