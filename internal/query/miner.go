package query

import (
	"math"
	"time"

	"github.com/Hello-Storage/hello-storage-proxy/internal/config"
)

// Constants
const initialRewardRate = 1.0             // Starting tokens per hour
const targetDurationHours = 10 * 365 * 24 // 10 years in hours
const smallestMeaningfulRate = 1.0        // Smallest meaningful reward rate

// Calculate decay constant τ (tau)
var tau = float64(targetDurationHours) / math.Log(initialRewardRate/smallestMeaningfulRate)

// calculateRewardRate computes the reward rate at the given currentTime.
func calculateRewardRate(epochZero, currentTime int64) float64 {
	// Calculate time elapsed since mining began, in hours
	timeElapsed := float64(currentTime-epochZero) / 3600.0

	// Calculate the reward rate using logarithmic decay
	rewardRate := initialRewardRate * math.Exp(-timeElapsed/tau)

	// Round to a realistic token issuance amount
	//rewardRate = math.Round(rewardRate*10000) / 10000 // Round to 4 decimal places

	// Ensure the reward rate doesn't drop below a minimum value
	if rewardRate < 0.000000000001 {
		rewardRate = 0.000000000001
	}

	return rewardRate
}

func GetGlobalRewardRate() float64 {
	epochZero := config.Env().EpochZero
	currentTime := time.Now().Unix() // Current timestamp

	//log epochZero and currentTime
	log.Printf("Epoch Zero: %d\n", epochZero)
	log.Printf("Current Time: %d\n", currentTime)

	rewardRate := calculateRewardRate(epochZero, currentTime)

	log.Printf("Reward rate: %f tokens per hour\n", rewardRate)

	return rewardRate
}
