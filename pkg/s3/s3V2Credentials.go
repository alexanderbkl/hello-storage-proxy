package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/Hello-Storage/hello-storage-proxy/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	s3V2 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// CustomEndpointResolver implements aws.EndpointResolver to provide a custom S3 endpoint.
type CustomEndpointResolver struct {
	Endpoint string
	Region   string
}

// ResolveEndpoint resolves the endpoint based on the service and region provided.
func (c CustomEndpointResolver) ResolveEndpoint(service, region string) (aws.Endpoint, error) {
	log.Printf(c.Region)
	if service == s3V2.ServiceID {
		return aws.Endpoint{
			URL:           c.Endpoint,
			SigningRegion: c.Region,
		}, nil
	}
	return aws.Endpoint{}, fmt.Errorf("unknown service: %s", service)
}

func GetS3V2Client() (*s3V2.Client, error) {
	awsCredentials := credentials.NewStaticCredentialsProvider(
		config.Env().StorageAccessKey,
		config.Env().StorageSecretKey,
		"",
	)

	ctx := context.TODO()

	// Use custom endpoint resolver
	customResolver := CustomEndpointResolver{
		Endpoint: config.Env().StorageEndpoint,
		Region:   config.Env().StorageRegion,
	}

	s3Config, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithCredentialsProvider(awsCredentials),
		awsConfig.WithRegion(config.Env().StorageRegion),
		awsConfig.WithEndpointResolver(customResolver), // Use custom endpoint resolver
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3V2Client := s3V2.NewFromConfig(s3Config)

	if s3V2Client == nil {
		return nil, fmt.Errorf("failed to create S3 client")
	}

	return s3V2Client, nil
}
