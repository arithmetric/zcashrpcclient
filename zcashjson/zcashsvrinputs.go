// Copyright (c) 2016 arithmetric
// Based on btcd by the btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package zcashjson

// ZSendManyEntry models the inputs for the z_sendmany command.
type ZSendManyEntry struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
	Memo    *string `json:"memo"`
}
