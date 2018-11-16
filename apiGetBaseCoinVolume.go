package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_basecoincol struct {
	Code   int
	Result base_volume
}

type base_volume struct {
	Volume string `json:"volume" bson:"-" gorm:"-"`
}

// Возвращает количество базовой монеты (BIP или MNT), существующей в сети.
// Он подсчитывает награды блоков, премиальные и ретранслируемые награды.
func (c *SDK) GetBaseCoinVolume(height int) float32 {
	url := fmt.Sprintf("%s/api/bipVolume?height=%d", c.MnAddress, height)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_basecoincol
	json.Unmarshal(body, &data)

	return pipStr2bip_f32(data.Result.Volume)
}
