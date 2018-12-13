package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// запрос на всех кандидатов (curl -s 'localhost:8841/api/candidates')
type node_candidates struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []CandidateInfo
	Error   ErrorStruct
}

// type CandidateInfo struct --- в apiGetCandidate.go

// Возвращает список нод валидаторов и кандидатов
func (c *SDK) GetCandidates() ([]CandidateInfo, error) {
	url := fmt.Sprintf("%s/candidates", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		return []CandidateInfo{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []CandidateInfo{}, err
	}

	var data node_candidates
	json.Unmarshal(body, &data)

	for i1, _ := range data.Result {
		data.Result[i1].TotalStake = pipStr2bip_f32(data.Result[i1].TotalStakeTx)
		data.Result[i1].Commission, _ = strconv.Atoi(data.Result[i1].CommissionTx)
		data.Result[i1].CreatedAtBlock, _ = strconv.Atoi(data.Result[i1].CreatedAtBlockTx)
	}
	return data.Result, nil
}
