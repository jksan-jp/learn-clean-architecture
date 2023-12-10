package usecase

import "learn-clean-architecture/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}
