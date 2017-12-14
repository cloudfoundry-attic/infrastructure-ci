package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Name     interface{} `yaml:"name"`
	Releases []struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"releases"`
	Stemcells []struct {
		Alias   string `yaml:"alias"`
		OS      string `yaml:"os"`
		Version string `yaml:"version"`
	} `yaml:"stemcells"`
	InstanceGroups []struct {
		Instances    int      `yaml:"instances"`
		Name         string   `yaml:"name"`
		Lifecycle    string   `yaml:"lifecycle"`
		AZs          []string `yaml:"azs"`
		VMExtensions []string `yaml:"vm_extensions"`
		VMType       string   `yaml:"vm_type"`
		Stemcell     string   `yaml:"stemcell"`
		Networks     []struct {
			Name string `yaml:"name"`
		} `yaml:"networks"`
		Jobs []struct {
			Name    string `yaml:"name"`
			Release string `yaml:"release"`
		} `yaml:"jobs"`
	} `yaml:"instance_groups"`
	Properties struct {
		Consul struct {
			AcceptanceTests struct {
				BOSH struct {
					Target         string `yaml:"target"`
					Username       string `yaml:"username"`
					Password       string `yaml:"password"`
					DirectorCACert string `yaml:"director_ca_cert"`
				} `yaml:"bosh"`
				Jumpbox struct {
					PrivateKey string `yaml:"private_key,omitempty"`
					URL        string `yaml:"url,omitempty"`
				} `yaml:"jumpbox,omitempty"`
				ParallelNodes              int    `yaml:"parallel_nodes"`
				ConsulReleaseVersion       string `yaml:"consul_release_version"`
				LatestConsulReleaseVersion string `yaml:"latest_consul_release_version"`
				EnableTurbulenceTests      bool   `yaml:"enable_turbulence_tests"`
				WindowsClients             bool   `yaml:"windows_clients"`
			} `yaml:"acceptance_tests"`
		} `yaml:"consul"`
	} `yaml:"properties"`
	Update interface{} `yaml:"update"`
}

func main() {
	output, err := Generate(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, string(output))
}

func Generate(exampleManifestFilePath string) ([]byte, error) {
	contents, err := ioutil.ReadFile(exampleManifestFilePath)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	err = yaml.Unmarshal(contents, &manifest)
	if err != nil {
		return nil, err
	}

	manifest.Releases[0].Version = os.Getenv("CONSUL_RELEASE_VERSION")
	manifest.Stemcells[0].Version = os.Getenv("STEMCELL_VERSION")

	manifest.Properties.Consul.AcceptanceTests.BOSH.Target = os.Getenv("BOSH_ENVIRONMENT")
	manifest.Properties.Consul.AcceptanceTests.BOSH.Username = os.Getenv("BOSH_CLIENT")
	manifest.Properties.Consul.AcceptanceTests.BOSH.Password = os.Getenv("BOSH_CLIENT_SECRET")
	manifest.Properties.Consul.AcceptanceTests.BOSH.DirectorCACert = os.Getenv("BOSH_CA_CERT")

	jumpboxPrivateKey := os.Getenv("JUMPBOX_PRIVATE_KEY")
	if jumpboxPrivateKey != "" {
		privateKey, err := ioutil.ReadFile(jumpboxPrivateKey)
		if err != nil {
			return nil, err
		}
		manifest.Properties.Consul.AcceptanceTests.Jumpbox.PrivateKey = string(privateKey)
		manifest.Properties.Consul.AcceptanceTests.Jumpbox.URL = os.Getenv("JUMPBOX_URL")
	}

	manifest.Properties.Consul.AcceptanceTests.ConsulReleaseVersion = os.Getenv("CONSUL_RELEASE_VERSION")
	manifest.Properties.Consul.AcceptanceTests.LatestConsulReleaseVersion = os.Getenv("LATEST_CONSUL_RELEASE_VERSION")
	manifest.Properties.Consul.AcceptanceTests.EnableTurbulenceTests = (os.Getenv("ENABLE_TURBULENCE_TESTS") == "true")
	manifest.Properties.Consul.AcceptanceTests.WindowsClients = (os.Getenv("WINDOWS_CLIENTS") == "true")

	parallelNodes, err := strconv.Atoi(os.Getenv("PARALLEL_NODES"))
	if err != nil {
		return nil, err
	}
	manifest.Properties.Consul.AcceptanceTests.ParallelNodes = parallelNodes

	contents, err = yaml.Marshal(manifest)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
