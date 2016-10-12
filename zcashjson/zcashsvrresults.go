// Copyright (c) 2016 arithmetric
// Based on btcd by the btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package zcashjson

// ZOperationStatusError models the error data in ZGetOperationStatusResult.
type ZOperationStatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ZGetOperationStatusResult models the data from the z_getoperationresult and
// z_getoperationstatus commands.
type ZGetOperationStatusResult struct {
	Id            string                `json:"id"`
	Status        string                `json:"status"`
	CreationTime  int                   `json:"creation_time"`
	Result        map[string]string     `json:"result"`
	Error         ZOperationStatusError `json:"error"`
	ExecutionSecs float64               `json:"execution_secs"`
}

// ZGetTotalBalanceResult models the data from the z_gettotalbalance command.
type ZGetTotalBalanceResult struct {
	Transparent string `json:"transparent"`
	Private     string `json:"private"`
	Total       string `json:"total"`
}

// ZListReceivedByAddressResult models the data from the z_listreceivedbyaddress
// command.
type ZListReceivedByAddressResult struct {
	TxID   string  `json:"txid"`
	Amount float64 `json:"amount"`
	Memo   string  `json:"memo"`
}
