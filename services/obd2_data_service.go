package services

import (
	"github.com/steventux/obd2-data-api/models"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func SaveObd2Data(obd2Data *models.Obd2Data) (bool, error) {
	err := Obd2DataCollection().Insert(obd2Data)

	if err != nil {
		log.Fatal(err)
		return false, err
	} else {
		return true, nil
	}
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
