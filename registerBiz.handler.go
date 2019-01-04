package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BizInfo struct for Business Information from registration
type BizInfo struct {
	NinjaID   int64  `json:"ninjaID"`
	Name      string `json:"name"`
	YelpBizID string `json:"yelpBizID"`
	Phone     string `json:"phone"`
	Address   struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		State    string `json:"state"`
		ZipCode  string `json:"zipCode"`
	} `json:"address"`
	YelpURL       string `json:"yelpURL"`
	Img           string `json:"img"`
	Cuisine       string `json:"cuisine"`
	Reservation   bool   `json:"reservation"`
	MobilePayment bool   `json:"mobilePayment"`
}

// registerBiz will extract incoming POST request body for businessInfo & return the same data as response
func registerBiz(w http.ResponseWriter, r *http.Request) {
	// read r.Body and assign the value to bizStruct after Unmarshal
	bizBody, _ := ioutil.ReadAll(r.Body)
	var bizStruct BizInfo
	json.Unmarshal([]byte(string(bizBody)), &bizStruct)

	// pass bizDoc as arguments to addBizFirestore func for database storage
	bizDoc := make(map[string]interface{})
	bizDocAddress := make(map[string]interface{})
	bizDoc["ninjaID"] = bizStruct.NinjaID
	bizDoc["name"] = bizStruct.Name
	bizDoc["yelpBizID"] = bizStruct.YelpBizID
	bizDoc["phone"] = bizStruct.Phone
	bizDoc["address"] = bizDocAddress
	bizDocAddress["address1"] = bizStruct.Address.Address1
	bizDocAddress["address2"] = bizStruct.Address.Address2
	bizDocAddress["city"] = bizStruct.Address.City
	bizDocAddress["state"] = bizStruct.Address.State
	bizDocAddress["zipCode"] = bizStruct.Address.ZipCode
	bizDoc["yelpURL"] = bizStruct.YelpURL
	bizDoc["img"] = bizStruct.Img
	bizDoc["cuisine"] = bizStruct.Cuisine
	bizDoc["reservation"] = bizStruct.Reservation
	bizDoc["mobilePayment"] = bizStruct.MobilePayment
	addBizFirestore(bizDoc)

	// Return JSON as response back to the client
	bizMarshal, err := json.Marshal(bizDoc)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bizMarshal)
}
