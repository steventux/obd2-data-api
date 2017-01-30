package services

import (
	"github.com/steventux/obd2-data-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

func SaveObd2Data(obd2Data *models.Obd2Data) (bool, error) {
	defer Session().Close()

	obd2Data.CreatedAt = time.Now()
	err := Obd2DataCollection().Insert(obd2Data)

	if err != nil {
		log.Fatal(err)
		return false, err
	} else {
		return true, nil
	}
}

func GetLatestObd2DataSession() string {
	defer Session().Close()
	var result *models.Obd2Data

	err := Obd2DataCollection().Find(nil).Sort("-createdat").One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Session
}

func GetObd2DataForSession(session string) []models.Obd2Data {
	defer Session().Close()
	var results []models.Obd2Data

	err := Obd2DataCollection().Find(bson.M{"session": session}).Sort("createdat").All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func Obd2DataCollection() *mgo.Collection {
	return Session().DB(os.Getenv("MONGODB_DB")).C(os.Getenv("MONGODB_COLLECTION"))
}

func Session() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("MONGODB_HOSTS"))
	if err != nil {
		log.Fatal(err)
	}

	return session
}
