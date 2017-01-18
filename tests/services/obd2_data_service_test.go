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
	m, _ := clean_like_gopher.NewCleaningGopher("mongo",
		map[string]string{"host": "127.0.0.1", "dbName": "test", "port": "27017"})

	services.SaveObd2Data(&models.Obd2Data{
		Torque:    "198",
		Voltage:   "14.4",
		EngineRPM: "2500",
	})

	var result struct {
		Text string `bson:"text"`
	}

	err := services.Obd2DataCollection().Find(bson.M{"voltage": "14.4"}).One(&result)

	if err != nil {
		t.Errorf("error while retrieving data: %v", err)
	}

	if &result == nil {
		t.Errorf("service failed to save Torque data: got %v", result)
	}

	m.Clean(nil)
	m.Close()
}
