package API

import (
	"RestfulDemo/Model"
	"RestfulDemo/Service"
)

type UserHandler struct {
	userService *Service.UserService
}

func DefaultUserHandler(userService *Service.UserService) UserHandler {
	userHandler := new(UserHandler)
	userHandler.userService = userService
	return userHandler
}

func (u *UserHandler) GetUsers() interface{} {

	users, err := u.userService.GetUsers()
	if err != nil {
		return err
	}

	return users
}

func (u *UserHandler) CreateUser() interface{} {

	name := "123"
	user, err := u.userService.CreateUser(name)
	if err != nil {
		return err
	}

	return user
}

func (u *UserHandler) ChangeUserRelation() interface{} {

	r, err := u.userService.ChangeUserRelation(0, 0, 1)
	if err != nil {
		return err
	}

	return fmtRelationShip(r)
}

func (u *UserHandler) GetAllRelations() interface{} {

	rs, err := u.userService.GetUserRelations(0)
	if err != nil {
		return err
	}

	relations := make([]interface{}, 0, 0)
	for _, r := range rs {
		relations = append(relations, fmtRelationShip(r))
	}

	return relations
}

func fmtRelationShip(r *Model.Relation) interface{} {
	relation := make(map[string]interface{})
	relation["id"] = r.TargetId
	relation["type"] = "relationship"
	switch r.Relation {
	case -1:
		relation["state"] = "disliked"
	case 1:
		relation["state"] = "liked"
	default:
		relation["state"] = "matched"
	}

	return relation
}
