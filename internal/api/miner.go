package api

import (
	"net/http"

	"github.com/Hello-Storage/hello-storage-proxy/internal/constant"
	"github.com/Hello-Storage/hello-storage-proxy/internal/entity"
	"github.com/Hello-Storage/hello-storage-proxy/internal/query"
	"github.com/Hello-Storage/hello-storage-proxy/pkg/token"
	"github.com/gin-gonic/gin"
)

// LoadUser
//
// GET /api/load/miner
func LoadMiner(router *gin.RouterGroup) {
	router.GET("/load/miner", func(ctx *gin.Context) {
		authPayload := ctx.MustGet(constant.AuthorizationPayloadKey).(*token.Payload)

		u := query.FindUserWithWallet(authPayload.UserID)
		if u == nil {
			log.Errorf("user not found: %d", authPayload.UserID)
			ctx.JSON(http.StatusNotFound, "user not found")
			return
		}

		miner := query.FindMinerWithUserID(u.ID)

		// get global reward rate
		rewardRate := query.GetGlobalRewardRate()

		var resp = struct {
			Miner      *entity.Miner `json:"miner"`
			RewardRate float64       `json:"rewardRate"`
		}{
			Miner:      miner,
			RewardRate: rewardRate,
		}

		ctx.JSON(http.StatusOK, resp)
	})
}
