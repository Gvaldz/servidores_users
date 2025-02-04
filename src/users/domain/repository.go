package domain

import(
	"usuarios/src/users/domain/entities"
)

type IUser interface {
	GetAll() ([]entities.User, error)
	Save(User entities.User) error
	Update(User entities.User) error
	Delete(id int32) error 
}