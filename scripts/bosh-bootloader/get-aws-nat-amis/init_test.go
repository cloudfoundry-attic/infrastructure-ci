package main_test

import (
	"os"
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAWSNATAMIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scripts/bosh-bootloader/get-aws-nat-amis")
}

var natBinaryPath string

var _ = BeforeSuite(func() {
	key := os.Getenv("AWS_ACCESS_KEY_ID")
	if key == "" {
		Fail("please set AWS_ACCESS_KEY_ID before running tests")
	}

	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret == "" {
		Fail("please set AWS_SECRET_ACCESS_KEY before running tests")
	}

	var err error
	natBinaryPath, err = gexec.Build("github.com/cloudfoundry/infrastructure-ci/scripts/bosh-bootloader/get-aws-nat-amis")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
