package helpers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/steventux/obd2-data-api/helpers"
	"github.com/steventux/obd2-data-api/models"
	"net/http"
)

var (
	obd2Data *models.Obd2Data
)

var _ = Describe("helpers.BuildObd2Data", func() {
	BeforeEach(func() {
		req, _ := http.NewRequest("GET", "/", nil)

		query := req.URL.Query()
		query.Set("k42", "14.4")
		query.Add("kc", "2550")
		query.Add("kff1225", "199")
		query.Add("kff5201", "36")

		req.URL.RawQuery = query.Encode()

		obd2Data = helpers.BuildObd2Data(req)
	})

	It("should convert request params to model attributes", func() {
		Expect(obd2Data.Voltage).To(Equal("14.4"))
		Expect(obd2Data.Torque).To(Equal("199"))
		Expect(obd2Data.EngineRPM).To(Equal("2550"))
		Expect(obd2Data.MPGLongTermAverage).To(Equal("36"))
	})
})
