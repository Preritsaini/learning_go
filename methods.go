package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	prerit := User{"Prerit", "sainiprerit6@gmail.com", true, 20}
	fmt.Println(prerit)
	fmt.Printf("Prerit details are %+v\n", prerit)
	fmt.Printf("Name is %v and Email is %v\n", prerit.Name, prerit.Email)

	prerit.GetStatus()
	prerit.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active:", u.Status)
}
func (u User) NewMail() {
	u.Email = "gogo.dev"
}
