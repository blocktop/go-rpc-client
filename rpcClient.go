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
