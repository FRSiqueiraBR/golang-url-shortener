package repositories

import (
	"database/sql"
	"fmt"

	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/entities"
)

type UrlShortRepository struct {
	Db *sql.DB
}

func NewUrlShortRepository(db *sql.DB) *UrlShortRepository {
	return &UrlShortRepository{
		Db: db,
	}
}

func (repo *UrlShortRepository) Save(url *entities.UrlShortEntity) error {
	_, err := repo.Db.Exec("INSERT INTO short_url(long, hash, expiration, timestamp) values (?,?,?,?)", url.Long, url.Hash, url.Expiration, url.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UrlShortRepository) FindAll() (*[]entities.UrlShortEntity, error) {
	rows, err := repo.Db.Query("SELECT long, hash, expiration, timestamp FROM short_url")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var shortUrls []entities.UrlShortEntity

    for rows.Next() {
        var shortUrl entities.UrlShortEntity
        if err := rows.Scan(&shortUrl.Long, &shortUrl.Hash, &shortUrl.Expiration, &shortUrl.Timestamp); err != nil {
            return &shortUrls, err
        }
        shortUrls = append(shortUrls, shortUrl)
    }
    if err = rows.Err(); err != nil {
        return &shortUrls, err
    }
    return &shortUrls, nil
}

func (repo *UrlShortRepository) FindByHash(hash string) (*string, error){
    var url string
    // Query for a value based on a single row.
    if err := repo.Db.QueryRow("SELECT long FROM short_url where hash=?",
        hash).Scan(&url); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("URL não encontrada")
        }
    }

    return &url, nil
}