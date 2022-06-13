// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package birden

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BirdenTokenInfoStruct is an auto generated low-level Go binding around an user-defined struct.
type BirdenTokenInfoStruct struct {
	TokenId  *big.Int
	URI      string
	Creator  common.Address
	RentedAt *big.Int
	RentedBy common.Address
	Balance  *big.Int
}

// BirdenMetaData contains all meta data concerning the Birden contract.
var BirdenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_comission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"rentPriceCelo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"tokensInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rentedAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rentedBy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBirden.TokenInfoStruct[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"rent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rent_token_address\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BirdenABI is the input ABI used to generate the binding from.
// Deprecated: Use BirdenMetaData.ABI instead.
var BirdenABI = BirdenMetaData.ABI

// Birden is an auto generated Go binding around an Ethereum contract.
type Birden struct {
	BirdenCaller     // Read-only binding to the contract
	BirdenTransactor // Write-only binding to the contract
	BirdenFilterer   // Log filterer for contract events
}

// BirdenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BirdenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BirdenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BirdenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BirdenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BirdenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BirdenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BirdenSession struct {
	Contract     *Birden           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BirdenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BirdenCallerSession struct {
	Contract *BirdenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BirdenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BirdenTransactorSession struct {
	Contract     *BirdenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BirdenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BirdenRaw struct {
	Contract *Birden // Generic contract binding to access the raw methods on
}

// BirdenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BirdenCallerRaw struct {
	Contract *BirdenCaller // Generic read-only contract binding to access the raw methods on
}

// BirdenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BirdenTransactorRaw struct {
	Contract *BirdenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBirden creates a new instance of Birden, bound to a specific deployed contract.
func NewBirden(address common.Address, backend bind.ContractBackend) (*Birden, error) {
	contract, err := bindBirden(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Birden{BirdenCaller: BirdenCaller{contract: contract}, BirdenTransactor: BirdenTransactor{contract: contract}, BirdenFilterer: BirdenFilterer{contract: contract}}, nil
}

// NewBirdenCaller creates a new read-only instance of Birden, bound to a specific deployed contract.
func NewBirdenCaller(address common.Address, caller bind.ContractCaller) (*BirdenCaller, error) {
	contract, err := bindBirden(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BirdenCaller{contract: contract}, nil
}

// NewBirdenTransactor creates a new write-only instance of Birden, bound to a specific deployed contract.
func NewBirdenTransactor(address common.Address, transactor bind.ContractTransactor) (*BirdenTransactor, error) {
	contract, err := bindBirden(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BirdenTransactor{contract: contract}, nil
}

// NewBirdenFilterer creates a new log filterer instance of Birden, bound to a specific deployed contract.
func NewBirdenFilterer(address common.Address, filterer bind.ContractFilterer) (*BirdenFilterer, error) {
	contract, err := bindBirden(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BirdenFilterer{contract: contract}, nil
}

// bindBirden binds a generic wrapper to an already deployed contract.
func bindBirden(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BirdenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Birden *BirdenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Birden.Contract.BirdenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Birden *BirdenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Birden.Contract.BirdenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Birden *BirdenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Birden.Contract.BirdenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Birden *BirdenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Birden.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Birden *BirdenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Birden.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Birden *BirdenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Birden.Contract.contract.Transact(opts, method, params...)
}

// Comission is a free data retrieval call binding the contract method 0x06ef93ad.
//
// Solidity: function _comission() view returns(uint256)
func (_Birden *BirdenCaller) Comission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "_comission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Comission is a free data retrieval call binding the contract method 0x06ef93ad.
//
// Solidity: function _comission() view returns(uint256)
func (_Birden *BirdenSession) Comission() (*big.Int, error) {
	return _Birden.Contract.Comission(&_Birden.CallOpts)
}

// Comission is a free data retrieval call binding the contract method 0x06ef93ad.
//
// Solidity: function _comission() view returns(uint256)
func (_Birden *BirdenCallerSession) Comission() (*big.Int, error) {
	return _Birden.Contract.Comission(&_Birden.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Birden.Contract.Allowance(&_Birden.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Birden.Contract.Allowance(&_Birden.CallOpts, owner, spender)
}

// Allowance0 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenCaller) Allowance0(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "allowance0", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance0 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenSession) Allowance0(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Birden.Contract.Allowance0(&_Birden.CallOpts, owner, spender)
}

// Allowance0 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenCallerSession) Allowance0(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Birden.Contract.Allowance0(&_Birden.CallOpts, owner, spender)
}

// Allowance1 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenCaller) Allowance1(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "allowance1", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance1 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenSession) Allowance1(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Birden.Contract.Allowance1(&_Birden.CallOpts, owner, spender)
}

// Allowance1 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Birden *BirdenCallerSession) Allowance1(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Birden.Contract.Allowance1(&_Birden.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Birden *BirdenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Birden *BirdenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf(&_Birden.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Birden *BirdenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf(&_Birden.CallOpts, owner)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenCaller) BalanceOf0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenSession) BalanceOf0(account common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf0(&_Birden.CallOpts, account)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenCallerSession) BalanceOf0(account common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf0(&_Birden.CallOpts, account)
}

// BalanceOf1 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Birden *BirdenCaller) BalanceOf1(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf1", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf1 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Birden *BirdenSession) BalanceOf1(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf1(&_Birden.CallOpts, owner)
}

// BalanceOf1 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Birden *BirdenCallerSession) BalanceOf1(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf1(&_Birden.CallOpts, owner)
}

// BalanceOf2 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenCaller) BalanceOf2(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf2", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf2 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenSession) BalanceOf2(account common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf2(&_Birden.CallOpts, account)
}

// BalanceOf2 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenCallerSession) BalanceOf2(account common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf2(&_Birden.CallOpts, account)
}

