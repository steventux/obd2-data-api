package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/steventux/obd2-data-api/models"
	"github.com/steventux/obd2-data-api/services"
	"github.com/tborisova/clean_like_gopher"
	"gopkg.in/mgo.v2/bson"
)

var (
	m clean_like_gopher.Generic
)

var _ = Describe("Obd2Data service functions", func() {
	BeforeEach(func() {
		m, _ = clean_like_gopher.NewCleaningGopher("mongo",
			map[string]string{"host": "localhost", "dbName": "test", "port": "27017"})
	})

	AfterEach(func() {
		m.Clean(nil)
		m.Close()
	})

	Describe("SaveObd2Data", func() {
		It("should save Obd2Data", func() {
			services.SaveObd2Data(&models.Obd2Data{
				Torque:    "198",
				Voltage:   "14.4",
				EngineRPM: "2500",
			})

			var result struct {
				Torque    string `bson:"torque"`
				Voltage   string `bson:"voltage"`
				EngineRPM string `bson:"enginerpm"`
			}

			err := services.Obd2DataCollection().Find(bson.M{"voltage": "14.4"}).One(&result)

			Expect(err).To(BeNil())
			Expect(result.Torque).To(Equal("198"))
			Expect(result.Voltage).To(Equal("14.4"))
			Expect(result.EngineRPM).To(Equal("2500"))
		})
	})
})
