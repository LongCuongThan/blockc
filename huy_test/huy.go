package huy_test

import (
	"fmt"
	"log"

	"huy_smart/uniswapv2Router"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	infuraUrl         = "wss://mainnet.infura.io/ws/v3/5e18a9566f8f4d83b56984c2c15b8f58"
	uniswapv2Contract = "0xbF16540c857B4e32cE6C37d2F7725C8eEC869B8b"
	uniswapv3Contract = "0x4585fe77225b41b697c938b018e2ac67ac5a20c0"
	uniswapUniversalC = "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"
)

func Testing() {
	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client", err)
	}
	// Tạo một kênh để nhận các giao dịch mới
	address := common.HexToAddress(uniswapv2Contract)
	contract, err := uniswapv2Router.NewUniswapv2Router(address, client)
	if err != nil {
		log.Fatal("Failed to create Uniswapv2Router", err)
	}
	callOpts := &bind.CallOpts{}
	reserve, err := contract.GetReserves(callOpts)
	if err != nil {
		log.Fatal("Failed to get reserves", err)
	}
	fmt.Println(reserve.Reserve0, reserve.Reserve1, reserve.BlockTimestampLast)

	token0, err := contract.Token0(callOpts)
	if err != nil {
		log.Fatal("Failed to get token0", err)
	}
	fmt.Println(token0)
	token1, err := contract.Token1(callOpts)
	if err != nil {
		log.Fatal("Failed to get token1", err)
	}
	fmt.Println(token1)
	decimalToken0 := getDecimalByToken(client, token0.String())
	fmt.Printf("decimal token 0: %v\n", decimalToken0)
	decimalToken1 := getDecimalByToken(client, token1.String())
	fmt.Printf("decimal token 1: %v", decimalToken1)
}

func getDecimalByToken(client *ethclient.Client, address string) int {
	addr := common.HexToAddress(address)
	contract, err := uniswapv2Router.NewUniswapv2Router(addr, client)
	if err != nil {
		log.Fatal("Failed to create Uniswapv2Router", err)
	}
	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to get decimals", err)
	}
	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to get Token name", err)
	}
	fmt.Printf("\nToken name: %v, decimals: %v\n", name, decimals)
	return int(decimals)
}

func handleTransaction(client ethclient.Client, txHash common.Hash) {

}
