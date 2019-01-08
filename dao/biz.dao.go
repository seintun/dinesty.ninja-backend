package dao

import (
	"log"

	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type BizDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	BCOLLECTION = "business"
)

// Connect Establish a connection to database
func (b *BizDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(b.Database)
}

// Queries

// FetchBiz list of bizs
func (b *BizDAO) FetchBiz() ([]Biz, error) {
	var bizs []Biz
	err := db.C(BCOLLECTION).Find(bson.M{}).All(&bizs)
	return bizs, err
}

// FindBizByID return specified Biz
func (b *BizDAO) FindBizByID(id string) (Biz, error) {
	var biz Biz
	err := db.C(BCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&biz)
	return biz, err
}

// RegisterBiz a biz into database
func (b *BizDAO) RegisterBiz(biz Biz) error {
	err := db.C(BCOLLECTION).Insert(&biz)
	return err
}

// UpdateBizByID an existing biz
func (b *BizDAO) UpdateBizByID(id string, biz Biz) error {
	err := db.C(BCOLLECTION).Update(bson.M{"_id": bson.ObjectIdHex(id)}, &biz)
	return err
}

// // DeactivateBizByID an existing biz
// func (b *BizDAO) DeactivateBizByID(id string) error {
// 	err := db.C(BCOLLECTION).Update(bson.ObjectIdHex(id), {$set:{"active":false}})
// 	return err
// }

// DeleteBizByID an existing biz
func (b *BizDAO) DeleteBizByID(id string) error {
	err := db.C(BCOLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
