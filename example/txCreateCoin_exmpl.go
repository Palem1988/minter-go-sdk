package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccPrivateKey: "...",
	}

	Gas, _ := sdk.GetMinGas()

	creatDt := m.TxCreateCoinData{
		Name:                 "Test coin 24",
		Symbol:               "ABCDEF24",
		InitialAmount:        100,
		InitialReserve:       100,
		ConstantReserveRatio: 50,
		// Gas
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxCreateCoin(&creatDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
