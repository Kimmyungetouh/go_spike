package main

type User struct {
	username string
	password string
	status   bool
}

type Profile struct {
	user  User
	age   int
	hobby string
}
type Authentication interface {
	login() bool
	logout() bool
	checkStatus() bool
}

func (user *User) login() bool {
	user.status = true
	return user.status
}

func (user *User) logout() bool {
	user.status = false
	return user.status
}

func (user *User) checkStatus() bool {
	return user.status
}

func signin(auth Authentication) {
	auth.login()
}

func signout(auth Authentication) {
	auth.logout()
}

func checkStatus(auth Authentication) bool {
	return auth.checkStatus()
}

func (profile Profile) whoami() string {
	return "*" + profile.user.username + "+, and i like to " + profile.hobby
}
