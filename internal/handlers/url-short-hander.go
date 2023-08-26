package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/ports"
	"github.com/go-chi/chi/v5"
)

type UrlShortHandler struct {
	svc ports.UrlShortService
}

func NewUrlShortHandler(service ports.UrlShortService) *UrlShortHandler {
	return &UrlShortHandler{
		svc: service,
	}
}

func (handler *UrlShortHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}

	expiration, err := convertStringToTime(data["expiration"].(string))
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	handler.svc.Save(data["url"].(string), r.Header.Get("x-real-ip"), expiration)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *UrlShortHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	domains, err := handler.svc.FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domains)
}

func (handler *UrlShortHandler) FindByHash(w http.ResponseWriter, r *http.Request) {
	url, err := handler.svc.FindByHash(chi.URLParam(r, "hash"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)
}

func convertStringToTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	expiration, err := time.Parse(layout, str)
	if err != nil {
		return time.Now(), errors.New("Expiration in invalid format")
	}

	return expiration, err
}
