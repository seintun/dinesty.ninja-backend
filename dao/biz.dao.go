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

// ValidateYelp biz with Yelp
func (b *BizDAO) ValidateYelp(yURL string, bearer string) (BizYelpJSN, error) {
	request, _ := http.NewRequest("GET", yURL, nil)
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
	query := bson.M{"active": true}
	var bizs []Biz
	err := db.C(BCOLLECTION).Find(query).All(&bizs)
	return bizs, err
}

// FindBizByID return specified Biz
func (b *BizDAO) FindBizByID(id string) (Biz, error) {
	query := bson.ObjectIdHex(id)
	var biz Biz
	err := db.C(BCOLLECTION).FindId(query).One(&biz)
	return biz, err
}

// RegisterBiz a biz into database
func (b *BizDAO) RegisterBiz(biz Biz) error {
	err := db.C(BCOLLECTION).Insert(&biz)
	return err
}

// UpdateBizByID an existing biz
func (b *BizDAO) UpdateBizByID(id string, biz Biz) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(BCOLLECTION).Update(query, &biz)
	return err
}

// DeactivateBizByID an existing biz
func (b *BizDAO) DeactivateBizByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	deactived := bson.M{"$set": bson.M{"active": false}}
	err := db.C(BCOLLECTION).Update(query, deactived)
	return err
}

// DeleteBizByID an existing biz
func (b *BizDAO) DeleteBizByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(BCOLLECTION).Remove(query)
	return err
}
