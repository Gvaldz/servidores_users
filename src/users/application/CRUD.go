package application

import (
	entities "usuarios/src/users/domain/entities"
	domain "usuarios/src/users/domain"
)

type CreateUser struct {
	repo domain.IUser
}

func NewCreateUser(repo domain.IUser) *CreateUser {
	return &CreateUser{repo: repo}
}

func (cp *CreateUser) Execute(user entities.User) error {
	return cp.repo.Save(user)
}

type UpdateUser struct {
	repo domain.IUser
}

func NewUpdateUser(repo domain.IUser) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (up *UpdateUser) Execute(user entities.User) error {
	return up.repo.Update(user)
}

type ListUsers struct {
	repo domain.IUser
}

func NewListUsers(repo domain.IUser) *ListUsers {
	return &ListUsers{repo: repo}
}

func (lp *ListUsers) Execute() ([]entities.User, error) {
	return lp.repo.GetAll()
}

type DeleteUser struct {
	repo domain.IUser
}

func NewDeleteUser(repo domain.IUser) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (dp *DeleteUser) Execute(id int32) error {
	return dp.repo.Delete(id)
}
