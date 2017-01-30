package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/steventux/obd2-data-api/models"
	"github.com/steventux/obd2-data-api/services"
	"github.com/tborisova/clean_like_gopher"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	cleaner clean_like_gopher.Generic
)

var _ = Describe("Obd2Data service functions", func() {
	BeforeEach(func() {
		cleaner, _ = clean_like_gopher.NewCleaningGopher("mongo",
			map[string]string{"host": "localhost", "dbName": "test", "port": "27017"})
	})

	AfterEach(func() {
		cleaner.Clean(nil)
		cleaner.Close()
	})

	Describe("SaveObd2Data", func() {
		It("should save Obd2Data", func() {
			services.SaveObd2Data(&models.Obd2Data{
				Torque:    "198",
				Voltage:   "14.4",
				EngineRPM: "2500",
			})

			var result models.Obd2Data

			err := services.Obd2DataCollection().Find(bson.M{"voltage": "14.4"}).One(&result)

			Expect(err).To(BeNil())
			Expect(result.Torque).To(Equal("198"))
			Expect(result.Voltage).To(Equal("14.4"))
			Expect(result.EngineRPM).To(Equal("2500"))
			Expect(result.CreatedAt).To(BeTemporally("~", time.Now(), time.Second))
		})
	})

	Describe("GetLatestObd2DataSession", func() {
		BeforeEach(func() {
			services.SaveObd2Data(&models.Obd2Data{
				Session: "12712bbl21jb312o",
			})

			services.SaveObd2Data(&models.Obd2Data{
				Session: "zyxababaw8qwhdqw",
			})

			services.SaveObd2Data(&models.Obd2Data{
				Session: "abcdefg12345678",
			})
		})

		It("should return the latest session identifier", func() {
			sessionId := services.GetLatestObd2DataSession()

			Expect(sessionId).To(Equal("abcdefg12345678"))
		})
	})

	Describe("GetObd2DataForSession", func() {
		BeforeEach(func() {
			services.SaveObd2Data(&models.Obd2Data{
				Session: "12712bbl21jb312o",
			})

			services.SaveObd2Data(&models.Obd2Data{
				Session: "abcdefg12345678",
			})

			services.SaveObd2Data(&models.Obd2Data{
				Session: "abcdefg12345678",
			})
		})

		It("retrieves data for a given session ID", func() {
			results := services.GetObd2DataForSession("abcdefg12345678")

			Expect(results).To(HaveLen(2))
			Expect(results[0].Session).To(Equal("abcdefg12345678"))
			Expect(results[1].Session).To(Equal("abcdefg12345678"))
			Expect(results[0].CreatedAt).To(BeTemporally("<", results[1].CreatedAt))
		})
	})
})
