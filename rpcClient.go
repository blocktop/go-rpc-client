// Copyright © 2018 J. Strobus White.
// This file is part of the blocktop blockchain development kit.
//
// Blocktop is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Blocktop is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with blocktop. If not, see <http://www.gnu.org/licenses/>.

package rpcclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	Id     string      `json:"id"`
}

type RPCResponse struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
	Id     string      `json:"id"`
}

func SendRPCRequest(req interface{}) (map[string]interface{}, error) {
	reqb, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("http://localhost:%d/rpc", viper.GetInt("rpc.port"))
	res, err := http.Post(url, "application/json", strings.NewReader(string(reqb)))
	if err != nil {
		return nil, err
	}
	resb, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	var data RPCResponse
	err = json.Unmarshal(resb, &data)
	if err != nil {
		return nil, err
	}
	if data.Error != "" {
		return nil, errors.New(data.Error)
	}

	return data.Result.(map[string]interface{}), nil
}
