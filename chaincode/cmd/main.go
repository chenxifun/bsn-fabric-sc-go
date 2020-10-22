package main

import (
	"bsn-fabric-sc/chaincode"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(chaincode.SCChaincode))
	if err != nil {
		fmt.Printf("Error starting SCChaincode: %s", err)
	}
}
