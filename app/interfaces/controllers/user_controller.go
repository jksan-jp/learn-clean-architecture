package controllers

import (
	"learn-clean-architecture/interfaces/database"
	"learn-clean-architecture/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHander database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHander,
			},
		},
	}
}
