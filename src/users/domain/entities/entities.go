package entities

type User struct {
	ID       int
	Username string
	Name   string
}

func newUser(id int, username, name string) *User {
	return &User{
		ID:       id,
		Username: username,
		Name:   name,
	}
}