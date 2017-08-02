package main_test

import (
	"encoding/json"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

type AMI struct {
	CreationDate string `json:"creationDate"`
	ID           string `json:"id"`
}

var _ = Describe("get-aws-nat-amis", func() {
	It("returns a JSON list of AMIs", func() {
		command := exec.Command(
			natBinaryPath,
			"--key", os.Getenv("AWS_ACCESS_KEY_ID"),
			"--secret", os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"--region", "us-west-1",
		)

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "120s").Should(gexec.Exit(0))

		var natAMIMap map[string]string
		err = json.Unmarshal(session.Out.Contents(), &natAMIMap)
		Expect(err).NotTo(HaveOccurred())

		Expect(natAMIMap).To(HaveKeyWithValue("us-west-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("us-west-2", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("us-east-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("us-east-2", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("sa-east-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("eu-west-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("eu-central-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("ap-southeast-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("ap-southeast-2", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("ap-northeast-1", MatchRegexp("ami-[0-9a-f]+")))
		Expect(natAMIMap).To(HaveKeyWithValue("ap-northeast-2", MatchRegexp("ami-[0-9a-f]+")))
	})
})
