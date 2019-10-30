package user

type Auth struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string
}

func (a *Auth) CheckAuth() bool {
	return true
}
