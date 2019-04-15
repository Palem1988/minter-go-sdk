package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Делегирования
type TxDelegateData struct {
	PubKey string
	Coin   string
	Stake  float32
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Делегирование
func (c *SDK) TxDelegate(t *TxDelegateData) (string, error) {
	coin := getStrCoin(t.Coin)
	coinGas := getStrCoin(t.GasCoin)
	value := bip2pip_f64(float64(t.Stake))
	valueGas := big.NewInt(t.GasPrice)
	pubkey := publicKey2Byte(t.PubKey)
	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.DelegateData{
		PubKey: pubkey,
		Coin:   coin,
		Value:  value,
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	_, nowNonce, err := c.GetAddress(c.AccAddress)
	if err != nil {
		return "", err
	}

	if c.ChainMainnet {
		ChainID = ChainMainnet
	} else {
		ChainID = ChainTestnet
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		ChainID:       ChainID,
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeDelegate,
		Data:          encodedData,
		Payload:       []byte(t.Payload),
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
