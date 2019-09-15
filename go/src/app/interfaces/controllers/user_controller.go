package controllers

import (
	"app/domain"
	"app/interfaces/database"
	"app/usecase"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, nil)
}

func (controller *UserController) Show(c Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) All(c Context) {
	users, err := controller.Interactor.GetAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}