// BalanceOf3 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenCaller) BalanceOf3(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf3", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf3 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenSession) BalanceOf3(account common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf3(&_Birden.CallOpts, account)
}

// BalanceOf3 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Birden *BirdenCallerSession) BalanceOf3(account common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf3(&_Birden.CallOpts, account)
}

// BalanceOf4 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Birden *BirdenCaller) BalanceOf4(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf4", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf4 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Birden *BirdenSession) BalanceOf4(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf4(&_Birden.CallOpts, owner)
}

// BalanceOf4 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Birden *BirdenCallerSession) BalanceOf4(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf4(&_Birden.CallOpts, owner)
}

// BalanceOf5 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Birden *BirdenCaller) BalanceOf5(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "balanceOf5", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf5 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Birden *BirdenSession) BalanceOf5(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf5(&_Birden.CallOpts, owner)
}

// BalanceOf5 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Birden *BirdenCallerSession) BalanceOf5(owner common.Address) (*big.Int, error) {
	return _Birden.Contract.BalanceOf5(&_Birden.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Birden *BirdenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Birden *BirdenSession) Decimals() (uint8, error) {
	return _Birden.Contract.Decimals(&_Birden.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Birden *BirdenCallerSession) Decimals() (uint8, error) {
	return _Birden.Contract.Decimals(&_Birden.CallOpts)
}

// Decimals0 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Birden *BirdenCaller) Decimals0(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "decimals0")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals0 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Birden *BirdenSession) Decimals0() (uint8, error) {
	return _Birden.Contract.Decimals0(&_Birden.CallOpts)
}

// Decimals0 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Birden *BirdenCallerSession) Decimals0() (uint8, error) {
	return _Birden.Contract.Decimals0(&_Birden.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Birden *BirdenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Birden *BirdenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved(&_Birden.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Birden *BirdenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved(&_Birden.CallOpts, tokenId)
}

// GetApproved0 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Birden *BirdenCaller) GetApproved0(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "getApproved0", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved0 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Birden *BirdenSession) GetApproved0(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved0(&_Birden.CallOpts, tokenId)
}

// GetApproved0 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Birden *BirdenCallerSession) GetApproved0(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved0(&_Birden.CallOpts, tokenId)
}

// GetApproved1 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Birden *BirdenCaller) GetApproved1(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "getApproved1", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved1 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Birden *BirdenSession) GetApproved1(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved1(&_Birden.CallOpts, tokenId)
}

// GetApproved1 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Birden *BirdenCallerSession) GetApproved1(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved1(&_Birden.CallOpts, tokenId)
}

// GetApproved2 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Birden *BirdenCaller) GetApproved2(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "getApproved2", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved2 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Birden *BirdenSession) GetApproved2(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved2(&_Birden.CallOpts, tokenId)
}

// GetApproved2 is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Birden *BirdenCallerSession) GetApproved2(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.GetApproved2(&_Birden.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll0 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCaller) IsApprovedForAll0(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "isApprovedForAll0", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll0 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenSession) IsApprovedForAll0(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll0(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll0 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCallerSession) IsApprovedForAll0(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll0(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll1 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCaller) IsApprovedForAll1(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "isApprovedForAll1", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll1 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenSession) IsApprovedForAll1(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll1(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll1 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCallerSession) IsApprovedForAll1(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll1(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll2 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCaller) IsApprovedForAll2(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "isApprovedForAll2", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll2 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenSession) IsApprovedForAll2(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll2(&_Birden.CallOpts, owner, operator)
}

// IsApprovedForAll2 is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Birden *BirdenCallerSession) IsApprovedForAll2(owner common.Address, operator common.Address) (bool, error) {
	return _Birden.Contract.IsApprovedForAll2(&_Birden.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenSession) Name() (string, error) {
	return _Birden.Contract.Name(&_Birden.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCallerSession) Name() (string, error) {
	return _Birden.Contract.Name(&_Birden.CallOpts)
}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCaller) Name0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "name0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenSession) Name0() (string, error) {
	return _Birden.Contract.Name0(&_Birden.CallOpts)
}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCallerSession) Name0() (string, error) {
	return _Birden.Contract.Name0(&_Birden.CallOpts)
}

// Name1 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCaller) Name1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "name1")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name1 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenSession) Name1() (string, error) {
	return _Birden.Contract.Name1(&_Birden.CallOpts)
}

// Name1 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCallerSession) Name1() (string, error) {
	return _Birden.Contract.Name1(&_Birden.CallOpts)
}

// Name2 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCaller) Name2(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "name2")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name2 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenSession) Name2() (string, error) {
	return _Birden.Contract.Name2(&_Birden.CallOpts)
}

// Name2 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCallerSession) Name2() (string, error) {
	return _Birden.Contract.Name2(&_Birden.CallOpts)
}

// Name3 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCaller) Name3(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "name3")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name3 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenSession) Name3() (string, error) {
	return _Birden.Contract.Name3(&_Birden.CallOpts)
}

// Name3 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Birden *BirdenCallerSession) Name3() (string, error) {
	return _Birden.Contract.Name3(&_Birden.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Birden *BirdenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Birden *BirdenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf(&_Birden.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Birden *BirdenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf(&_Birden.CallOpts, tokenId)
}

// OwnerOf0 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Birden *BirdenCaller) OwnerOf0(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "ownerOf0", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf0 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Birden *BirdenSession) OwnerOf0(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf0(&_Birden.CallOpts, tokenId)
}

// OwnerOf0 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Birden *BirdenCallerSession) OwnerOf0(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf0(&_Birden.CallOpts, tokenId)
}

