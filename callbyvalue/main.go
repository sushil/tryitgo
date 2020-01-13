package main

type User struct {
	Name string
}

// based on http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/
func main() {
	u := User{Name: "ying"}
	uptr := &u

	println("name before ", u.Name)

	Modify(uptr)
	println("after Modify ", u.Name)

	Modify0(u)
	println("after Modify0 ", u.Name)

	Modify1(u)
	println("after Modify1 ", u.Name)

	Modify2(&uptr)
	println("after Modify2 ", u.Name)

	Modify3(uptr)
	println("after Modify3 ", u.Name)

	Modify4(uptr)
	println("after Modify4 ", u.Name)

	u2 := &User{Name: "ying2"}

	println("name before ", u2.Name)

	Modify(u2)
	println("after Modify ", u2.Name)

	Modify0(u)
	println("after Modify0 ", u2.Name)

	Modify1(*u2)
	println("after Modify1 ", u2.Name)

	Modify2(&u2)
	println("after Modify2 ", u2.Name)

	Modify3(u2)
	println("after Modify3 ", u2.Name)

	Modify4(u2)
	println("after Modify4 ", u2.Name)
}

func Modify(u *User) {
	u = &User{Name: "yang"}
}

func Modify0(u User) {
	u = User{Name: "yang-0"}
}

func Modify1(u User) {
	u.Name = "yang-1"
}

func Modify2(u **User) {
	*u = &User{Name: "yang-2"}
}

func Modify3(u *User) {
	*u = User{Name: "yang-3"}
}

func Modify4(u *User) {
	u.Name = "yang-4"
}
