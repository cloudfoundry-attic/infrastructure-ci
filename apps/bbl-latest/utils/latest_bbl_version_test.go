package utils_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/cloudfoundry/infrastructure-ci/apps/bbl-latest/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	expectedLastModified = "Thu, 20 Apr 2017 16:57:57 GMT"
)

var _ = Describe("LatestBBLVersion", func() {
	var (
		server *httptest.Server

		okCallCount          int
		notModifiedCallCount int
	)

	BeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			switch req.URL.Path {
			case "/repos/cloudfoundry/bosh-bootloader/releases/latest":
				if req.Header.Get("If-Modified-Since") == expectedLastModified {
					notModifiedCallCount++
					w.WriteHeader(http.StatusNotModified)
					return
				} else {
					okCallCount++
					w.Header().Set("Last-Modified", expectedLastModified)
					w.Write([]byte(`{"tag_name": "some-version"}`))
				}
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
		}))

		utils.SetEndpoint(server.URL)
	})

	AfterEach(func() {
		okCallCount = 0
		notModifiedCallCount = 0
		utils.ResetEndpoint()
	})

	It("returns the latest bbl version", func() {
		version, _, err := utils.LatestBBLVersion("", "")
		Expect(err).NotTo(HaveOccurred())
		Expect(version).To(Equal("some-version"))
	})

	It("returns cached version if unmodified", func() {
		version, lastModified, err := utils.LatestBBLVersion("", "")
		Expect(err).NotTo(HaveOccurred())
		Expect(lastModified).To(Equal(expectedLastModified))

		Expect(version).To(Equal("some-version"))
		Expect(okCallCount).To(Equal(1))

		version, lastModified, err = utils.LatestBBLVersion(version, lastModified)
		Expect(err).NotTo(HaveOccurred())
		Expect(lastModified).To(Equal(expectedLastModified))

		Expect(version).To(Equal("some-version"))
		Expect(okCallCount).To(Equal(1))
		Expect(notModifiedCallCount).To(Equal(1))
	})

	It("returns updated version if modified", func() {
		version, lastModified, err := utils.LatestBBLVersion("old-version", "Mon, 17 Apr 2017 10:00:00 GMT")
		Expect(err).NotTo(HaveOccurred())
		Expect(lastModified).To(Equal(expectedLastModified))

		Expect(version).To(Equal("some-version"))
		Expect(okCallCount).To(Equal(1))
		Expect(notModifiedCallCount).To(Equal(0))
	})
})
