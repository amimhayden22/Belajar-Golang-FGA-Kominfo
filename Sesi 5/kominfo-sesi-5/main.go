package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Id			string
	Name		string
	Address		string
	Occupation 	string
	Reason 		string
}

var friends []Friend = []Friend{
	Friend{
		Id: "C1",
		Name: "Gus",
		Address: "Jalan Bangka",
		Occupation: "Student",
		Reason: "Ingin belajar hal baru",
	},
	{
		Id: "C2",
		Name: "Khamim",
		Address: "Jalan Grabag",
		Occupation: "Student",
		Reason: "Saya suka golang",
	},
}

func searchFriendById(id string)  {
	for _, v := range friends {
		if v.Id == id {
			fmt.Printf("Name: %s\n", v.Name)
			fmt.Printf("Address: %s\n", v.Address)
			fmt.Printf("Occupation: %s\n", v.Occupation)
			fmt.Printf("Reason: %s\n", v.Reason)

			return
		}
	}

	fmt.Printf("Friend with id %s can not be found\n", id)
}

func main() {
	var arguments []string = os.Args

	if len(arguments) <= 2 {
		fmt.Println("Usage: go run main.go [id]")
		return
	}

	searchFriendById(arguments[1])

	// fmt.Printf("args => %#v\n", arguments)
}