package lib_signature

import (
	"crypto/ecdsa"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/ethereum/go-ethereum/crypto"
)

var PrivateKey *ecdsa.PrivateKey
var PublicKeyBytes []byte

func initialize() {
	var err error
	PrivateKey, err = crypto.HexToECDSA(config.Cfg.PrivateKey)
	if err != nil {
		log.Error().Err(err).Msg("error casting public key to ECDSA")
	}

	publicKey := PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Msg("error casting public key to ECDSA")
	}
	PublicKeyBytes = crypto.FromECDSAPub(publicKeyECDSA)
}
