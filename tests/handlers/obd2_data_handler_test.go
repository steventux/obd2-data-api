package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/steventux/obd2-data-api/handlers"
	"github.com/tborisova/clean_like_gopher"
	"net/http"
	"net/http/httptest"
)

var (
	m clean_like_gopher.Generic
)

var _ = Describe("Obd2DataHandler authenticated", func() {
	BeforeEach(func() {
		m, _ = clean_like_gopher.NewCleaningGopher("mongo",
			map[string]string{"host": "localhost", "dbName": "test", "port": "27017"})

		req, _ := http.NewRequest("GET", "/", nil)
		query := req.URL.Query()
		query.Add("k42", "14.4")
		query.Add("kff1202", "120")
		query.Add("kff1225", "198")

		req.URL.RawQuery = query.Encode()

		rr = httptest.NewRecorder()
		handler = http.HandlerFunc(handlers.CreateObd2Data)

		handler.ServeHTTP(rr, req)
	})

	AfterEach(func() {
		m.Clean(nil)
		m.Close()
	})

	Describe("response status code", func() {
		It("should be OK", func() {
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("response body", func() {
		It("should respond with 'OK'", func() {
			Expect(rr.Body.String()).To(Equal("OK"))
		})
	})
})

var _ = Describe("Obd2DataHandler unauthenticated", func() {
	BeforeEach(func() {
		req, _ := http.NewRequest("GET", "/", nil)
		query := req.URL.Query()
		// Add an id param to authenticate with.
		query.Add("id", "123abc")
		query.Add("k42", "14.4")
		query.Add("kff1202", "120")
		query.Add("kff1225", "198")

		req.URL.RawQuery = query.Encode()

		rr = httptest.NewRecorder()
		handler = http.HandlerFunc(handlers.CreateObd2Data)

		handler.ServeHTTP(rr, req)
	})

	Describe("response status code", func() {
		It("should be 403", func() {
			Expect(rr.Code).To(Equal(403))
		})
	})

	Describe("response body", func() {
		It("should respond with 'Forbidden'", func() {
			Expect(rr.Body.String()).To(Equal("Forbidden"))
		})
	})
})