// OwnerOf1 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Birden *BirdenCaller) OwnerOf1(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "ownerOf1", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf1 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Birden *BirdenSession) OwnerOf1(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf1(&_Birden.CallOpts, tokenId)
}

// OwnerOf1 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Birden *BirdenCallerSession) OwnerOf1(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf1(&_Birden.CallOpts, tokenId)
}

// OwnerOf2 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Birden *BirdenCaller) OwnerOf2(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "ownerOf2", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf2 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Birden *BirdenSession) OwnerOf2(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf2(&_Birden.CallOpts, tokenId)
}

// OwnerOf2 is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Birden *BirdenCallerSession) OwnerOf2(tokenId *big.Int) (common.Address, error) {
	return _Birden.Contract.OwnerOf2(&_Birden.CallOpts, tokenId)
}

// RentPriceCelo is a free data retrieval call binding the contract method 0xa69bd72c.
//
// Solidity: function rentPriceCelo() view returns(uint256)
func (_Birden *BirdenCaller) RentPriceCelo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "rentPriceCelo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RentPriceCelo is a free data retrieval call binding the contract method 0xa69bd72c.
//
// Solidity: function rentPriceCelo() view returns(uint256)
func (_Birden *BirdenSession) RentPriceCelo() (*big.Int, error) {
	return _Birden.Contract.RentPriceCelo(&_Birden.CallOpts)
}

// RentPriceCelo is a free data retrieval call binding the contract method 0xa69bd72c.
//
// Solidity: function rentPriceCelo() view returns(uint256)
func (_Birden *BirdenCallerSession) RentPriceCelo() (*big.Int, error) {
	return _Birden.Contract.RentPriceCelo(&_Birden.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface0 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCaller) SupportsInterface0(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "supportsInterface0", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface0 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenSession) SupportsInterface0(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface0(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface0 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCallerSession) SupportsInterface0(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface0(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface1 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCaller) SupportsInterface1(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "supportsInterface1", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface1 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenSession) SupportsInterface1(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface1(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface1 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCallerSession) SupportsInterface1(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface1(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface2 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCaller) SupportsInterface2(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "supportsInterface2", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface2 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenSession) SupportsInterface2(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface2(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface2 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCallerSession) SupportsInterface2(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface2(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface3 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCaller) SupportsInterface3(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "supportsInterface3", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface3 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenSession) SupportsInterface3(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface3(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface3 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCallerSession) SupportsInterface3(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface3(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface4 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCaller) SupportsInterface4(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "supportsInterface4", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface4 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenSession) SupportsInterface4(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface4(&_Birden.CallOpts, interfaceId)
}

// SupportsInterface4 is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Birden *BirdenCallerSession) SupportsInterface4(interfaceId [4]byte) (bool, error) {
	return _Birden.Contract.SupportsInterface4(&_Birden.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenSession) Symbol() (string, error) {
	return _Birden.Contract.Symbol(&_Birden.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCallerSession) Symbol() (string, error) {
	return _Birden.Contract.Symbol(&_Birden.CallOpts)
}

// Symbol0 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCaller) Symbol0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "symbol0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol0 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenSession) Symbol0() (string, error) {
	return _Birden.Contract.Symbol0(&_Birden.CallOpts)
}

// Symbol0 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCallerSession) Symbol0() (string, error) {
	return _Birden.Contract.Symbol0(&_Birden.CallOpts)
}

// Symbol1 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCaller) Symbol1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "symbol1")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol1 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenSession) Symbol1() (string, error) {
	return _Birden.Contract.Symbol1(&_Birden.CallOpts)
}

// Symbol1 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCallerSession) Symbol1() (string, error) {
	return _Birden.Contract.Symbol1(&_Birden.CallOpts)
}

// Symbol2 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCaller) Symbol2(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "symbol2")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol2 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenSession) Symbol2() (string, error) {
	return _Birden.Contract.Symbol2(&_Birden.CallOpts)
}

// Symbol2 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCallerSession) Symbol2() (string, error) {
	return _Birden.Contract.Symbol2(&_Birden.CallOpts)
}

// Symbol3 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCaller) Symbol3(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "symbol3")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol3 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenSession) Symbol3() (string, error) {
	return _Birden.Contract.Symbol3(&_Birden.CallOpts)
}

// Symbol3 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Birden *BirdenCallerSession) Symbol3() (string, error) {
	return _Birden.Contract.Symbol3(&_Birden.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Birden.Contract.TokenURI(&_Birden.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Birden.Contract.TokenURI(&_Birden.CallOpts, tokenId)
}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenCaller) TokenURI0(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "tokenURI0", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenSession) TokenURI0(tokenId *big.Int) (string, error) {
	return _Birden.Contract.TokenURI0(&_Birden.CallOpts, tokenId)
}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenCallerSession) TokenURI0(tokenId *big.Int) (string, error) {
	return _Birden.Contract.TokenURI0(&_Birden.CallOpts, tokenId)
}

// TokenURI1 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenCaller) TokenURI1(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "tokenURI1", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI1 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenSession) TokenURI1(tokenId *big.Int) (string, error) {
	return _Birden.Contract.TokenURI1(&_Birden.CallOpts, tokenId)
}

// TokenURI1 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Birden *BirdenCallerSession) TokenURI1(tokenId *big.Int) (string, error) {
	return _Birden.Contract.TokenURI1(&_Birden.CallOpts, tokenId)
}

