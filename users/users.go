package users

import (
	"fmt"
	"go_basic/modelo"
	"time"
)

func AltaUser() {
	u := new(modelo.User)
	u.AddUser(1, "alta", time.Now(), true)
	fmt.Println(u)
}
