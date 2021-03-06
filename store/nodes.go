package store

import (
	"launchpad.net/mgo/bson"
)

// Node array type
type Nodes []Node

func (n *Nodes) GetAll(q bson.M) (err error) {
	db, err := DBConnect()
	if err != nil {
		return
	}
	defer db.Close()
	err = db.GetAll(q, n)
	return
}

func (n *Nodes) GetAllLimitOffset(q bson.M, limit int, offset int) (err error) {
	db, err := DBConnect()
	if err != nil {
		return
	}
	defer db.Close()
	err = db.GetAllLimitOffset(q, n, limit, offset)
	return
}
