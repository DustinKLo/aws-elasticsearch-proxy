package awscredentials

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
)

func GetAwsSigner() *v4.Signer {
	provider := credentials.SharedCredentialsProvider{}

	fileCreds, err := provider.Retrieve()
	if err != nil {
		L.Logging.Error(err)
		panic("Cannot retrieve credentials from $HOME/.aws/credentials")
	}

	L.Logging.Info("AWS Credentials:", fileCreds)
	os.Setenv("AWS_ACCESS_KEY_ID", fileCreds.AccessKeyID) // setting the environment variables
	os.Setenv("AWS_SECRET_ACCESS_KEY", fileCreds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", fileCreds.SessionToken)

	creds := credentials.NewEnvCredentials() // taken from AWS environment variables

	return v4.NewSigner(creds)
}
