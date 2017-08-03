package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	key := flag.String("key", os.Getenv("AWS_ACCESS_KEY_ID"), "AWS access key ID")
	secret := flag.String("secret", os.Getenv("AWS_SECRET_ACCESS_KEY"), "AWS secret access key")
	region := flag.String("region", "us-west-1", "AWS region")

	flag.Parse()

	awsCredentials := credentials.NewCredentials(&credentials.StaticProvider{
		credentials.Value{
			AccessKeyID:     *key,
			SecretAccessKey: *secret,
		},
	})

	awsSession := session.Must(session.NewSession(&aws.Config{
		Credentials: awsCredentials,
		Region:      region,
	}))

	ec2Client := ec2.New(awsSession)

	awsRegionsOutput, err := ec2Client.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		log.Fatalf("failed describing regions: %s", err) //not tested
	}

	nameFilter := &ec2.Filter{
		Name:   aws.String("name"),
		Values: []*string{aws.String("amzn-ami-vpc-nat-hvm*")},
	}

	AMIs := make(map[string]string)

	for _, region := range awsRegionsOutput.Regions {
		awsSession := session.Must(session.NewSession(&aws.Config{
			Credentials: awsCredentials,
			Region:      region.RegionName,
		}))

		ec2Client := ec2.New(awsSession)

		describeImagesInput := ec2.DescribeImagesInput{
			Filters: []*ec2.Filter{nameFilter},
			Owners:  []*string{aws.String("amazon")},
		}

		imagesOutput, err := ec2Client.DescribeImages(&describeImagesInput)
		if err != nil {
			log.Fatalf("failed describing images: %s", err) //not tested
		}

		sort.Sort(ImageSlice(imagesOutput.Images))

		AMIs[*region.RegionName] = *imagesOutput.Images[0].ImageId
	}

	err = json.NewEncoder(os.Stdout).Encode(AMIs)
	if err != nil {
		log.Fatalf("failed encoding json: %s", err) //not tested
	}
}

type ImageSlice []*ec2.Image

func (is ImageSlice) Len() int {
	return len(is)
}

func (is ImageSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is ImageSlice) Less(i, j int) bool {
	return *is[i].CreationDate > *is[j].CreationDate
}
