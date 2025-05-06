package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"study/greeting"
	"study/user"
)

func main() {
	greeting.SayHello()
	i := greeting.GiveMeInt()
	fmt.Println("i:", i)

	u := user.User{
		Name: "Григорий",
		// age не видно, потому что age private
	}

	u.ChangeAge(50)
	// changeName не видно, потому что changeName private
	// user.changeName

	pp.Println("user:", u)
}
