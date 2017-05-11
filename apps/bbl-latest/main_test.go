package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("bbl-latest", func() {
	var (
		port string
	)

	BeforeEach(func() {
		var err error
		port, err = openPort()
		Expect(err).NotTo(HaveOccurred())
		os.Setenv("PORT", port)

		bblLatest()
		waitForServerToStart(port)
	})

	Context("when someone hits /", func() {
		It("returns status not found", func() {
			response, err := http.Get(fmt.Sprintf("http://localhost:%s/", port))
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Context("when someone hits /latest", func() {
		DescribeTable("with various query parameters it redirects to latest github release", func(os string, osType string) {
			request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%s/latest?os=%s", port, os), nil)
			Expect(err).NotTo(HaveOccurred())

			response, err := http.DefaultTransport.RoundTrip(request)
			Expect(err).NotTo(HaveOccurred())

			Expect(response.StatusCode).To(Equal(http.StatusFound))

			bblVersion := getLatestBBLVersion()
			Expect(response.Header["Location"][0]).To(Equal(fmt.Sprintf("https://github.com/cloudfoundry/bosh-bootloader/releases/download/%[1]s/bbl-%[1]s_%s", bblVersion, osType)))
		},
			Entry("os=linux", "linux", "linux_x86-64"),
			Entry("os=osx", "osx", "osx"),
		)

		Context("when os param is invalid", func() {
			It("returns a 400", func() {
				request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%s/latest?os=invalid", port), nil)
				Expect(err).NotTo(HaveOccurred())

				response, err := http.DefaultTransport.RoundTrip(request)
				Expect(err).NotTo(HaveOccurred())

				Expect(response.StatusCode).To(Equal(http.StatusBadRequest))

				body, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal("missing required query parameter: os=[osx,linux]"))
			})
		})
	})
})
