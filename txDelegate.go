package mintersdk

import (
	"encoding/hex"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
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

func (c *SDK) TxDelegateRLP(t *TxDelegateData) (string, error) {
	coin := getStrCoin(t.Coin)
	coinGas := getStrCoin(t.GasCoin)
	value := bip2pip_f64(float64(t.Stake))
	valueGas := uint32(t.GasPrice)
	pubkey := publicKey2Byte(t.PubKey)
	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	if c.AccAddress == "" {
		c.AccAddress, err = GetAddressPrivateKey(c.AccPrivateKey)
		if err != nil {
			return "", err
		}
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

	var _ChainID types.ChainID
	if c.ChainMainnet {
		_ChainID = types.ChainMainnet
	} else {
		_ChainID = types.ChainTestnet
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		ChainID:       _ChainID,
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

	encodedTx, err := tx.Serialize()
	if err != nil {
		return "", err
	}

	strTxRPL := hex.EncodeToString(encodedTx)

	strRlpEnc := string(strTxRPL)

	return strRlpEnc, err
}

// Транзакция - Делегирование
func (c *SDK) TxDelegate(t *TxDelegateData) (string, error) {
	strRlpEnc, err := c.TxDelegateRLP(t)
	if err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(strRlpEnc)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
