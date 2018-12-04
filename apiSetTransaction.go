package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bytes"
	"encoding/hex"
	"errors"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Ответ транзакции
type send_transaction struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  TransSendResponse
}
type TransSendResponse struct {
	Code int    `json:"code" bson:"code" gorm:"code"`
	Log  string `json:"log" bson:"log" gorm:"log"`
	Data string `json:"data" bson:"data" gorm:"data"`
	Hash string `json:"hash" bson:"hash" gorm:"hash"`
}

// Исполнение транзакции закодированной RLP
func (c *SDK) SetTransaction(tx *tr.Transaction) (string, error) {

	encodedTx, err := tx.Serialize()
	if err != nil {
		fmt.Println("ERROR: SetCandidateTransaction::tx.Serialize")
		return "", err
	}

	strTxRPL := hex.EncodeToString(encodedTx)

	message := map[string]interface{}{
		"transaction": strTxRPL,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		fmt.Println("ERROR: SetCandidateTransaction::json.Marshal")
		return "", err
	}

	/*url := fmt.Sprintf("%s/api/sendTransaction", c.MnAddress)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))*/
	url := fmt.Sprintf("%s/send_transaction?tx=%s", c.MnAddress, string(bytesRepresentation))
	res, err := http.Get(url)
	if err != nil {
		//fmt.Println("ERROR: TxSign::http.Post")
		fmt.Println("ERROR: TxSign::http.Get")
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR: TxSign::ioutil.ReadAll")
		return "", err
	}

	var data send_transaction
	json.Unmarshal(body, &data)

	if data.Result.Code == 0 {
		return data.Result.Hash, nil
	} else {
		fmt.Printf("ERROR: TxSign: %#v\n", data)
		return data.Result.Log, errors.New(fmt.Sprintf("Err:%d-%s", data.Result.Code, data.Result.Log))
	}
}
