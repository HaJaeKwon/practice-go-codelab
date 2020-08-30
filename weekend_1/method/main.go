package main

import "fmt"

func main() {
	user := User{
		ID: "hjk",
		Name: "jaekwon",
	}
	fmt.Printf("%p\n", &user)
	user.Print()
	user.ChangeID("123")

	name := Name("Hello World")
	name.Print()
}

type User struct {
	ID   string
	Name string
}

func (u User) Print() {
	fmt.Printf("%p\n", &u)
	fmt.Println("ID:", u.ID)
	fmt.Println("Names:", u.Name)
}

func (u *User) ChangeID(new string)  {
	fmt.Printf("%p\n", u)
	(*u).ID = new
	fmt.Println("ID:", u.ID)
	fmt.Println("Names:", u.Name)
}

type Name string

func (n Name) Print() {
	fmt.Println("Names:", n)
}
