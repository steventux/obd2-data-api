// services_test.go
package services_test

import (
	"github.com/steventux/obd2-data-api/models"
	"github.com/steventux/obd2-data-api/services"
	"github.com/tborisova/clean_like_gopher"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestSaveObd2Data(t *testing.T) {
	options := map[string]string{"host": "localhost", "dbName": "test", "port": "27017"}
	m, _ := clean_like_gopher.NewCleaningGopher("mongo", options)

	services.SaveObd2Data(&models.Obd2Data{
		Torque:    "198",
		Voltage:   "14.4",
		EngineRPM: "2500",
	})

	var result struct {
		Text string `bson:"text"`
	}

	err := services.Obd2DataCollection().Find(bson.M{"voltage": "14.4"}).One(&result)

	m.Clean(nil)
	m.Close()

	if err != nil {
		t.Errorf("error while retrieving data: %v", err)
	}

	if &result == nil {
		t.Errorf("service failed to save Torque data: got %v", result)
	}
}
