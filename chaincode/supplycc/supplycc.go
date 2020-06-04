package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SupplyChainContract - basic smart contract structure
type SupplyChainContract struct {
}

// Init is called when the chaincode is instantiated
func (c *SupplyChainContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke is called when data needs to be written to the blockchain
func (c *SupplyChainContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Query is called to read data from the blockchain
func (c *SupplyChainContract) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SupplyChainContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
