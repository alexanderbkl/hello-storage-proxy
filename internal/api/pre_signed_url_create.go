package api

import (
	"net/http"

	"github.com/Hello-Storage/hello-storage-proxy/internal/config"
	"github.com/Hello-Storage/hello-storage-proxy/pkg/s3"
	s3V2 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

// GeneratePresignedURL function will generate a presigned URL for client to upload directly to S3

func GeneratePutPresignedObject(s3V2Client *s3V2.Client, router *gin.RouterGroup) {

	router.PUT(("presigned-url"), func(c *gin.Context) {

		presigner := s3.Presigner{
			PresignClient: s3V2.NewPresignClient(s3V2Client),
		}

		presignedHTTPRequest, err := presigner.PutObject(config.Env().StorageBucket, "objectKey", 60)

		log.Printf("Presigned URL: %s", presignedHTTPRequest.URL)
		log.Printf("Presigned Method: %s", presignedHTTPRequest.Method)
		log.Printf("Presigned Headers: %v", presignedHTTPRequest.SignedHeader)

		if err != nil {
			log.Errorf("failed to generate presigned URL, %v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"presigned_url": presignedHTTPRequest.URL,
			"method":        presignedHTTPRequest.Method,
			"headers":       presignedHTTPRequest.SignedHeader,
		})
	})

}
