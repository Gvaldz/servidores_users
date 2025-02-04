package infrastructure

import (
	"errors"
	"sync"
	entities "usuarios/src/users/domain/entities"
)

type UserRepository struct {
	users []entities.User
	mu    sync.Mutex // Para evitar condiciones de carrera
	nextID int       // Para generar IDs únicos
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []entities.User{},
		nextID: 1, // Iniciar el ID en 1
	}
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	return r.users, nil
}

func (r *UserRepository) Save(user entities.User) error {
	r.mu.Lock()         
	defer r.mu.Unlock()

	user.ID = r.nextID 
	r.nextID++         
	r.users = append(r.users, user)
	return nil
}

func (r *UserRepository) Update(user entities.User) error {
	r.mu.Lock()         
	defer r.mu.Unlock()

	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) Delete(id int32) error {
	r.mu.Lock()         
	defer r.mu.Unlock()

	for i, u := range r.users {
		if u.ID == int(id) {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}