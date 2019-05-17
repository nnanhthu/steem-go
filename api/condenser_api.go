package api

import (
	"encoding/json"
	"steem-go/types"

	"steem-go/transports"
	_ "steem-go/types"
)

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call("condenser_api", "get_config", transports.EmptyParams, &resp)
	return &resp, err
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := api.call("condenser_api", "get_dynamic_global_properties", transports.EmptyParams, &resp)
	return &resp, err
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	err := api.call("condenser_api", "get_block", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	err := api.call("condenser_api", "get_block_header", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

// Set callback to invoke as soon as a new block is applied
func (api *API) SetBlockAppliedCallback(notice func(header *BlockHeader, error error)) (err error) {
	err = api.setCallback("condenser_api", "set_block_applied_callback", func(raw json.RawMessage) {
		var header []BlockHeader
		if err := json.Unmarshal(raw, &header); err != nil {
			notice(nil, err)
		}
		for _, b := range header {
			notice(&b, nil)
		}
	})
	return
}

//BroadcastTransaction api request broadcast_transaction
func (api *API) BroadcastTransaction(tx *types.Transaction) error {
	return api.call("condenser_api", "broadcast_transaction", []interface{}{tx}, nil)
}

//BroadcastTransactionSynchronous api request broadcast_transaction_synchronous
func (api *API) BroadcastTransactionSynchronous(tx *types.Transaction) (*BroadcastResponse, error) {
	var resp BroadcastResponse
	err := api.call("condenser_api", "broadcast_transaction_synchronous", []interface{}{tx}, &resp)
	return &resp, err
}
