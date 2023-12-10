package usecase

import "learn-clean-architecture/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
}
