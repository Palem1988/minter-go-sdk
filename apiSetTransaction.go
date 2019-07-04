package mintersdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Ответ транзакции

//easyjson:json
type send_transaction struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      string            `json:"id"`
	Result  TransSendResponse `json:"result"`
	Error   ErrorStruct       `json:"error"`
}
type TransSendResponse struct {
	Code int    `json:"code" bson:"code" gorm:"code" db:"code"`
	Log  string `json:"log" bson:"log" gorm:"log" db:"log"`
	Data string `json:"data" bson:"data" gorm:"data" db:"data"`
	Hash string `json:"hash" bson:"hash" gorm:"hash" db:"hash"`
}

// Исполнение транзакции закодированной RLP
func (c *SDK) SetTransaction(strRlpEnc string) (string, error) {
	fmt.Println("TX RLP:", strRlpEnc)

	url := fmt.Sprintf("%s/send_transaction?tx=0x%s", c.MnAddress, strRlpEnc)
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
	//json.Unmarshal(body, &data) -- переход на easyjson

	err = data.UnmarshalJSON(body)
	if err != nil {
		panic(err)
	}

	if data.Error.Code != 0 {
		err = errors.New(fmt.Sprint(data.Error.Code, " - ", data.Error.Message, " - ", data.Error.TxResult.Log))
		return "", err
	}

	if data.Result.Code == 0 {
		return fmt.Sprintf("Mt%s", strings.ToLower(data.Result.Hash)), nil
	} else {
		fmt.Printf("ERROR: TxSign: %#v\n", data)
		return data.Result.Log, errors.New(fmt.Sprintf("Err:%d-%s", data.Result.Code, data.Result.Log))
	}
}
