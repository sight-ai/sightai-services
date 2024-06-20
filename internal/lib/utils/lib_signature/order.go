package lib_signature

import (
	"encoding/json"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

func VerifyDataSig(address, signature string, data []byte) (bool, error) {
	log.Info().Msgf("addr %s, sig %s, data %s", address, signature, string(data))
	return sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(address),
		data,
		signature,
	)
}

func VerifyTypedDataSig(address, signature string, data []byte) (bool, error) {
	var typedData apitypes.TypedData
	if err := json.Unmarshal(data, &typedData); err != nil {
		return false, err
	}
	return sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress(address),
		typedData,
		signature,
	)
}
