package services

type service struct {}

func New() *service {
	return &service{}
}

func (srv *service) Get(id string) () {
}
