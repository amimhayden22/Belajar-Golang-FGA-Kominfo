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
	{
		Id: "C3",
		Name: "Amim",
		Address: "Jalan Secang",
		Occupation: "Student",
		Reason: "Saya penasaran dengan golang",
	},
	{
		Id: "C4",
		Name: "Hayden",
		Address: "Jalan Suka Maju",
		Occupation: "Student",
		Reason: "Saya ingin bisa golang",
	},
	{
		Id: "C5",
		Name: "Muhammad",
		Address: "Jalan Jogja Magelang KM 22",
		Occupation: "Student",
		Reason: "Saya belum bisa ngoding",
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

	// Untuk memeriksa apakah user input id atau tidak
	if len(arguments) <= 1 {
		fmt.Println("Usage: go run main.go [id]")
		return
	}

	searchFriendById(arguments[1])
}