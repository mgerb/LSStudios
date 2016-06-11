/*

package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goskeleton/mywebsite/db"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	collection = "temperature_data"
)

type Data struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	MaxTemp   float64       `json:"maxtemp" bson:"maxtemp"`
	MinTemp   float64       `json:"mintemp" bson:"mintemp"`
	Location  string        `json:"location" bson:"location"`
	Month     int           `json:"month" bson:"month"`
	MonthName string        `json:"monthname" bson:"monthname"`
	Day       int           `json:"day" bson:"day"`
	Year      int           `json:"year" bson:"year"`
}

//convert struct to json string
func (s *Data) toJson() string {

	b, err := json.MarshalIndent(s, "", "    ")

	if err != nil {
		fmt.Println(err.Error)
	}

	return string(b)

}

//store info in database
func (s *Data) StoreData() error {

	if db.Mongo.Connected() {

		log.Println("Inserting data into " + collection)

		session := db.Mongo.Session.Copy()
		defer session.Close()

		c := session.DB(db.Mongo.Info.Database).C(collection)

		// Insert Datas
		err := c.Insert(s)

		if err != nil {
			return err
		}
	}

	return nil
}

//function to update the daily temperature max or min for a location
func (s *Data) UpdateData() error {

	if db.Mongo.Connected() {

		log.Println("Updating data")

		session := db.Mongo.Session.Copy()
		defer session.Close()

		c := session.DB(db.Mongo.Info.Database).C(collection)

		colQuerier := bson.M{"location": s.Location, "month": s.Month, "monthname": s.MonthName, "day": s.Day, "year": s.Year}
		change := bson.M{"$set": bson.M{"maxtemp": s.MaxTemp, "mintemp": s.MinTemp}}

		err := c.Update(colQuerier, change)

		if err != nil {
			return err
		}
	}

	return nil

}

*/