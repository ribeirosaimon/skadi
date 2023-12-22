package repository

type Entity interface {
	GetId() any
}

type skadiInterface interface {
	FindById(any) (any, error)
}
