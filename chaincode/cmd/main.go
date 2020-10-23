package main

import (
	"fmt"
	"github.com/chenxifun/bsn-fabric-sc-go/chaincode"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(chaincode.SCChaincode))
	if err != nil {
		fmt.Printf("Error starting SCChaincode: %s", err)
	}
}
