package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Создания монеты
type TxCreateCoinData struct {
	Name                 string
	Symbol               string
	InitialAmount        int64
	InitialReserve       int64
	ConstantReserveRatio uint
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Создание монеты
func (c *SDK) TxCreateCoin(t *TxCreateCoinData) (string, error) {
	toCreate := getStrCoin(t.Symbol)
	reserve := bip2pip_i64(t.InitialReserve)
	amount := bip2pip_i64(t.InitialAmount)
	coinGas := getStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.CreateCoinData{
		Name:                 t.Name,
		Symbol:               toCreate,
		InitialAmount:        amount,
		InitialReserve:       reserve,
		ConstantReserveRatio: t.ConstantReserveRatio,
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	nowNonce, err := c.GetNonce(c.AccAddress)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeCreateCoin,
		Data:          encodedData,
		SignatureType: tr.SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(&tx)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
