package bot

import "fmt"

type User struct {
	Address
}

type Address struct {
	City   string
	Street string
}

func (u *User) ShowAddress() {
	u.Show()
}

func (a *Address) Show() {
	fmt.Println(a)
}
