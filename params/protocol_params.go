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

// DO NOT EDIT!!!
// AUTOGENERATED FROM generators/defaults.go

package params

import "math/big"

var (
	MaximumExtraDataSize   = big.NewInt(32)     // Maximum size extra data may be after Genesis.
	ExpByteNrg             = big.NewInt(10)     // Times ceil(log256(exponent)) for the EXP instruction.
	SloadNrg               = big.NewInt(50)     // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransferNrg   = big.NewInt(9000)   // Paid for CALL when the value transfer is non-zero.
	CallNewAccountNrg      = big.NewInt(25000)  // Paid for CALL when the destination address didn't exist prior.
	TxNrg                  = big.NewInt(21000)  // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxNrgContractCreation  = big.NewInt(53000)  // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroNrg          = big.NewInt(4)      // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	DifficultyBoundDivisor = big.NewInt(256)    // The bound divisor of the difficulty, used in the update calculations.
	QuadCoeffDiv           = big.NewInt(512)    // Divisor for the quadratic particle of the memory cost equation.
	GenesisDifficulty      = big.NewInt(1500) // Difficulty of the Genesis block.
	DurationLimit          = big.NewInt(60)     // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	SstoreSetNrg           = big.NewInt(20000)  // Once per SLOAD operation.
	LogDataNrg             = big.NewInt(8)      // Per byte in a LOG* operation's data.
	CallStipend            = big.NewInt(2300)   // Free nrg given at beginning of call.
	EcrecoverNrg           = big.NewInt(3000)   //
	Sha256WordNrg          = big.NewInt(12)     //
	MinNrgLimit            = big.NewInt(5000)    // Minimum the nrg limit may ever be.
	GenesisNrgLimit        = big.NewInt(31415926) // Nrg limit of the Genesis block.
	Sha3Nrg              = big.NewInt(30)     // Once per SHA3 operation.
	Sha256Nrg            = big.NewInt(60)     //
	IdentityWordNrg      = big.NewInt(3)      //
	Sha3WordNrg          = big.NewInt(6)      // Once per word of the SHA3 operation's data.
	SstoreResetNrg       = big.NewInt(5000)   // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearNrg       = big.NewInt(5000)   // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundNrg      = big.NewInt(15000)  // Once per SSTORE operation if the zeroness changes to zero.
	JumpdestNrg          = big.NewInt(1)      // Refunded nrg, once per SSTORE operation if the zeroness changes to zero.
	IdentityNrg          = big.NewInt(15)     //
	NrgLimitBoundDivisor = big.NewInt(1024)   // The bound divisor of the nrg limit, used in update calculations.
	EpochDuration        = big.NewInt(30000)  // Duration between proof-of-work epochs.
	CallNrg              = big.NewInt(40)     // Once per CALL operation & message call transaction.
	CreateDataNrg        = big.NewInt(200)    //
	Ripemd160Nrg         = big.NewInt(600)    //
	Ripemd160WordNrg     = big.NewInt(120)    //
	MinimumDifficulty    = big.NewInt(131072) // The minimum that the difficulty may ever be.
	CallCreateDepth      = big.NewInt(1024)   // Maximum depth of call/create stack.
	ExpNrg               = big.NewInt(10)     // Once per EXP instruction.
	LogNrg               = big.NewInt(375)    // Per LOG* operation.
	CopyNrg              = big.NewInt(3)      //
	StackLimit           = big.NewInt(1024)   // Maximum size of VM stack allowed.
	TierStepNrg          = big.NewInt(0)      // Once per operation, for a selection of them.
	LogTopicNrg          = big.NewInt(375)    // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicNrg, LOG4 incurs 4 * c_txLogTopicNrg.
	CreateNrg            = big.NewInt(32000)  // Once per CREATE operation & contract-creation transaction.
	SuicideRefundNrg     = big.NewInt(24000)  // Refunded following a suicide operation.
	MemoryNrg            = big.NewInt(3)      // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	TxDataNonZeroNrg     = big.NewInt(68)     // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.
)