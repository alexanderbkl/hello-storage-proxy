package web3

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func ValidateMessageSignature(walletAddress, signature string, message []byte) bool {
	if signature[0:2] != "0x" || len(signature) != 132 {
		log.Printf("invalid signature: %s", signature)
		return false
	}

	sig, err := hexutil.Decode(signature)
	if err != nil {
		log.Printf("failed to decode signature: %s", err)
		//print signature
		log.Printf("invalid signature: %s", signature)
		return false
	}

	message = accounts.TextHash(message)
	sig[crypto.RecoveryIDOffset] -= 27

	recovered, err := crypto.SigToPub(message, sig)
	if err != nil {
		log.Printf("failed to recover public key: %s", err)
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	//convert recoveredAddr to lower case string
	recoveredAddrStr := strings.ToLower(recoveredAddr.Hex())
	//convert walletAddress to lower case
	walletAddress = strings.ToLower(walletAddress)

	if walletAddress == recoveredAddrStr {
		return true
	} else {
		//print walletAddress and recoveredAddr
		//log.Printf("walletAddress: %s, recoveredAddr: %s", walletAddress, recoveredAddr.Hex())
		//log.Printf("invalid signature: %s", signature)
		return false
	}
}
