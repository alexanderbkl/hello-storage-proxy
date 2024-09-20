package server

import (
	"github.com/Hello-Storage/hello-storage-proxy/internal/api"
	"github.com/Hello-Storage/hello-storage-proxy/internal/config"
	"github.com/Hello-Storage/hello-storage-proxy/internal/middlewares"
	"github.com/Hello-Storage/hello-storage-proxy/pkg/token"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	var APIv1 *gin.RouterGroup
	var AuthAPIv1 *gin.RouterGroup
	tokenMaker, err := token.NewPasetoMaker(config.Env().TokenSymmetricKey)
	if err != nil {
		log.Errorf("cannot create token maker: %s", err)
		panic(err)
	}

	// Create router groups.
	APIv1 = router.Group("/api")
	AuthAPIv1 = router.Group("/api")
	AuthAPIv1.Use(middlewares.AuthMiddleware(tokenMaker))

	// routes
	api.Ping(APIv1)
	/*
		//api keys routes
		api.ApiKey(AuthAPIv1, tokenMaker)
		// auth routes
		api.LoginUser(APIv1, tokenMaker)
		api.RenewAccessToken(APIv1, tokenMaker)
		api.OAuthGoogle(APIv1, tokenMaker)
		api.RequestNonce(APIv1)
		api.StartOTP(APIv1)
		api.VerifyOTP(APIv1, tokenMaker)

		// user routes
		api.LoadUser(AuthAPIv1)
		api.GetUserDetail(AuthAPIv1)
	*/

	// file routes
	/*
		FileRoutes := AuthAPIv1.Group("/file")
		api.GetFile(FileRoutes)
		api.PutUploadFiles(FileRoutes)
		api.CreateFile(FileRoutes)
		api.DeleteFile(FileRoutes)
		api.DownloadFile(FileRoutes)
		api.DownloadMultipartFile(FileRoutes)
		api.UpdateFileRoot(FileRoutes)
		api.CheckFilesExistInPool(FileRoutes)
		api.GetShareState(FileRoutes)
		api.PublishFile(FileRoutes)
		api.UnpublishFile(FileRoutes)
		api.GetPublishedFile(FileRoutes)
		api.EncryptFile(FileRoutes)
		api.UploadFileMultipart(FileRoutes)
		api.UpdateFileIpfs(FileRoutes)

		api.GetPublishedFileName(router.Group("/api/file"))

		// folder routes
		api.SearchFolderByRoot(AuthAPIv1)
		api.CreateFolder(AuthAPIv1)
		api.GetFolderFiles(AuthAPIv1)
		api.DownloadFolder(AuthAPIv1)
		api.DownloadMultipartFolder(AuthAPIv1)
		api.DeleteFolder(AuthAPIv1)
		api.UpdateFolderRoot(AuthAPIv1)

	*/

}
