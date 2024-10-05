package repository

type Repository interface {
	GetRecord() (string, error)
}
