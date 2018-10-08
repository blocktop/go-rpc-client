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

func GetMetrics(format string) (*GetMetricsReply, error) {
	params := []GetMetricsArgs{GetMetricsArgs{format}}
	req := &rpcclient.RPCRequest{
		Method: "Kernel.GetMetrics",
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
