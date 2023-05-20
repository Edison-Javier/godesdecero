package user

import (
	"fmt"
	"time"

	"github.com/godesdecero/models"
)

func ValidarUsuario() {
	u := new(models.User)
	u.AddUser(18401608, "Edison", time.Now(), true)
	fmt.Println(u)
}
