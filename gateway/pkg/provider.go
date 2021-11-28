package pkg

func NewService() *service {
	return new(service)
}

type ServiceProvider interface {
	Run() error
}
