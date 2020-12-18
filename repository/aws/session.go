package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// createSession connects to AWS S3 and returns the new session.
func createSession(region string) *session.Session {
	s := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return s
}
