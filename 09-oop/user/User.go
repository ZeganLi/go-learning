package user

type User struct {
	Id   int
	Name string
	Age  int
}

// go中不支持函数的重载，但又想按照传入的参数来初始化一个对象
// func NewUser(name string) *User {
// 	return &User{Name: name}
// }
//
// func NewUser(id int, name string) *User {
// 	return &User{Id: id, Name: name}
// }

func NewUser(fn func(u *User)) *User {
	u := new(User)
	fn(u)
	return u
}

func WithId(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
	}

}
