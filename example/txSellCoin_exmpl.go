package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "...",
	}

	slDt := m.TxSellCoinData{
		CoinToSell:  "MNT",
		CoinToBuy:   "ABCDEF24",
		ValueToSell: 10,
		GasCoin:     "MNT",
		GasPrice:    1,
	}

	resHash, err := sdk.TxSellCoin(&slDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
