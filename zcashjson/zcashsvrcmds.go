// Copyright (c) 2016 arithmetric
// Based on btcd by the btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// Zcash.

package zcashjson

import(
  "github.com/btcsuite/btcd/btcjson"
)

// ZGetBalanceCmd defines the z_getbalance JSON-RPC command.
type ZGetBalanceCmd struct {
	Address *string
	MinConf *int `jsonrpcdefault:"1"`
}

// NewZGetBalanceCmd returns a new instance which can be used to issue a
// z_getbalance JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZGetBalanceCmd(address *string, minConf *int) *ZGetBalanceCmd {
	return &ZGetBalanceCmd{
		Address: address,
		MinConf: minConf,
	}
}

// ZGetNewAddressCmd defines the z_getnewaddress JSON-RPC command.
type ZGetNewAddressCmd struct {
	MinConf *int `jsonrpcdefault:"1"`
}

// NewGetNewAddressCmd returns a new instance which can be used to issue a
// getnewaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional. Passing nil
// for optional parameters will use the default value.
func NewZGetNewAddressCmd() *ZGetNewAddressCmd {
	return &ZGetNewAddressCmd{}
}

// ZGetTotalBalanceCmd defines the z_gettotalbalance JSON-RPC command.
type ZGetTotalBalanceCmd struct {
	MinConf *int `jsonrpcdefault:"1"`
}

// NewZGetTotalBalanceCmd returns a new instance which can be used to issue a
// z_gettotalbalance JSON-RPC command.
//
// The parameters which are pointers indicate they are optional. Passing nil
// for optional parameters will use the default value.
func NewZGetTotalBalanceCmd() *ZGetTotalBalanceCmd {
	return &ZGetTotalBalanceCmd{}
}

// ZListReceivedByAddressCmd defines the z_listreceivedbyaddress JSON-RPC command.
type ZListReceivedByAddressCmd struct {
	Address string
	MinConf *int `jsonrpcdefault:"1"`
}

// NewZListReceivedByAddressCmd returns a new instance which can be used to issue
// a z_listreceivedbyaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZListReceivedByAddressCmd(address string, minConf *int) *ZListReceivedByAddressCmd {
	return &ZListReceivedByAddressCmd{
		Address: address,
		MinConf: minConf,
	}
}

// ImportKeyCmd defines the z_importkey JSON-RPC command.
type ZImportKeyCmd struct {
	ZKey   string
	Rescan *bool `jsonrpcdefault:"true"`
}

// NewImportKeyCmd returns a new instance which can be used to issue a
// z_importkey JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZImportKeyCmd(zkey string, rescan *bool) *ZImportKeyCmd {
	return &ZImportKeyCmd{
		ZKey:   zkey,
		Rescan: rescan,
	}
}

// ZImportWalletCmd defines the z_importwallet JSON-RPC command.
type ZImportWalletCmd struct {
	Filename string
}

// NewZImportWalletCmd returns a new instance which can be used to issue a
// z_importwallet JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZImportWalletCmd(filename string) *ZImportWalletCmd {
	return &ZImportWalletCmd{
		Filename: filename,
	}
}

// ZExportKeyCmd defines the z_exportkey JSON-RPC command.
type ZExportKeyCmd struct {
	Address string
}

// NewZExportKeyCmd returns a new instance which can be used to issue a
// z_exportkey JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZExportKeyCmd(address string) *ZExportKeyCmd {
	return &ZExportKeyCmd{
		Address: address,
	}
}

// ZExportWalletCmd defines the z_exportwallet JSON-RPC command.
type ZExportWalletCmd struct {
	Filename string
}

// NewZExportWalletCmd returns a new instance which can be used to issue a
// z_exportwallet JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZExportWalletCmd(filename string) *ZExportWalletCmd {
	return &ZExportWalletCmd{
		Filename: filename,
	}
}

// ZListAccountsCmd defines the z_listaccounts JSON-RPC command.
type ZListAddressesCmd struct {
	MinConf *int `jsonrpcdefault:"1"`
}

// NewZListAccountsCmd returns a new instance which can be used to issue a
// z_listaccounts JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZListAddressesCmd(minConf *int) *ZListAddressesCmd {
	return &ZListAddressesCmd{
		MinConf: minConf,
	}
}

// ZListOperationIdsCmd defines the z_listoperationids JSON-RPC command.
type ZListOperationIdsCmd struct {
	MinConf *int `jsonrpcdefault:"1"`
}

// ZListOperationIdsCmd returns a new instance which can be used to issue a
// z_listoperationids JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZListOperationIdsCmd(minConf *int) *ZListOperationIdsCmd {
	return &ZListOperationIdsCmd{
		MinConf: minConf,
	}
}

// ZGetOperationResultCmd defines the z_getoperationresult JSON-RPC command.
type ZGetOperationResultCmd struct {
	MinConf *int `jsonrpcdefault:"1"`
}

// ZGetOperationResultCmd returns a new instance which can be used to issue a
// z_getoperationresult JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZGetOperationResultCmd(minConf *int) *ZGetOperationResultCmd {
	return &ZGetOperationResultCmd{
		MinConf: minConf,
	}
}

// ZGetOperationStatusCmd defines the z_getoperationstatus JSON-RPC command.
type ZGetOperationStatusCmd struct {
	MinConf *int `jsonrpcdefault:"1"`
}

// ZGetOperationStatusCmd returns a new instance which can be used to issue a
// z_getoperationstatus JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZGetOperationStatusCmd(minConf *int) *ZGetOperationStatusCmd {
	return &ZGetOperationStatusCmd{
		MinConf: minConf,
	}
}

// ZSendManyCmd defines the z_sendmany JSON-RPC command.
type ZSendManyCmd struct {
	FromAccount string
	Amounts     []ZSendManyEntry `jsonrpcusage:"{\"address\":address,\"amount\":amount,...}"`
	MinConf     *int             `jsonrpcdefault:"1"`
}

// NewZSendManyCmd returns a new instance which can be used to issue a z_sendmany
// JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewZSendManyCmd(fromAccount string, amounts []ZSendManyEntry, minConf *int) *ZSendManyCmd {
	return &ZSendManyCmd{
		FromAccount: fromAccount,
		Amounts:     amounts,
		MinConf:     minConf,
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	flags := btcjson.UFWalletOnly

	btcjson.MustRegisterCmd("z_exportkey", (*ZExportKeyCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_exportwallet", (*ZExportWalletCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_getbalance", (*ZGetBalanceCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_getoperationresult", (*ZGetOperationResultCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_getoperationstatus", (*ZGetOperationStatusCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_getnewaddress", (*ZGetNewAddressCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_gettotalbalance", (*ZGetTotalBalanceCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_importkey", (*ZImportKeyCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_importwallet", (*ZImportWalletCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_listaddresses", (*ZListAddressesCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_listoperationids", (*ZListOperationIdsCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_listreceivedbyaddress", (*ZListReceivedByAddressCmd)(nil), flags)
	btcjson.MustRegisterCmd("z_sendmany", (*ZSendManyCmd)(nil), flags)
}
