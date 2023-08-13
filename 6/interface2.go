package main

var _ User = &superUser{}

type superUser struct {
	Name      string
	Age       int
	isBlocked bool
}

// ChangeAddress implements User.
// func (*superUser) ChangeAddress(newAddress string) {
// 	panic("unimplemented")
// }

// // ChangeFio implements User.
// func (*superUser) ChangeFio(newfio string) int {
// 	panic("unimplemented")
// }

func (u *superUser) Block() {
	u.isBlocked = true
}

var _ User = &user{}

type user struct {
	Adress, Fio, Phone string
	isBlocked          bool
}

// ChangeAddress implements User.
//
//	func (u *user) ChangeAddress(newAddress string) {
//		u.Adress = newAddress
//	}
//
//	func (u *user) ChangeFio(newfio string) int {
//		u.Fio = newfio
//		return 0
//	}
func (u *user) Block() {
	u.isBlocked = true
}

type User interface {
	// ChangeFio(newfio string) int
	// ChangeAddress(newAddress string)
	Block()
}

// ChangeFio implements User.

func NewUser(fio, adress, phone string) User {
	u := user{}
	return &u
}

func main() {
	us := NewUser("dasd", "sada", "asd")
	us.Block()
}
