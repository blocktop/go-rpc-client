package rpcconsensus

import (
	rpcclient "github.com/blckit/go-rpc-client"
	"github.com/google/uuid"
)

func GetTree(format string) (*GetTreeReply, error) {
	params := []GetTreeArgs{GetTreeArgs{format}}
	req := &rpcclient.RPCRequest{
		Method: "Consensus.GetTree",
		Params: params,
		Id:     uuid.New().String()}

	res, err := rpcclient.SendRPCRequest(req)
	if err != nil {
		return nil, err
	}

	reply := GetTreeReply{Tree: res["Tree"].(string)}

	return &reply, nil
}

type GetTreeArgs struct {
	Format string
}

type GetTreeReply struct {
	Tree string
}
