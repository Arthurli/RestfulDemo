package main

import (
	"RestfulDemo/Database"
	"fmt"
)

func main() {
	var userData *Database.UserData
	var err error
	userData, err = Database.DefaultUserData()
	if err != nil {
		fmt.Println(err)
	}

	user1, err := userData.InsertUser("123")
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return
	}
	fmt.Println("user1", user1)

	user2, err := userData.InsertUser("456")
	if err != nil {
		fmt.Println("1.1")
		fmt.Println(err)
		return
	}
	fmt.Println("user2", user2)

	users, err := userData.GetUsers()
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return
	}
	fmt.Println(users)

	a, err := userData.InsertRelation(user1.Id, user2.Id, 1)
	if err != nil {
		fmt.Println("3")
		fmt.Println(err)
		return
	}
	fmt.Println(a)

	err = userData.UpdateRelation(user1.Id, user2.Id, 2)
	if err != nil {
		fmt.Println("4")
		fmt.Println(err)
		return
	}

	b, err := userData.GetRelation(user1.Id, user2.Id)
	if err != nil {
		fmt.Println("5")
		fmt.Println(err)
		return
	}
	fmt.Println(b)

	c, err := userData.GetRelations(user1.Id)
	if err != nil {
		fmt.Println("6")
		fmt.Println(err)
		return
	}
	fmt.Println(c)

}
