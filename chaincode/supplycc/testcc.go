package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/segmentio/ksuid"
)

func testCreate(stub shim.ChaincodeStubInterface) pb.Response {
	var args []string
	id := ksuid.New().String()
	args = append(args, id)

	obj := "{\"farmer-name\":\"Reliance Farm\",\"farm-address\":\"Nashik, Maharashtra, India\",\"exporter-name\":\"Express Export & Import Service\",\"importer-name\":\"FQ Export & Import Service\"}"
	args = append(args, obj)

	return createBatch(stub, args)
}

func testUpdate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	nextStage := FarmInspectorStage
	args = append(args, nextStage)

	obj := "{\"seed-type\":\"Coffee Arabica\",\"coffee-family\":\"Arabica Coffee Family\",\"fertilizer-used\":\"Organic Fertilizer\"}"
	args = append(args, obj)

	return updateBatch(stub, args)
}
