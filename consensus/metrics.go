package rpcconsensus

import (
	rpcclient "github.com/blckit/go-rpc-client"
	"github.com/google/uuid"
)

func GetMetrics(format string) (*GetMetricsReply, error) {
	params := []GetMetricsArgs{GetMetricsArgs{format}}
	req := &rpcclient.RPCRequest{
		Method: "Consensus.GetMetrics",
		Params: params,
		Id:     uuid.New().String()}

	res, err := rpcclient.SendRPCRequest(req)
	if err != nil {
		return nil, err
	}
	
	reply := GetMetricsReply{Metrics: res["Metrics"].(string)}

	return &reply, nil
}

type GetMetricsArgs struct {
	Format string
}

type GetMetricsReply struct {
	Metrics string
}
