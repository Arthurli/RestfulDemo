package Database

import (
	"testing"

	"RestfulDemo/Model"
	pg "gopkg.in/pg.v3"
)

var testUserDB *UserData
var user1, user2 *Model.User
var err error

func Test_Init(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User: "test",
	})
	user := new(UserData)
	user.db = db
	err = user.Init()
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	testUserDB = user
}

func Test_CreateUser(t *testing.T) {
	if testUserDB == nil {
		return
	}

	user1, err = testUserDB.InsertUser("user1")
	if err != nil {
		t.Fatal(err.Error())
	}

	user2, err = testUserDB.InsertUser("user2")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_GetUser(t *testing.T) {
	if testUserDB == nil {
		return
	}
	_, err = testUserDB.GetUser(user1.Id)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_GetUsers(t *testing.T) {
	if testUserDB == nil {
		return
	}

	_, err = testUserDB.GetUsers()
	if err != nil {
		t.Fatal(err.Error())
	}
}
func Test_InsertRelation(t *testing.T) {
	if testUserDB == nil ||
		user1 == nil ||
		user2 == nil {
		return
	}

	r, err := testUserDB.InsertRelation(user1.Id, user2.Id, 1)
	if err != nil {
		t.Fatal(err.Error())
	}

	if r.Relation != 1 {
		t.Fatal("Insert error")
	}
}
func Test_UpdateRelation(t *testing.T) {
	if testUserDB == nil ||
		user1 == nil ||
		user2 == nil {
		return
	}

	r, err := testUserDB.UpdateRelation(user1.Id, user2.Id, -1)
	if err != nil {
		t.Fatal(err.Error())
	}

	if r.Relation != -1 {
		t.Fatal("Update error")
	}
}
func Test_GetRelation(t *testing.T) {
	if testUserDB == nil ||
		user1 == nil ||
		user2 == nil {
		return
	}

	_, err = testUserDB.GetRelation(user1.Id, user2.Id)
	if err != nil {
		t.Fatal(err.Error())
	}
}
func Test_GetRelations(t *testing.T) {
	if testUserDB == nil ||
		user1 == nil ||
		user2 == nil {
		return
	}

	_, err = testUserDB.GetRelations(user1.Id)
	if err != nil {
		t.Fatal(err.Error())
	}
}
