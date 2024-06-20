package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/playground/rs/sha/encode"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/storyicon/sigverify"
	"math/big"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 0xc0271BDA95f78EF80728152eE9B6c5A915E91DA5 (address)
// 0x3d256e3d56dd084baf0b23de9d6b67a181a02e7431880f84850efc9a75152223 (private key)

//Vault deployed to 0x5FbDB2315678afecb367f032d93F642f64180aa3
//BUSD deployed to 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
//WBTC deployed to 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
//WETH deployed to 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9

type SignInRequest struct {
	WalletAddress string `json:"wallet_address"`
}

func main() {
	//address := "0xc0271BDA95f78EF80728152eE9B6c5A915E91DA5"
	//sig := "0x191c5bbb176070a37089cfa50b545271205647721ec00d393c6af1723e9eb53e3766eeb84cfc4d1e0c6cac8e640decfd573d6e2fe6a68655204de225c2f0970c1c"
	privateKey, err := crypto.HexToECDSA("10667ccdfed55f1e03481b6c8dffa2ff29f357de261226f520dcf0210a374bef")
	checkErr(err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKeyBytes ", hexutil.Encode(publicKeyBytes))

	nonce := strconv.FormatUint(uint64(2), 10)
	amount := "100100000000000000000"

	msg := encode.EncodePacked(
		encode.EncodeUint256(nonce),
		encode.EncodeUint256(amount),
	)

	msgHash := crypto.Keccak256Hash(msg)
	fmt.Println("msgHash ", hexutil.Encode(msgHash.Bytes()))

	data := crypto.Keccak256Hash(encode.EncodePacked(
		[]byte("\x19Ethereum Signed Message:\n32"),
		msgHash.Bytes()))
	fmt.Println("data ", hexutil.Encode(data.Bytes()))

	signature, err := crypto.Sign(data.Bytes(), privateKey)
	checkErr(err)
	fmt.Println("signature ", hexutil.Encode(signature))

	r, s, v := toSignatureValues(signature)
	fmt.Println("r ", r)
	fmt.Println("s ", s)
	fmt.Println("v ", v)

	newSig := toSignatureBytes(r, s, v)
	fmt.Println("signature with ethID ", hexutil.Encode(newSig))

	verifySig(publicKeyBytes, signature, data)

	//r := &SignInRequest{
	//	WalletAddress: address,
	//}
	//b, err := json.Marshal(r)
	//checkErr(err)
	//var typedData apitypes.TypedData
	//if err := json.Unmarshal(b, &typedData); err != nil {
	//	panic(err)
	//}
	//valid, err := VerifyTypedDataSig(address, sig, b)
	//checkErr(err)
	//fmt.Println(valid)
}

var ChainIDNonceMap = map[uint]byte{97: 27}

func toSignatureValues(sig []byte) (r, s, v *big.Int) {
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + ChainIDNonceMap[97]})

	return r, s, v
}

func toSignatureBytes(r, s, v *big.Int) []byte {
	sig := make([]byte, crypto.SignatureLength)

	copy(sig, r.Bytes())
	copy(sig[32:], s.Bytes())
	sig[64] = byte(v.Uint64())

	return sig
}

func verifySig(publicKey, signature []byte, hash common.Hash) {
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	checkErr(err)
	fmt.Println("sigPublicKey ", hexutil.Encode(sigPublicKey))

	matches := bytes.Equal(sigPublicKey, publicKey)
	fmt.Println("matches: ", matches) // true
}

func VerifyTypedDataSig(address, signature string, data []byte) (bool, error) {
	return sigverify.VerifyEllipticCurveHexSignatureEx(
		common.HexToAddress(address),
		data,
		signature,
	)
}
