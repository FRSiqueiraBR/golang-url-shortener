package services

import (
	"crypto/md5"
	"encoding/base32"
	"strconv"
	"strings"
	"time"
)

type HashService struct{}

func NewHashService() *HashService {
	return &HashService{}
}

func (service *HashService) Create(url string, ip string) (string, error) {
	now := time.Now()
	sec := now.UnixNano()

	ipWithoutDots := strings.ReplaceAll(ip, ".", "")
	stringToHash := ipWithoutDots + strconv.FormatInt(sec, 10)

	data := []byte(stringToHash)
	md5Hash := md5.Sum(data)
	
	urlShortened := base32.HexEncoding.EncodeToString(md5Hash[:])[0:7]

	return urlShortened, nil
}