package s3

import (
	"context"

	"github.com/Hello-Storage/hello-storage-proxy/internal/config"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	s3V2 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3V2Client() (*s3V2.Client, error) {
	awsCredentials := credentials.NewStaticCredentialsProvider(
		config.Env().StorageAccessKey,
		config.Env().StorageSecretKey,
		"",
	)

	ctx := context.TODO()

	s3Config, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithCredentialsProvider(awsCredentials), awsConfig.WithRegion(config.Env().StorageRegion))
	if err != nil {
		return nil, err
	}

	s3V2Client := s3V2.NewFromConfig(s3Config)

	return s3V2Client, nil
}
