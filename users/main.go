package main

import "fmt"

func main() {
	fmt.Println("Welcome into our app !")
	fmt.Println("Let's register  .... ")
	guess := User{username: "johndoe", password: "**pass**", status: false}
	fmt.Println("Guess what, we know you")
	fmt.Println("Sign in now ...")
	fmt.Println("Login ........")
	guess.login()
	fmt.Println("You are connected")
	fmt.Println("Leaving now ? ok, bye !!!")
	guess.logout()
}
