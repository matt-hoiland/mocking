package user

import "github.com/matt-hoiland/mocking/simple/doer"

type User struct {
	Service doer.Doer
}

func (u *User) AccessService() {
	u.Service.Do("five", 5)
}
