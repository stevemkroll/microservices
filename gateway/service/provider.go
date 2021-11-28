package service

func NewService() *service {
	return new(service)
}

type Provider interface {
	Run() error
}
