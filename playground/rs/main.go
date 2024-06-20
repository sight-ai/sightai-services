package main

import (
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/storyicon/sigverify"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//addr 0x882304271ee4851133005f817af762f97d9dbd07, sig 0x30032cefd919b22637de66dd1ecdf1af09b4e4239ce5d9105fb28a3c7ef1765c64029c8a5f31465fc68f971a10d597d6e7765926f55b94f2267043cb8c755f471c, data {"from_currency_id":1,"to_currency_id":2,"amount_in":"0.01","amount_out_min":"223.59960996099608","to_address":"0x882304271Ee4851133005f817AF762f97D9dbd07","nonce":102}

func main() {
	data := `{"order_ids":[82736940],"nonce":16}`
	addr := "0x3A8cd1C3bae091ba181256845b3Ae59693f12207"
	sig := "0x0fdd5981cbee4d7c71aaf5c3b5264f411f0388a800d6a19476b26609b7c3b5ce765028496cec76c95e254876ecfea425787f57fe40c9c8fde7f2e44148fb20a91c"

	fmt.Printf("addr %s, sig %s, data %s\n", addr, sig, data)
	b, err := sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(addr),
		[]byte(data),
		sig,
	)

	checkErr(err)
	fmt.Println(b)
}
