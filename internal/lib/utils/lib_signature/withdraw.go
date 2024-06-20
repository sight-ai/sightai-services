package lib_signature

import (
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
)

func GetSignedWithdrawTicket(nonce uint, amount string) (string, error) {
	if PublicKeyBytes == nil {
		initialize()
	}

	msg := encodePacked(
		encodeUint256(strconv.FormatUint(uint64(nonce), 10)),
		encodeUint256(amount),
	)

	msgHash := crypto.Keccak256Hash(msg)
	fmt.Println("msgHash ", hexutil.Encode(msgHash.Bytes()))

	data := crypto.Keccak256Hash(encodePacked(
		[]byte("\x19Ethereum Signed Message:\n32"),
		msgHash.Bytes()))
	fmt.Println("data ", hexutil.Encode(data.Bytes()))

	signature, err := crypto.Sign(data.Bytes(), PrivateKey)
	if err != nil {
		return "", err
	}

	chainId, _ := strconv.Atoi(config.Cfg.SightChainId)
	r, s, v := toSignatureValues(signature, uint(chainId))

	newSig := toSignatureBytes(r, s, v)

	return hexutil.Encode(newSig), nil
}

var ChainIDNonceMap = map[uint]byte{97: 27, 3780: 27}
