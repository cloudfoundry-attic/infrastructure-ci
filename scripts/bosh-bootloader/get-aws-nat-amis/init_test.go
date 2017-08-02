package main_test

import (
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
	var err error
	natBinaryPath, err = gexec.Build("github.com/cloudfoundry/infrastructure-ci/scripts/bosh-bootloader/get-aws-nat-amis")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
