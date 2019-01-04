package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BizInfo struct for Business Information from registration
type BizInfo struct {
	UUID      int64  `json:"uuid"`
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

	// JSON data can now be accessed in the form of bizStruct.Address

	// Re-formatting the JSON through Marshal to be sent back to the client
	bizMarshal, err := json.Marshal(bizStruct)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bizMarshal)
}