// TokensInfo is a free data retrieval call binding the contract method 0xe79c28fe.
//
// Solidity: function tokensInfo() view returns((uint256,string,address,uint256,address,uint256)[])
func (_Birden *BirdenCaller) TokensInfo(opts *bind.CallOpts) ([]BirdenTokenInfoStruct, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "tokensInfo")

	if err != nil {
		return *new([]BirdenTokenInfoStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]BirdenTokenInfoStruct)).(*[]BirdenTokenInfoStruct)

	return out0, err

}

// TokensInfo is a free data retrieval call binding the contract method 0xe79c28fe.
//
// Solidity: function tokensInfo() view returns((uint256,string,address,uint256,address,uint256)[])
func (_Birden *BirdenSession) TokensInfo() ([]BirdenTokenInfoStruct, error) {
	return _Birden.Contract.TokensInfo(&_Birden.CallOpts)
}

// TokensInfo is a free data retrieval call binding the contract method 0xe79c28fe.
//
// Solidity: function tokensInfo() view returns((uint256,string,address,uint256,address,uint256)[])
func (_Birden *BirdenCallerSession) TokensInfo() ([]BirdenTokenInfoStruct, error) {
	return _Birden.Contract.TokensInfo(&_Birden.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenSession) TotalSupply() (*big.Int, error) {
	return _Birden.Contract.TotalSupply(&_Birden.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenCallerSession) TotalSupply() (*big.Int, error) {
	return _Birden.Contract.TotalSupply(&_Birden.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenCaller) TotalSupply0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "totalSupply0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenSession) TotalSupply0() (*big.Int, error) {
	return _Birden.Contract.TotalSupply0(&_Birden.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenCallerSession) TotalSupply0() (*big.Int, error) {
	return _Birden.Contract.TotalSupply0(&_Birden.CallOpts)
}

// TotalSupply1 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenCaller) TotalSupply1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Birden.contract.Call(opts, &out, "totalSupply1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply1 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenSession) TotalSupply1() (*big.Int, error) {
	return _Birden.Contract.TotalSupply1(&_Birden.CallOpts)
}

// TotalSupply1 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Birden *BirdenCallerSession) TotalSupply1() (*big.Int, error) {
	return _Birden.Contract.TotalSupply1(&_Birden.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve(&_Birden.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve(&_Birden.TransactOpts, to, tokenId)
}

// Approve0 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) Approve0(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve0", spender, amount)
}

// Approve0 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenSession) Approve0(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve0(&_Birden.TransactOpts, spender, amount)
}

// Approve0 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) Approve0(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve0(&_Birden.TransactOpts, spender, amount)
}

// Approve1 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) Approve1(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve1", to, tokenId)
}

// Approve1 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) Approve1(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve1(&_Birden.TransactOpts, to, tokenId)
}

// Approve1 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) Approve1(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve1(&_Birden.TransactOpts, to, tokenId)
}

// Approve2 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) Approve2(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve2", spender, amount)
}

// Approve2 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenSession) Approve2(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve2(&_Birden.TransactOpts, spender, amount)
}

// Approve2 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) Approve2(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve2(&_Birden.TransactOpts, spender, amount)
}

// Approve3 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) Approve3(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve3", spender, amount)
}

// Approve3 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenSession) Approve3(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve3(&_Birden.TransactOpts, spender, amount)
}

// Approve3 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) Approve3(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve3(&_Birden.TransactOpts, spender, amount)
}

// Approve4 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) Approve4(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve4", to, tokenId)
}

// Approve4 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) Approve4(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve4(&_Birden.TransactOpts, to, tokenId)
}

// Approve4 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) Approve4(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve4(&_Birden.TransactOpts, to, tokenId)
}

// Approve5 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) Approve5(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "approve5", to, tokenId)
}

// Approve5 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) Approve5(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve5(&_Birden.TransactOpts, to, tokenId)
}

// Approve5 is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) Approve5(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Approve5(&_Birden.TransactOpts, to, tokenId)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Birden *BirdenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Birden *BirdenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.DecreaseAllowance(&_Birden.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Birden *BirdenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.DecreaseAllowance(&_Birden.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Birden *BirdenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Birden *BirdenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.IncreaseAllowance(&_Birden.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Birden *BirdenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.IncreaseAllowance(&_Birden.TransactOpts, spender, addedValue)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Birden *BirdenTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Birden *BirdenSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.Contract.OnERC721Received(&_Birden.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Birden *BirdenTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.Contract.OnERC721Received(&_Birden.TransactOpts, operator, from, tokenId, data)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 tokenId) returns()
func (_Birden *BirdenTransactor) Redeem(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "redeem", tokenId)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 tokenId) returns()
func (_Birden *BirdenSession) Redeem(tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Redeem(&_Birden.TransactOpts, tokenId)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) Redeem(tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Redeem(&_Birden.TransactOpts, tokenId)
}

// Rent is a paid mutator transaction binding the contract method 0x7456be7d.
//
// Solidity: function rent(uint256 tokenId) returns()
func (_Birden *BirdenTransactor) Rent(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "rent", tokenId)
}

// Rent is a paid mutator transaction binding the contract method 0x7456be7d.
//
// Solidity: function rent(uint256 tokenId) returns()
func (_Birden *BirdenSession) Rent(tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Rent(&_Birden.TransactOpts, tokenId)
}

// Rent is a paid mutator transaction binding the contract method 0x7456be7d.
//
// Solidity: function rent(uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) Rent(tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Rent(&_Birden.TransactOpts, tokenId)
}

// SafeMint is a paid mutator transaction binding the contract method 0xc6f05771.
//
// Solidity: function safeMint(string _tokenURI, address rent_token_address) returns()
func (_Birden *BirdenTransactor) SafeMint(opts *bind.TransactOpts, _tokenURI string, rent_token_address common.Address) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeMint", _tokenURI, rent_token_address)
}

