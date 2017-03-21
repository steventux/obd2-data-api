package services

import (
	"github.com/steventux/obd2-data-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

func SaveObd2Data(db *mgo.Database, obd2Data *models.Obd2Data) (bool, error) {
	obd2Data.CreatedAt = time.Now()
	err := obd2DataCollection(db).Insert(obd2Data)

	if err != nil {
		log.Fatal(err)
		return false, err
	} else {
		return true, nil
	}
}

func GetLatestObd2DataSession(db *mgo.Database) string {
	var result *models.Obd2Data

	err := obd2DataCollection(db).Find(nil).Sort("-createdat").One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Session
}

func GetObd2DataForSession(db *mgo.Database, session string) []models.Obd2Data {
	var results []models.Obd2Data

	err := obd2DataCollection(db).Find(bson.M{"session": session}).Sort("createdat").All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func obd2DataCollection(db *mgo.Database) *mgo.Collection {
	return db.C(os.Getenv("MONGODB_COLLECTION"))
}
