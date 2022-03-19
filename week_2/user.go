package main

type User struct {
	Id   int
	Name string
}

type UserRepo interface {
	FindUser(string) (*User, error)
}

type UserUC struct {
	repo UserRepo
}

func NewUserUC(r UserRepo) *UserUC {
	return &UserUC{r}
}

func (u *UserUC) FindUser(name string) (*User, error) {
	return u.repo.FindUser(name)
}
