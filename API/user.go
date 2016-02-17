package API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"RestfulDemo/Model"
	"RestfulDemo/Service"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	pg "gopkg.in/pg.v3"
)

type UserHandler struct {
	userService *Service.UserService
}

func DefaultUserHandler(userService *Service.UserService) *UserHandler {
	userHandler := new(UserHandler)
	userHandler.userService = userService
	return userHandler
}

//Get /users
func (u *UserHandler) GetUsers(request *http.Request, rd render.Render, params martini.Params) {

	users, err := u.userService.GetUsers()
	if err != nil {
		rd.Text(500, err.Error())
		return
	}

	rd.JSON(200, users)
}

//Post /users
func (u *UserHandler) CreateUser(request *http.Request, rd render.Render, params martini.Params) {

	// 获取 json body 参数
	decoder := json.NewDecoder(request.Body)
	p := struct {
		Name string `json:"name"`
	}{}
	err := decoder.Decode(&p)
	if err != nil || len(p.Name) == 0 {
		rd.Text(400, "The request cannot be fulfilled due to bad syntax.")
		return
	}

	user, err := u.userService.CreateUser(p.Name)
	if err != nil {
		rd.Text(500, err.Error())
		return
	}

	rd.JSON(200, user)
}

//Put /users/:user_id/relationships/:other_user_id
func (u *UserHandler) ChangeUserRelation(request *http.Request, rd render.Render, params martini.Params) {

	// 获取 url 参数
	userid, err1 := strconv.ParseInt(params["user_id"], 10, 64)
	otherid, err2 := strconv.ParseInt(params["other_user_id"], 10, 64)
	if err1 != nil || err2 != nil {
		rd.Text(400, "The request cannot be fulfilled due to bad syntax.")
		return
	}

	// 获取 json body 参数
	decoder := json.NewDecoder(request.Body)
	p := struct {
		State string `json:"state"`
	}{}
	err := decoder.Decode(&p)

	//验证 user 是否存在
	_, err = u.userService.GetUser(userid)
	if err != nil {
		if err == pg.ErrNoRows {
			rd.Text(400, fmt.Sprintln("user", userid, "not exist"))
		} else {
			rd.Text(500, err.Error())
		}
		return
	}
	_, err = u.userService.GetUser(otherid)
	if err != nil {
		if err == pg.ErrNoRows {
			rd.Text(400, fmt.Sprintln("user", otherid, "not exist"))
		} else {
			rd.Text(500, err.Error())
		}
		return
	}

	if _, ok := Model.AllowedRelation[p.State]; err != nil || !ok {
		rd.Text(400, "The request cannot be fulfilled due to bad syntax.")
		return
	}

	r, err := u.userService.ChangeUserRelation(userid, otherid, Model.AllowedRelation[p.State])
	if err != nil {
		rd.Text(500, err.Error())
		return
	}

	rd.JSON(200, fmtRelationShip(r))
}

//Get /users/:user_id/relationships
func (u *UserHandler) GetAllRelations(request *http.Request, rd render.Render, params martini.Params) {
	// 获取 url 参数
	userid, err := strconv.ParseInt(params["user_id"], 10, 64)
	if err != nil {
		rd.Text(400, "The request cannot be fulfilled due to bad syntax.")
		return
	}

	//验证 user 是否存在
	_, err = u.userService.GetUser(userid)
	if err != nil {
		if err == pg.ErrNoRows {
			rd.Text(400, fmt.Sprintln("user", userid, "not exist"))
		} else {
			rd.Text(500, err.Error())
		}
		return
	}

	rs, err := u.userService.GetUserRelations(userid)
	if err != nil {
		rd.Text(500, err.Error())
		return
	}

	relations := make([]interface{}, 0, 0)
	for _, r := range rs {
		relations = append(relations, fmtRelationShip(r))
	}

	rd.JSON(200, relations)
}

// 格式化返回值
func fmtRelationShip(r *Model.Relation) map[string]interface{} {
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
