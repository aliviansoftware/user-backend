package mock

import "user-backend/pkg"

type UserService struct {
	CreateUserFn      func(u *pkg.User) error
	CreateUserInvoked bool

	GetUserByUsernameFn      func(username string) (error, pkg.User)
	GetUserByUsernameInvoked bool

	LoginFn      func(c pkg.Credentials) (error, pkg.User)
	LoginInvoked bool
}

func (us *UserService) CreateUser(u *pkg.User) error {
	us.CreateUserInvoked = true
	return us.CreateUserFn(u)
}

func (us *UserService) GetUserByUsername(username string) (error, pkg.User) {
	us.GetUserByUsernameInvoked = true
	return us.GetUserByUsernameFn(username)
}

func (us *UserService) Login(c pkg.Credentials) (error, pkg.User) {
	us.LoginInvoked = true
	return us.LoginFn(c)
}
