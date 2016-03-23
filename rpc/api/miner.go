// Copyright 2015 The shift Authors
// This file is part of the shift library.
//
// The shift library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The shift library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the shift library. If not, see <http://www.gnu.org/licenses/>.

package api

import (
	"github.com/shiftcurrency/ethash"
	"github.com/shiftcurrency/shift/common"
	"github.com/shiftcurrency/shift/shf"
	"github.com/shiftcurrency/shift/rpc/codec"
	"github.com/shiftcurrency/shift/rpc/shared"
)

const (
	MinerApiVersion = "1.0"
)

var (
	// mapping between methods and handlers
	MinerMapping = map[string]minerhandler{
		"miner_hashrate":     (*minerApi).Hashrate,
		"miner_makeDAG":      (*minerApi).MakeDAG,
		"miner_setExtra":     (*minerApi).SetExtra,
		"miner_setNrgPrice":  (*minerApi).SetNrgPrice,
		"miner_setEtherbase": (*minerApi).SetEtherbase,
		"miner_startAutoDAG": (*minerApi).StartAutoDAG,
		"miner_start":        (*minerApi).StartMiner,
		"miner_stopAutoDAG":  (*minerApi).StopAutoDAG,
		"miner_stop":         (*minerApi).StopMiner,
	}
)

// miner callback handler
type minerhandler func(*minerApi, *shared.Request) (interface{}, error)

// miner api provider
type minerApi struct {
	shift *shf.Shift
	methods  map[string]minerhandler
	codec    codec.ApiCoder
}

// create a new miner api instance
func NewMinerApi(shift *shf.Shift, coder codec.Codec) *minerApi {
	return &minerApi{
		shift: shift,
		methods:  MinerMapping,
		codec:    coder.New(nil),
	}
}

// Execute given request
func (self *minerApi) Execute(req *shared.Request) (interface{}, error) {
	if callback, ok := self.methods[req.Method]; ok {
		return callback(self, req)
	}

	return nil, &shared.NotImplementedError{req.Method}
}

// collection with supported methods
func (self *minerApi) Methods() []string {
	methods := make([]string, len(self.methods))
	i := 0
	for k := range self.methods {
		methods[i] = k
		i++
	}
	return methods
}

func (self *minerApi) Name() string {
	return shared.MinerApiName
}

func (self *minerApi) ApiVersion() string {
	return MinerApiVersion
}

func (self *minerApi) StartMiner(req *shared.Request) (interface{}, error) {
	args := new(StartMinerArgs)
	if err := self.codec.Decode(req.Params, &args); err != nil {
		return nil, err
	}
	if args.Threads == -1 { // (not specified by user, use default)
		args.Threads = self.shift.MinerThreads
	}

	self.shift.StartAutoDAG()
	err := self.shift.StartMining(args.Threads, "")
	if err == nil {
		return true, nil
	}

	return false, err
}

func (self *minerApi) StopMiner(req *shared.Request) (interface{}, error) {
	self.shift.StopMining()
	return true, nil
}

func (self *minerApi) Hashrate(req *shared.Request) (interface{}, error) {
	return self.shift.Miner().HashRate(), nil
}

func (self *minerApi) SetExtra(req *shared.Request) (interface{}, error) {
	args := new(SetExtraArgs)
	if err := self.codec.Decode(req.Params, &args); err != nil {
		return nil, err
	}

	if err := self.shift.Miner().SetExtra([]byte(args.Data)); err != nil {
		return false, err
	}

	return true, nil
}

func (self *minerApi) SetNrgPrice(req *shared.Request) (interface{}, error) {
	args := new(NrgPriceArgs)
	if err := self.codec.Decode(req.Params, &args); err != nil {
		return false, err
	}

	self.shift.Miner().SetNrgPrice(common.String2Big(args.Price))
	return true, nil
}

func (self *minerApi) SetEtherbase(req *shared.Request) (interface{}, error) {
	args := new(SetEtherbaseArgs)
	if err := self.codec.Decode(req.Params, &args); err != nil {
		return false, err
	}
	self.shift.SetEtherbase(args.Etherbase)
	return nil, nil
}

func (self *minerApi) StartAutoDAG(req *shared.Request) (interface{}, error) {
	self.shift.StartAutoDAG()
	return true, nil
}

func (self *minerApi) StopAutoDAG(req *shared.Request) (interface{}, error) {
	self.shift.StopAutoDAG()
	return true, nil
}

func (self *minerApi) MakeDAG(req *shared.Request) (interface{}, error) {
	args := new(MakeDAGArgs)
	if err := self.codec.Decode(req.Params, &args); err != nil {
		return nil, err
	}

	if args.BlockNumber < 0 {
		return false, shared.NewValidationError("BlockNumber", "BlockNumber must be positive")
	}

	err := ethash.MakeDAG(uint64(args.BlockNumber), "")
	if err == nil {
		return true, nil
	}
	return false, err
}