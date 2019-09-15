package usecase

import "app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	GetAll() (domain.Users, error)
}
