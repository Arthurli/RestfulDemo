package Model

import (
	"fmt"
)

type User struct {
	Id   int64  `pg:"id"`
	Name string `pg:"name"`
	Type string `pg:"type"`
}

func (u *User) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Name, u.Type)
}

type Relation struct {
	Id       int64 `pg:"id"`
	UserId   int64 `pg:"userid"`
	TargetId int64 `pg:"target"`
	Relation int64 `pg:"relation"`
}

var AllowedRelation = map[string]int64{
	RelationLiked:    1,
	RelationDisliked: -1,
}

const (
	RelationLiked    = "liked"
	RelationDisliked = "disliked"
)

func (r *Relation) String() string {
	return fmt.Sprintf("Relation<%d %d %d %d>", r.Id, r.UserId, r.TargetId, r.Relation)
}
