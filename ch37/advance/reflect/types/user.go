package types

import "fmt"

type User struct {
	Name string
	age  int
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) ChangeName(name string) {
	u.Name = name
}

func (u User) private() {
	fmt.Println("private")
}
