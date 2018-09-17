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

package rpcconsensus

import (
	rpcclient "github.com/blocktop/go-rpc-client"
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
