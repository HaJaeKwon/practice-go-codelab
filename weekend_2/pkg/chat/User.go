package chat

import (
	"net"
)

type User struct {
	Conn net.Conn
	Name string
}

type UserRegistry struct {
	Conns map[net.Conn]User
	Names map[string]User
}

func NewUserRegistry() *UserRegistry {
	return &UserRegistry{
		Conns: map[net.Conn]User{},
		Names: map[string]User{},
	}
}
