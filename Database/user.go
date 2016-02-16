package Database

import (
	"RestfulDemo/Model"
	pg "gopkg.in/pg.v3"
)

type User interface {
	GetUsers() ([]*Model.User, error)
	InsertUser(string) error
	GetRelation(int64, int64) (*Model.Relation, error)
	GetRelations(int64) ([]*Model.Relation, error)
	InsertRelation(int64, int64, int64) error
	UpdateRelation(int64, int64, int64) error
}

func DefaultUserData() (*UserData, error) {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	user := new(UserData)
	user.db = db
	err := user.Init()

	return user, err
}

type UserData struct {
	db *pg.DB
}

func (u *UserData) Init() error {
	// 由于 Demo 数据较少 不去添加索引
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (id serial, name text, type text)`,
		`CREATE TABLE IF NOT EXISTS relations (id serial, userid bigint, target bigint, relation integer)`,
	}
	for _, q := range queries {
		_, err := u.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UserData) GetUsers() ([]*Model.User, error) {
	var users []*Model.User
	_, err := u.db.Query(&users, `SELECT * FROM users`)
	return users, err
}

func (u *UserData) InsertUser(name string) (*Model.User, error) {
	var user Model.User
	_, err := u.db.QueryOne(&user, `
        INSERT INTO users (name, type) VALUES (?, ?)
        RETURNING id, name, type
    `, name, "user")
	return &user, err
}

func (u *UserData) GetRelation(userid, target int64) (*Model.Relation, error) {
	var relation Model.Relation
	_, err := u.db.QueryOne(&relation, `SELECT * FROM relations WHERE userid = ? AND target = ?`, userid, target)
	return &relation, err
}

func (u *UserData) GetRelations(userid int64) ([]*Model.Relation, error) {
	var relations []*Model.Relation
	_, err := u.db.Query(&relations, `SELECT * FROM relations WHERE userid = ?`, userid)
	return relations, err
}

func (u *UserData) InsertRelation(userid, targetid, relarion int64) (*Model.Relation, error) {
	var relationObject Model.Relation
	_, err := u.db.QueryOne(&relationObject, `
        INSERT INTO relations (userid, target, relation) VALUES (?, ?, ?)
        RETURNING id, userid, target, relation
    `, userid, targetid, relarion)
	return &relationObject, err
}

func (u *UserData) UpdateRelation(userid, targetid, relarion int64) error {
	_, err := u.db.Exec(`
        UPDATE relations SET relation = ? WHERE userid = ? AND target = ?
    `, relarion, userid, targetid)
	return err
}
