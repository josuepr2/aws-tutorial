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

func NewS3Connection() (*s3.S3, error){

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		exitErrorf("Unable to load a new session, %v", err)
		return nil, err
	}

	// Create S3 service client
	svc := s3.New(sess)
	return svc, nil
}

func (s3 *S3Instance)ListS3Buckets(){

	result, err := s3.Connection.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

