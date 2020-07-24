package repository

type User struct {}

func (u User) SaveUser(email, password string){
	Repository.Insert()
}