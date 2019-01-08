package dao

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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
	COLLECTION = "business"
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

// FetchYelpJSN biz with Yelp
func (b *BizDAO) FetchYelpJSN(yelpURL string, bearer string) (BizYelpJSN, error) {
	request, _ := http.NewRequest("GET", yelpURL, nil)
	request.Header.Add("Authorization", bearer)
	client := &http.Client{}
	yelpR, _ := client.Do(request)
	data, _ := ioutil.ReadAll(yelpR.Body)
	defer yelpR.Body.Close()

	var yJSN YelpJSN
	err := json.Unmarshal([]byte(data), &yJSN)
	bJSN := BizYelpJSN{
		YelpBizID: yJSN.Alias,
		YelpURL:   yJSN.URL,
		Name:      yJSN.Name,
		Phone:     yJSN.Phone,
		Address: Address{
			Address1: yJSN.Location.Address1,
			Address2: yJSN.Location.Address2,
			City:     yJSN.Location.City,
			State:    yJSN.Location.State,
			ZipCode:  yJSN.Location.ZipCode,
		},
		Img:           yJSN.ImageURL,
		Cuisine:       yJSN.Categories[0].Title,
		Reservation:   false,
		MobilePayment: false,
		Active:        true,
	}
	return bJSN, err
}

// FetchBiz return list of bizs
func (b *BizDAO) FetchBiz() ([]Biz, error) {
	var bizs []Biz
	err := db.C(COLLECTION).Find(bson.M{}).All(&bizs)
	return bizs, err
}

// FindBizByID return specified Biz
func (b *BizDAO) FindBizByID(id string) (Biz, error) {
	var biz Biz
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&biz)
	return biz, err
}

// Insert a biz into database
func (b *BizDAO) Insert(biz Biz) error {
	err := db.C(COLLECTION).Insert(&biz)
	return err
}

// UpdateBizByID an existing biz
func (b *BizDAO) UpdateBizByID(id string, biz Biz) error {
	err := db.C(COLLECTION).Update(bson.M{"_id": bson.ObjectIdHex(id)}, &biz)
	return err
}

// // DeactivateBizByID an existing biz
// func (b *BizDAO) DeactivateBizByID(id string) error {
// 	err := db.C(COLLECTION).Update(bson.ObjectIdHex(id), {$set:{"active":false}})
// 	return err
// }

// DeleteBizByID an existing biz
func (b *BizDAO) DeleteBizByID(id string) error {
	err := db.C(COLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