// SafeMint is a paid mutator transaction binding the contract method 0xc6f05771.
//
// Solidity: function safeMint(string _tokenURI, address rent_token_address) returns()
func (_Birden *BirdenSession) SafeMint(_tokenURI string, rent_token_address common.Address) (*types.Transaction, error) {
	return _Birden.Contract.SafeMint(&_Birden.TransactOpts, _tokenURI, rent_token_address)
}

// SafeMint is a paid mutator transaction binding the contract method 0xc6f05771.
//
// Solidity: function safeMint(string _tokenURI, address rent_token_address) returns()
func (_Birden *BirdenTransactorSession) SafeMint(_tokenURI string, rent_token_address common.Address) (*types.Transaction, error) {
	return _Birden.Contract.SafeMint(&_Birden.TransactOpts, _tokenURI, rent_token_address)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Birden *BirdenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom0(&_Birden.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom0(&_Birden.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom1 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom1(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom1", from, to, tokenId)
}

// SafeTransferFrom1 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) SafeTransferFrom1(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom1(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom1 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom1(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom1(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom2 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom2(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom2", from, to, tokenId, _data)
}

// SafeTransferFrom2 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Birden *BirdenSession) SafeTransferFrom2(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom2(&_Birden.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom2 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom2(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom2(&_Birden.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom3 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom3(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom3", from, to, tokenId)
}

// SafeTransferFrom3 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) SafeTransferFrom3(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom3(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom3 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom3(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom3(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom4 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom4(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom4", from, to, tokenId, data)
}

// SafeTransferFrom4 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Birden *BirdenSession) SafeTransferFrom4(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom4(&_Birden.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom4 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom4(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom4(&_Birden.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom5 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom5(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom5", from, to, tokenId)
}

// SafeTransferFrom5 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) SafeTransferFrom5(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom5(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom5 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom5(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom5(&_Birden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom6 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Birden *BirdenTransactor) SafeTransferFrom6(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "safeTransferFrom6", from, to, tokenId, data)
}

// SafeTransferFrom6 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Birden *BirdenSession) SafeTransferFrom6(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom6(&_Birden.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom6 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Birden *BirdenTransactorSession) SafeTransferFrom6(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Birden.Contract.SafeTransferFrom6(&_Birden.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Birden *BirdenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Birden *BirdenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll(&_Birden.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Birden *BirdenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll(&_Birden.TransactOpts, operator, approved)
}

// SetApprovalForAll0 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Birden *BirdenTransactor) SetApprovalForAll0(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "setApprovalForAll0", operator, approved)
}

// SetApprovalForAll0 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Birden *BirdenSession) SetApprovalForAll0(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll0(&_Birden.TransactOpts, operator, approved)
}

// SetApprovalForAll0 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Birden *BirdenTransactorSession) SetApprovalForAll0(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll0(&_Birden.TransactOpts, operator, approved)
}

// SetApprovalForAll1 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Birden *BirdenTransactor) SetApprovalForAll1(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "setApprovalForAll1", operator, _approved)
}

// SetApprovalForAll1 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Birden *BirdenSession) SetApprovalForAll1(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll1(&_Birden.TransactOpts, operator, _approved)
}

// SetApprovalForAll1 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Birden *BirdenTransactorSession) SetApprovalForAll1(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll1(&_Birden.TransactOpts, operator, _approved)
}

// SetApprovalForAll2 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Birden *BirdenTransactor) SetApprovalForAll2(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "setApprovalForAll2", operator, _approved)
}

// SetApprovalForAll2 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Birden *BirdenSession) SetApprovalForAll2(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll2(&_Birden.TransactOpts, operator, _approved)
}

// SetApprovalForAll2 is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Birden *BirdenTransactorSession) SetApprovalForAll2(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Birden.Contract.SetApprovalForAll2(&_Birden.TransactOpts, operator, _approved)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Transfer(&_Birden.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Transfer(&_Birden.TransactOpts, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) Transfer0(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transfer0", to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenSession) Transfer0(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Transfer0(&_Birden.TransactOpts, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) Transfer0(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Transfer0(&_Birden.TransactOpts, to, amount)
}

// Transfer1 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) Transfer1(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transfer1", to, amount)
}

// Transfer1 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenSession) Transfer1(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Transfer1(&_Birden.TransactOpts, to, amount)
}

// Transfer1 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) Transfer1(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.Transfer1(&_Birden.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) TransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom0", from, to, amount)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenSession) TransferFrom0(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom0(&_Birden.TransactOpts, from, to, amount)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) TransferFrom0(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom0(&_Birden.TransactOpts, from, to, amount)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) TransferFrom1(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom1", from, to, tokenId)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) TransferFrom1(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom1(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) TransferFrom1(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom1(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) TransferFrom2(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom2", from, to, amount)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenSession) TransferFrom2(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom2(&_Birden.TransactOpts, from, to, amount)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) TransferFrom2(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom2(&_Birden.TransactOpts, from, to, amount)
}

// TransferFrom3 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactor) TransferFrom3(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom3", from, to, amount)
}

// TransferFrom3 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenSession) TransferFrom3(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom3(&_Birden.TransactOpts, from, to, amount)
}

// TransferFrom3 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Birden *BirdenTransactorSession) TransferFrom3(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom3(&_Birden.TransactOpts, from, to, amount)
}

// TransferFrom4 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) TransferFrom4(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom4", from, to, tokenId)
}

// TransferFrom4 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) TransferFrom4(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom4(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom4 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) TransferFrom4(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom4(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom5 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactor) TransferFrom5(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.contract.Transact(opts, "transferFrom5", from, to, tokenId)
}

