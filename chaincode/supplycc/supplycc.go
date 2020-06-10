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
	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "storeTestBlock":
		return storeTestBlock(stub)
	case "updateTestBlock":
		return updateTestBlock(stub, args)
	case "queryTestBlock":
		return queryTestBlock(stub, args)
	default:
		return shim.Error("Function does not exist")
	}
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SupplyChainContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
