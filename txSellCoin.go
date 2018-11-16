package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Продажи монет
type TxSellCoinData struct {
	CoinToSell  string
	CoinToBuy   string
	ValueToSell float32
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Продажи монет
func (c *SDK) TxSellCoin(t *TxSellCoinData) (string, error) {
	coinBuy := getStrCoin(t.CoinToBuy)
	coinSell := getStrCoin(t.CoinToSell)
	value := bip2pip_f64(float64(t.ValueToSell))
	coinGas := getStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.SellCoinData{
		CoinToSell:  coinSell,
		ValueToSell: value,
		CoinToBuy:   coinBuy,
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeSellCoin,
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
