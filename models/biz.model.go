package models

import "gopkg.in/mgo.v2/bson"

// Biz struct represents a basic information of restaurant business
type Biz struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	YelpBizID string        `json:"yelpBizID"`
	YelpURL   string        `json:"yelpURL"`
	Name      string        `json:"name"`
	Phone     string        `json:"phone"`
	Address   struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		State    string `json:"state"`
		ZipCode  string `json:"zipCode"`
	} `json:"address"`
	Img           string   `json:"img"`
	Cuisine       string   `json:"cuisine"`
	Reservation   bool     `json:"reservation"`
	MobilePayment bool     `json:"mobilePayment"`
	Orders        []string `json:"orders"`
	Active        bool     `json:"active"`
}

// YelpJSN struct for incoming JSN from API
type YelpJSN struct {
	ID           string `json:"id"`
	Alias        string `json:"alias"`
	Name         string `json:"name"`
	ImageURL     string `json:"image_url"`
	IsClaimed    bool   `json:"is_claimed"`
	IsClosed     bool   `json:"is_closed"`
	URL          string `json:"url"`
	Phone        string `json:"phone"`
	DisplayPhone string `json:"display_phone"`
	ReviewCount  int    `json:"review_count"`
	Categories   []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	Rating   float64 `json:"rating"`
	Location struct {
		Address1       string      `json:"address1"`
		Address2       string      `json:"address2"`
		Address3       interface{} `json:"address3"`
		City           string      `json:"city"`
		ZipCode        string      `json:"zip_code"`
		Country        string      `json:"country"`
		State          string      `json:"state"`
		DisplayAddress []string    `json:"display_address"`
		CrossStreets   string      `json:"cross_streets"`
	} `json:"location"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Photos []string `json:"photos"`
	Price  string   `json:"price"`
	Hours  []struct {
		Open []struct {
			IsOvernight bool   `json:"is_overnight"`
			Start       string `json:"start"`
			End         string `json:"end"`
			Day         int    `json:"day"`
		} `json:"open"`
		HoursType string `json:"hours_type"`
		IsOpenNow bool   `json:"is_open_now"`
	} `json:"hours"`
	Transactions []interface{} `json:"transactions"`
}

// Address struct for address details
type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	ZipCode  string `json:"zipCode"`
}

// BizYelpJSN struct for compsing data of return Yelp JSN
type BizYelpJSN struct {
	YelpBizID     string  `json:"yelpBizID"`
	YelpURL       string  `json:"yelpURL"`
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	Address       Address `json:"address"`
	Img           string  `json:"img"`
	Cuisine       string  `json:"cuisine"`
	Reservation   bool    `json:"reservation"`
	MobilePayment bool    `json:"mobilePayment"`
	Active        bool    `json:"active"`
}
