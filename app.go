package main

import (
	"RestfulDemo/API"
	"RestfulDemo/Database"
	"RestfulDemo/Service"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	userdb, err := Database.DefaultUserData()
	if err != nil {
		return
	}
	userService := Service.DefaultUserService(userdb)
	userhandler := API.DefaultUserHandler(userService)

	m := martini.New()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(render.Renderer())

	r := martini.NewRouter()
	r.Get("/users", userhandler.GetUsers)
	r.Post("/users", userhandler.CreateUser)
	r.Get("/users/:user_id/relationships", userhandler.GetAllRelations)
	r.Put("/users/:user_id/relationships/:other_user_id", userhandler.ChangeUserRelation)

	m.Action(r.Handle)
	m.RunOnAddr(":8080")
}
