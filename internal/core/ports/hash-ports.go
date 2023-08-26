package ports

type HashService interface {
	Create(url string, ip string) (string, error)
}