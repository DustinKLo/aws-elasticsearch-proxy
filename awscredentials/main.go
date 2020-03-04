package awscredentials

import (
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
)

// https://docs.aws.amazon.com/sdk-for-go/api/aws/ec2metadata/#EC2Metadata.Region
func GetAwsSigner() *v4.Signer {
	session, err := session.NewSession()
	if err != nil {
		L.Logging.Error(err)
		panic(err)
	}

	creds := session.Config.Credentials
	return v4.NewSigner(creds)
}
