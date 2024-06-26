package encode

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

func EncodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func EncodeBytesString(v string) []byte {
	decoded, err := hex.DecodeString(v)
	if err != nil {
		panic(err)
	}
	return decoded
}

func EncodeStringArray(arr []string) []byte {
	var res [][]byte
	for _, v := range arr {
		res = append(res, []byte(v))
	}
	return bytes.Join(res, nil)
}

func EncodeString(v string) []byte {
	return []byte(v)
}

func EncodeAddress(v string) []byte {
	return common.HexToAddress(v).Bytes()
}

func EncodeUint256(v string) []byte {
	bn := new(big.Int)
	bn.SetString(v, 10)
	return math.U256Bytes(bn)
}

func EncodeUint256Array(arr []string) []byte {
	var res [][]byte
	for _, v := range arr {
		b := EncodeUint256(v)
		res = append(res, b)
	}

	return bytes.Join(res, nil)
}
