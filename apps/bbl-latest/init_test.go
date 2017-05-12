package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestBBLLatest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "bbl-latest")
}

var (
	pathToBBLLatest string
)

var _ = BeforeSuite(func() {
	var err error
	pathToBBLLatest, err = gexec.Build("github.com/cloudfoundry/infrastructure-ci/apps/bbl-latest")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func bblLatest() *gexec.Session {
	cmd := exec.Command(pathToBBLLatest)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}

func openPort() (string, error) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", err
	}

	defer l.Close()
	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return "", err
	}

	return port, nil
}

func waitForServerToStart(port string) {
	timer := time.After(0 * time.Second)
	timeout := time.After(10 * time.Second)
	for {
		select {
		case <-timeout:
			panic("Failed to boot!")
		case <-timer:
			_, err := http.Get("http://localhost:" + port)
			if err == nil {
				return
			}

			timer = time.After(2 * time.Second)
		}
	}
}

func getLatestBBLVersion() string {
	var latestJson struct {
		TagName string `json:"tag_name"`
	}
	accessToken := os.Getenv("GITHUB_OAUTH_TOKEN")
	response, err := http.Get(fmt.Sprintf("https://api.github.com/repos/cloudfoundry/bosh-bootloader/releases/latest?access_token=%s", accessToken))
	Expect(err).NotTo(HaveOccurred())
	Expect(response.StatusCode).To(Equal(http.StatusOK))
	defer response.Body.Close()

	bodyContents, err := ioutil.ReadAll(response.Body)
	Expect(err).NotTo(HaveOccurred())

	err = json.Unmarshal(bodyContents, &latestJson)
	Expect(err).NotTo(HaveOccurred())

	latestVersion := latestJson.TagName
	if latestVersion == "" {
		fmt.Printf("latestVersion is empty: GitHub Response Body:\n%s\n", string(bodyContents))
	}
	Expect(latestVersion).To(MatchRegexp(`v[0-9]+\.[0-9]+\.[0-9]+`))
	return latestVersion
}
