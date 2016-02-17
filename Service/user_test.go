package Service

import (
	"testing"

	"RestfulDemo/Model"
)

var userService *UserService

func Test_Init(t *testing.T) {
	userService = new(UserService)
	userService.userDatabase = new(TestUserDB)
}

func Test_GetUser(t *testing.T) {
	user, err := userService.GetUser(1)

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if user == nil {
		t.Fatal("Fetch user error")
	}
}

func Test_GetUsers(t *testing.T) {
	users, err := userService.GetUsers()

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if len(users) != 1 {
		t.Fatal("Fetch users length error")
	}
}

func Test_CreateUser(t *testing.T) {
	user, err := userService.userDatabase.InsertUser("test")

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if user == nil {
		t.Fatal("Create user fatal")
		return
	}

	if user.Name != "test" {
		t.Fatal("New user's name is wrong")
	}
}

func Test_GetUserRelations(t *testing.T) {
	rs, err := userService.GetUserRelations(1)

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if len(rs) != 1 {
		t.Fatal("Fetch relations length error")
	}
}

func Test_ChangeUserRelation(t *testing.T) {
	r, err := userService.ChangeUserRelation(1, 2, 1)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if r.Relation != 2 {
		t.Fatal("Insert user relation wrong")
		return
	}

	r, err = userService.ChangeUserRelation(1, 2, -1)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if r.Relation != -1 {
		t.Fatal("Update user relation wrong")
	}
}

type TestUserDB struct {
}

func (t *TestUserDB) GetUser(userid int64) (*Model.User, error) {
	return &Model.User{
		Id:   1,
		Name: "test",
		Type: "user",
	}, nil
}

func (t *TestUserDB) GetUsers() ([]*Model.User, error) {
	var users = make([]*Model.User, 0, 0)
	users = append(users, &Model.User{
		Id:   1,
		Name: "test",
		Type: "user",
	})
	return users, nil
}
func (t *TestUserDB) InsertUser(name string) (*Model.User, error) {
	return &Model.User{
		Id:   1,
		Name: name,
		Type: "user",
	}, nil
}
func (t *TestUserDB) GetRelation(userid int64, targetid int64) (*Model.Relation, error) {

	var r *Model.Relation

	if userid == 1 && targetid == 2 {
		r = &Model.Relation{
			Id:       1,
			UserId:   1,
			TargetId: 2,
			Relation: 1,
		}
	} else if userid == 2 && targetid == 1 {
		r = &Model.Relation{
			Id:       2,
			UserId:   2,
			TargetId: 1,
			Relation: 1,
		}
	}

	return r, nil
}
func (t *TestUserDB) GetRelations(userid int64) ([]*Model.Relation, error) {
	var rs = make([]*Model.Relation, 0, 0)
	if userid == 1 {
		rs = append(rs, &Model.Relation{
			Id:       1,
			UserId:   1,
			TargetId: 2,
			Relation: 1,
		})
	}

	return rs, nil
}
func (t *TestUserDB) InsertRelation(userid int64, targetid int64, relation int64) (*Model.Relation, error) {
	return &Model.Relation{
		Id:       1,
		UserId:   userid,
		TargetId: targetid,
		Relation: relation,
	}, nil
}
func (t *TestUserDB) UpdateRelation(userid int64, targetid int64, relation int64) (*Model.Relation, error) {
	return &Model.Relation{
		Id:       1,
		UserId:   userid,
		TargetId: targetid,
		Relation: relation,
	}, nil
}
