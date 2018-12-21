package core

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type S3Instance struct{
	Connection *s3.S3
}

type bucketList struct{}

func NewS3Connection() (*S3Instance, error){

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		exitErrorf("Unable to load a new session, %v", err)
		return nil, err
	}

	// Create S3 service client
	var svcInstance S3Instance
	svc := s3.New(sess)
	svcInstance.Connection = svc

	return &svcInstance, nil
}

func (s *S3Instance)ListS3Buckets() ([]string,error){

	var bucketList []string

	result, err := s.Connection.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
		return nil, err
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
			bucketList = append(bucketList, aws.StringValue(b.Name))
	}

	return bucketList, nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

