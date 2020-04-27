package awscredentials

import (
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
)

func createSession() *session.Session {
	session, err := session.NewSession()
	if err != nil {
		L.Logging.Error(err)
		panic(err)
	}
	return session
}

var curSession = createSession()
var creds = curSession.Config.Credentials
var signer = v4.NewSigner(creds)

// https://docs.aws.amazon.com/sdk-for-go/api/aws/ec2metadata/#EC2Metadata.Region
func GetSigner() *v4.Signer {
	if creds.IsExpired() == true { // refreshes credentials
		curSession = createSession()
		creds = curSession.Config.Credentials
		signer = v4.NewSigner(creds)
	}
	return signer
}
