// handlers_test.go
package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/steventux/obd2-data-api/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("HomeHandler", func() {
	BeforeEach(func() {
		req, _ := http.NewRequest("GET", "/", nil)

		rr = httptest.NewRecorder()
		handler = http.HandlerFunc(handlers.Home)

		handler.ServeHTTP(rr, req)
	})

	Describe("response status code", func() {
		It("should be OK", func() {
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("response body", func() {
		It("should reflect the status", func() {
			Expect(rr.Body.String()).To(Equal(`{"status":"OK"}`))
		})
	})
})