// TransferFrom5 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenSession) TransferFrom5(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom5(&_Birden.TransactOpts, from, to, tokenId)
}

// TransferFrom5 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Birden *BirdenTransactorSession) TransferFrom5(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Birden.Contract.TransferFrom5(&_Birden.TransactOpts, from, to, tokenId)
}

// BirdenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Birden contract.
type BirdenApprovalIterator struct {
	Event *BirdenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval represents a Approval event raised by the Birden contract.
type BirdenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BirdenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApprovalIterator{contract: _Birden.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BirdenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval)
				if err := _Birden.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseApproval(log types.Log) (*BirdenApproval, error) {
	event := new(BirdenApproval)
	if err := _Birden.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApproval0Iterator is returned from FilterApproval0 and is used to iterate over the raw logs and unpacked data for Approval0 events raised by the Birden contract.
type BirdenApproval0Iterator struct {
	Event *BirdenApproval0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApproval0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApproval0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApproval0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval0 represents a Approval0 event raised by the Birden contract.
type BirdenApproval0 struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval0 is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) FilterApproval0(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BirdenApproval0Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval0", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApproval0Iterator{contract: _Birden.contract, event: "Approval0", logs: logs, sub: sub}, nil
}

// WatchApproval0 is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) WatchApproval0(opts *bind.WatchOpts, sink chan<- *BirdenApproval0, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval0", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval0)
				if err := _Birden.contract.UnpackLog(event, "Approval0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval0 is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) ParseApproval0(log types.Log) (*BirdenApproval0, error) {
	event := new(BirdenApproval0)
	if err := _Birden.contract.UnpackLog(event, "Approval0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApproval1Iterator is returned from FilterApproval1 and is used to iterate over the raw logs and unpacked data for Approval1 events raised by the Birden contract.
type BirdenApproval1Iterator struct {
	Event *BirdenApproval1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApproval1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApproval1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApproval1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval1 represents a Approval1 event raised by the Birden contract.
type BirdenApproval1 struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval1 is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterApproval1(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BirdenApproval1Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval1", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApproval1Iterator{contract: _Birden.contract, event: "Approval1", logs: logs, sub: sub}, nil
}

// WatchApproval1 is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchApproval1(opts *bind.WatchOpts, sink chan<- *BirdenApproval1, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval1", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval1)
				if err := _Birden.contract.UnpackLog(event, "Approval1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval1 is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseApproval1(log types.Log) (*BirdenApproval1, error) {
	event := new(BirdenApproval1)
	if err := _Birden.contract.UnpackLog(event, "Approval1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApproval2Iterator is returned from FilterApproval2 and is used to iterate over the raw logs and unpacked data for Approval2 events raised by the Birden contract.
type BirdenApproval2Iterator struct {
	Event *BirdenApproval2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApproval2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApproval2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApproval2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval2 represents a Approval2 event raised by the Birden contract.
type BirdenApproval2 struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval2 is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) FilterApproval2(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BirdenApproval2Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval2", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApproval2Iterator{contract: _Birden.contract, event: "Approval2", logs: logs, sub: sub}, nil
}

// WatchApproval2 is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) WatchApproval2(opts *bind.WatchOpts, sink chan<- *BirdenApproval2, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval2", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval2)
				if err := _Birden.contract.UnpackLog(event, "Approval2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval2 is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) ParseApproval2(log types.Log) (*BirdenApproval2, error) {
	event := new(BirdenApproval2)
	if err := _Birden.contract.UnpackLog(event, "Approval2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApproval3Iterator is returned from FilterApproval3 and is used to iterate over the raw logs and unpacked data for Approval3 events raised by the Birden contract.
type BirdenApproval3Iterator struct {
	Event *BirdenApproval3 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApproval3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval3)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval3)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApproval3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApproval3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval3 represents a Approval3 event raised by the Birden contract.
type BirdenApproval3 struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval3 is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) FilterApproval3(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BirdenApproval3Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval3", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApproval3Iterator{contract: _Birden.contract, event: "Approval3", logs: logs, sub: sub}, nil
}

// WatchApproval3 is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) WatchApproval3(opts *bind.WatchOpts, sink chan<- *BirdenApproval3, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval3", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval3)
				if err := _Birden.contract.UnpackLog(event, "Approval3", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval3 is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Birden *BirdenFilterer) ParseApproval3(log types.Log) (*BirdenApproval3, error) {
	event := new(BirdenApproval3)
	if err := _Birden.contract.UnpackLog(event, "Approval3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApproval4Iterator is returned from FilterApproval4 and is used to iterate over the raw logs and unpacked data for Approval4 events raised by the Birden contract.
type BirdenApproval4Iterator struct {
	Event *BirdenApproval4 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApproval4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval4)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval4)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApproval4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApproval4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval4 represents a Approval4 event raised by the Birden contract.
type BirdenApproval4 struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval4 is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterApproval4(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BirdenApproval4Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval4", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApproval4Iterator{contract: _Birden.contract, event: "Approval4", logs: logs, sub: sub}, nil
}

// WatchApproval4 is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchApproval4(opts *bind.WatchOpts, sink chan<- *BirdenApproval4, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval4", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval4)
				if err := _Birden.contract.UnpackLog(event, "Approval4", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval4 is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseApproval4(log types.Log) (*BirdenApproval4, error) {
	event := new(BirdenApproval4)
	if err := _Birden.contract.UnpackLog(event, "Approval4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApproval5Iterator is returned from FilterApproval5 and is used to iterate over the raw logs and unpacked data for Approval5 events raised by the Birden contract.
type BirdenApproval5Iterator struct {
	Event *BirdenApproval5 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApproval5Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApproval5)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApproval5)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApproval5Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApproval5Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApproval5 represents a Approval5 event raised by the Birden contract.
type BirdenApproval5 struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval5 is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterApproval5(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BirdenApproval5Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Approval5", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApproval5Iterator{contract: _Birden.contract, event: "Approval5", logs: logs, sub: sub}, nil
}

// WatchApproval5 is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchApproval5(opts *bind.WatchOpts, sink chan<- *BirdenApproval5, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Approval5", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApproval5)
				if err := _Birden.contract.UnpackLog(event, "Approval5", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval5 is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseApproval5(log types.Log) (*BirdenApproval5, error) {
	event := new(BirdenApproval5)
	if err := _Birden.contract.UnpackLog(event, "Approval5", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Birden contract.
type BirdenApprovalForAllIterator struct {
	Event *BirdenApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApprovalForAll represents a ApprovalForAll event raised by the Birden contract.
type BirdenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BirdenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApprovalForAllIterator{contract: _Birden.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BirdenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApprovalForAll)
				if err := _Birden.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) ParseApprovalForAll(log types.Log) (*BirdenApprovalForAll, error) {
	event := new(BirdenApprovalForAll)
	if err := _Birden.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApprovalForAll0Iterator is returned from FilterApprovalForAll0 and is used to iterate over the raw logs and unpacked data for ApprovalForAll0 events raised by the Birden contract.
type BirdenApprovalForAll0Iterator struct {
	Event *BirdenApprovalForAll0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApprovalForAll0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApprovalForAll0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApprovalForAll0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApprovalForAll0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApprovalForAll0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApprovalForAll0 represents a ApprovalForAll0 event raised by the Birden contract.
type BirdenApprovalForAll0 struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll0 is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) FilterApprovalForAll0(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BirdenApprovalForAll0Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "ApprovalForAll0", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApprovalForAll0Iterator{contract: _Birden.contract, event: "ApprovalForAll0", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll0 is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) WatchApprovalForAll0(opts *bind.WatchOpts, sink chan<- *BirdenApprovalForAll0, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "ApprovalForAll0", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApprovalForAll0)
				if err := _Birden.contract.UnpackLog(event, "ApprovalForAll0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll0 is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) ParseApprovalForAll0(log types.Log) (*BirdenApprovalForAll0, error) {
	event := new(BirdenApprovalForAll0)
	if err := _Birden.contract.UnpackLog(event, "ApprovalForAll0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApprovalForAll1Iterator is returned from FilterApprovalForAll1 and is used to iterate over the raw logs and unpacked data for ApprovalForAll1 events raised by the Birden contract.
type BirdenApprovalForAll1Iterator struct {
	Event *BirdenApprovalForAll1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApprovalForAll1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApprovalForAll1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApprovalForAll1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApprovalForAll1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApprovalForAll1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApprovalForAll1 represents a ApprovalForAll1 event raised by the Birden contract.
type BirdenApprovalForAll1 struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll1 is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) FilterApprovalForAll1(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BirdenApprovalForAll1Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "ApprovalForAll1", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApprovalForAll1Iterator{contract: _Birden.contract, event: "ApprovalForAll1", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll1 is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) WatchApprovalForAll1(opts *bind.WatchOpts, sink chan<- *BirdenApprovalForAll1, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "ApprovalForAll1", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApprovalForAll1)
				if err := _Birden.contract.UnpackLog(event, "ApprovalForAll1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll1 is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) ParseApprovalForAll1(log types.Log) (*BirdenApprovalForAll1, error) {
	event := new(BirdenApprovalForAll1)
	if err := _Birden.contract.UnpackLog(event, "ApprovalForAll1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenApprovalForAll2Iterator is returned from FilterApprovalForAll2 and is used to iterate over the raw logs and unpacked data for ApprovalForAll2 events raised by the Birden contract.
type BirdenApprovalForAll2Iterator struct {
	Event *BirdenApprovalForAll2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenApprovalForAll2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenApprovalForAll2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenApprovalForAll2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenApprovalForAll2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenApprovalForAll2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenApprovalForAll2 represents a ApprovalForAll2 event raised by the Birden contract.
type BirdenApprovalForAll2 struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll2 is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) FilterApprovalForAll2(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BirdenApprovalForAll2Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "ApprovalForAll2", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BirdenApprovalForAll2Iterator{contract: _Birden.contract, event: "ApprovalForAll2", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll2 is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) WatchApprovalForAll2(opts *bind.WatchOpts, sink chan<- *BirdenApprovalForAll2, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "ApprovalForAll2", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenApprovalForAll2)
				if err := _Birden.contract.UnpackLog(event, "ApprovalForAll2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll2 is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Birden *BirdenFilterer) ParseApprovalForAll2(log types.Log) (*BirdenApprovalForAll2, error) {
	event := new(BirdenApprovalForAll2)
	if err := _Birden.contract.UnpackLog(event, "ApprovalForAll2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Birden contract.
type BirdenTransferIterator struct {
	Event *BirdenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer represents a Transfer event raised by the Birden contract.
type BirdenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BirdenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransferIterator{contract: _Birden.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BirdenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer)
				if err := _Birden.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseTransfer(log types.Log) (*BirdenTransfer, error) {
	event := new(BirdenTransfer)
	if err := _Birden.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransfer0Iterator is returned from FilterTransfer0 and is used to iterate over the raw logs and unpacked data for Transfer0 events raised by the Birden contract.
type BirdenTransfer0Iterator struct {
	Event *BirdenTransfer0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransfer0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransfer0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer0 represents a Transfer0 event raised by the Birden contract.
type BirdenTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer0 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BirdenTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransfer0Iterator{contract: _Birden.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

// WatchTransfer0 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BirdenTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer0)
				if err := _Birden.contract.UnpackLog(event, "Transfer0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer0 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) ParseTransfer0(log types.Log) (*BirdenTransfer0, error) {
	event := new(BirdenTransfer0)
	if err := _Birden.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransfer1Iterator is returned from FilterTransfer1 and is used to iterate over the raw logs and unpacked data for Transfer1 events raised by the Birden contract.
type BirdenTransfer1Iterator struct {
	Event *BirdenTransfer1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransfer1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransfer1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransfer1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer1 represents a Transfer1 event raised by the Birden contract.
type BirdenTransfer1 struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer1 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterTransfer1(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BirdenTransfer1Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer1", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransfer1Iterator{contract: _Birden.contract, event: "Transfer1", logs: logs, sub: sub}, nil
}

// WatchTransfer1 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchTransfer1(opts *bind.WatchOpts, sink chan<- *BirdenTransfer1, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer1", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer1)
				if err := _Birden.contract.UnpackLog(event, "Transfer1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer1 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseTransfer1(log types.Log) (*BirdenTransfer1, error) {
	event := new(BirdenTransfer1)
	if err := _Birden.contract.UnpackLog(event, "Transfer1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransfer2Iterator is returned from FilterTransfer2 and is used to iterate over the raw logs and unpacked data for Transfer2 events raised by the Birden contract.
type BirdenTransfer2Iterator struct {
	Event *BirdenTransfer2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransfer2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransfer2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransfer2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer2 represents a Transfer2 event raised by the Birden contract.
type BirdenTransfer2 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer2 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) FilterTransfer2(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BirdenTransfer2Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer2", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransfer2Iterator{contract: _Birden.contract, event: "Transfer2", logs: logs, sub: sub}, nil
}

// WatchTransfer2 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) WatchTransfer2(opts *bind.WatchOpts, sink chan<- *BirdenTransfer2, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer2", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer2)
				if err := _Birden.contract.UnpackLog(event, "Transfer2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer2 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) ParseTransfer2(log types.Log) (*BirdenTransfer2, error) {
	event := new(BirdenTransfer2)
	if err := _Birden.contract.UnpackLog(event, "Transfer2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransfer3Iterator is returned from FilterTransfer3 and is used to iterate over the raw logs and unpacked data for Transfer3 events raised by the Birden contract.
type BirdenTransfer3Iterator struct {
	Event *BirdenTransfer3 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransfer3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer3)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer3)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransfer3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransfer3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer3 represents a Transfer3 event raised by the Birden contract.
type BirdenTransfer3 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer3 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) FilterTransfer3(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BirdenTransfer3Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer3", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransfer3Iterator{contract: _Birden.contract, event: "Transfer3", logs: logs, sub: sub}, nil
}

// WatchTransfer3 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) WatchTransfer3(opts *bind.WatchOpts, sink chan<- *BirdenTransfer3, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer3", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer3)
				if err := _Birden.contract.UnpackLog(event, "Transfer3", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer3 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Birden *BirdenFilterer) ParseTransfer3(log types.Log) (*BirdenTransfer3, error) {
	event := new(BirdenTransfer3)
	if err := _Birden.contract.UnpackLog(event, "Transfer3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransfer4Iterator is returned from FilterTransfer4 and is used to iterate over the raw logs and unpacked data for Transfer4 events raised by the Birden contract.
type BirdenTransfer4Iterator struct {
	Event *BirdenTransfer4 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransfer4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer4)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer4)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransfer4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransfer4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer4 represents a Transfer4 event raised by the Birden contract.
type BirdenTransfer4 struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer4 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterTransfer4(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BirdenTransfer4Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer4", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransfer4Iterator{contract: _Birden.contract, event: "Transfer4", logs: logs, sub: sub}, nil
}

// WatchTransfer4 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchTransfer4(opts *bind.WatchOpts, sink chan<- *BirdenTransfer4, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer4", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer4)
				if err := _Birden.contract.UnpackLog(event, "Transfer4", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer4 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseTransfer4(log types.Log) (*BirdenTransfer4, error) {
	event := new(BirdenTransfer4)
	if err := _Birden.contract.UnpackLog(event, "Transfer4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BirdenTransfer5Iterator is returned from FilterTransfer5 and is used to iterate over the raw logs and unpacked data for Transfer5 events raised by the Birden contract.
type BirdenTransfer5Iterator struct {
	Event *BirdenTransfer5 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BirdenTransfer5Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirdenTransfer5)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BirdenTransfer5)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BirdenTransfer5Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirdenTransfer5Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirdenTransfer5 represents a Transfer5 event raised by the Birden contract.
type BirdenTransfer5 struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer5 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) FilterTransfer5(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BirdenTransfer5Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.FilterLogs(opts, "Transfer5", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BirdenTransfer5Iterator{contract: _Birden.contract, event: "Transfer5", logs: logs, sub: sub}, nil
}

// WatchTransfer5 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) WatchTransfer5(opts *bind.WatchOpts, sink chan<- *BirdenTransfer5, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Birden.contract.WatchLogs(opts, "Transfer5", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirdenTransfer5)
				if err := _Birden.contract.UnpackLog(event, "Transfer5", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer5 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Birden *BirdenFilterer) ParseTransfer5(log types.Log) (*BirdenTransfer5, error) {
	event := new(BirdenTransfer5)
	if err := _Birden.contract.UnpackLog(event, "Transfer5", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
