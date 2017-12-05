package main

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generate", func() {
	var variables map[string]string

	BeforeEach(func() {
		variables = map[string]string{
			"BOSH_ENVIRONMENT":              "some-bosh-target",
			"BOSH_CLIENT":                   "some-bosh-username",
			"BOSH_CLIENT_SECRET":            "some-bosh-password",
			"BOSH_CA_CERT":                  "some-bosh-director-ca-cert",
			"PARALLEL_NODES":                "10",
			"CONSUL_RELEASE_VERSION":        "some-consul-release-version",
			"STEMCELL_VERSION":              "some-stemcell-version",
			"LATEST_CONSUL_RELEASE_VERSION": "some-latest-consul-release-version",
			"ENABLE_TURBULENCE_TESTS":       "true",
			"WINDOWS_CLIENTS":               "true",
		}

		for name, value := range variables {
			variables[name] = os.Getenv(name)
			os.Setenv(name, value)
		}
	})

	AfterEach(func() {
		for name, value := range variables {
			os.Setenv(name, value)
		}
	})

	It("generates a manifest", func() {
		expectedManifest, err := ioutil.ReadFile("fixtures/expected.yml")
		Expect(err).NotTo(HaveOccurred())

		manifest, err := Generate("fixtures/example.yml")
		Expect(err).NotTo(HaveOccurred())

		Expect(manifest).To(MatchYAML(expectedManifest))
	})

	Context("failure cases", func() {
		It("returns an error when the parallel nodes is not an int", func() {
			os.Setenv("PARALLEL_NODES", "not an int")
			_, err := Generate("fixtures/example.yml")
			Expect(err).To(MatchError(ContainSubstring(`parsing "not an int": invalid syntax`)))
		})

		It("returns an error when the example manifest does not exist", func() {
			_, err := Generate("fixtures/doesnotexist.yml")
			Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
		})

		It("returns an error when the example manifest is malformed", func() {
			_, err := Generate("fixtures/malformed.yml")
			Expect(err).To(MatchError(ContainSubstring("cannot unmarshal !!str `this is...`")))
		})
	})
})
