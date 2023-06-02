package modelo

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (usuarios *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuarios.Id = id
	usuarios.Name = name
	usuarios.CreatedAt = createdAt
	usuarios.Status = status
}
