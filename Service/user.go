package Service

import (
	"log"
	"runtime"

	"RestfulDemo/Database"
	"RestfulDemo/Model"
	pg "gopkg.in/pg.v3"
)

type UserService struct {
	userDatabase Database.User
}

func DefaultUserService(userDatabase Database.User) *UserService {
	userService := new(UserService)
	userService.userDatabase = userDatabase

	return userService
}

func (u *UserService) GetUsers() ([]*Model.User, error) {
	users, err := u.userDatabase.GetUsers()
	if err != nil {
		Nlog(0, err)
	}

	return users, err
}

func (u *UserService) CreateUser(name string) (*Model.User, error) {
	user, err := u.userDatabase.InsertUser(name)
	if err != nil {
		Nlog(0, err)
	}
	return user, err
}

func (u *UserService) GetUserRelations(userid int64) ([]*Model.Relation, error) {
	rs, err := u.userDatabase.GetRelations(userid)
	if err != nil {
		Nlog(0, err)
		return nil, err
	}
	relations := make([]*Model.Relation, 0)
	for _, r := range rs {

		relation, err := u.getMutualRelationship(r)
		if err != nil {
			Nlog(0, err)
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}

func (u *UserService) ChangeUserRelation(userid, targetid, relation int64) (*Model.Relation, error) {
	_, err := u.userDatabase.GetRelation(userid, targetid)
	if err != nil && err != pg.ErrNoRows {
		Nlog(0, err)
		return nil, err
	}

	var r *Model.Relation

	// 已经建立关系的更新关系, 没有创建关系的 创建新关系
	if err == pg.ErrNoRows {
		r, err = u.userDatabase.InsertRelation(userid, targetid, relation)
	} else {
		r, err = u.userDatabase.UpdateRelation(userid, targetid, relation)
	}

	if err != nil {
		Nlog(0, err)
		return nil, err
	}

	r, err = u.getMutualRelationship(r)
	if err != nil {
		Nlog(0, err)
	}

	return r, err
}

// 由单一的关系获取双方相互的关系 如果均为 like, 则显示为 match
func (u *UserService) getMutualRelationship(r *Model.Relation) (*Model.Relation, error) {
	relation, err := u.userDatabase.GetRelation(r.TargetId, r.UserId)
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}

	// 如果对方没有与你建立关系, 则直接返回
	if err == pg.ErrNoRows {
		return r, nil
	}

	if r.Relation == 1 && relation.Relation == 1 {
		r.Relation = 2
	}

	return r, nil
}

func Nlog(callDepth int, err error) {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	} else {
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
	}

	log.Println(file, line, err.Error())
}
