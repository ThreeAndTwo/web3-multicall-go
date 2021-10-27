package multicall_test

import (
	"encoding/json"
	"fmt"
	"github.com/alethio/web3-go/ethrpc"
	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
	"github.com/alethio/web3-multicall-go/multicall"
	"testing"
	"time"
)

func TestExampleViwCall(t *testing.T) {
	eth, err := getETH("https://mainnet.infura.io/v3/17ed7fe26d014e5b9be7dfff5368c69d")
	vc := multicall.NewViewCall(
		"key.1",
		"0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643",
		"name()(string)",  // Compound Dai
		[]interface{}{},
	)

	vc1 := multicall.NewViewCall(
		"key.2",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		"name()(string)",	// Dai Stablecoin
		[]interface{}{},
	)

	vcs := multicall.ViewCalls{vc, vc1}
	mc, _ := multicall.New(eth, multicall.ContractAddress("0x5eb3fa2dfecdde21c950813c665e9364fa609bd2"))
	block := "latest"
	res, err := mc.Call(vcs, block)
	if err != nil {
		panic(err)
	}

	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
	fmt.Println(res)
	fmt.Println(err)

}

func getETH(url string) (ethrpc.ETHInterface, error) {
	provider, err := httprpc.New(url)
	if err != nil {
		return nil, err
	}
	provider.SetHTTPTimeout(5 * time.Second)
	return ethrpc.New(provider)
}
