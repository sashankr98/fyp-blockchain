package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"github.com/segmentio/ksuid"
)

// storeTestBlock - function to write a basic block to check working of chaincode
func storeTestBlock(stub shim.ChaincodeStubInterface) pb.Response {
	var item CultivationData

	item.BatchID = ksuid.New().String()
	item.Stage = CultivationStage
	item.FarmerName = "Reliance Farm"
	item.FarmAddress = "Nashik, Maharashtra, India"
	item.ExporterName = "Express Export & Import Service"
	item.ImporerName = "FQ Export & Import Service"
	itemAsBytes, err := json.Marshal(item)
	if err != nil {
		return shim.Error(err.Error())
	}

	stageAsBytes, _ := json.Marshal(item.Stage)
	stub.PutState(item.BatchID, stageAsBytes)

	itemKey, _ := stub.CreateCompositeKey(CultivationStage, []string{item.BatchID, item.Stage})
	stub.PutState(itemKey, itemAsBytes)

	rawAsset, err := stub.GetState(item.BatchID)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(rawAsset)
}

// updateTestBlock - function to change the state of an existing batch
func updateTestBlock(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("updateTestBlock expects only one arg - BatchID")
	}

	stageAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	var stage string
	json.Unmarshal(stageAsBytes, &stage)

	switch stage {
	default:
		return shim.Error("Error while querying the blockchain")
	case CultivationStage:
		stage = FarmInspectorStage
		break
	case FarmInspectorStage:
		stage = HarvesterStage
		break
	case HarvesterStage:
		stage = ExporterStage
		break
	case ExporterStage:
		stage = ImporterStage
		break
	case ImporterStage:
		stage = ProcessorStage
		break
	case ProcessorStage:
		stage = CompletedStage
		break
	}
	// TODO: Make store and update more generic
	nextItem := FarmInspectorData{}
	nextItem.Stage = stage
	nextItem.SeedType = "Coffee Arabica"
	nextItem.CoffeeFamily = "Arabica Coffee Family"
	nextItem.FertilizerUsed = "Organic Fertilizer"

	nextItemAsBytes, err := json.Marshal(nextItem)
	key, _ := stub.CreateCompositeKey(stage, []string{args[0], stage})
	stub.PutState(key, nextItemAsBytes)

	return shim.Success(nil)
}

// queryTestBlock - function to retrieve test block history
func queryTestBlock(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("queryTestBlock expects only one arg - BatchID")
	}

	stageAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	var stage string
	json.Unmarshal(stageAsBytes, &stage)
	key, _ := stub.CreateCompositeKey(stage, []string{args[0], stage})
	itemAsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(itemAsBytes)
	// switch stage {
	// default:
	// 	return shim.Error("Error while querying the blockchain")
	// case CultivationStage:
	// 	item := CultivationData{}
	// 	json.Unmarshal(itemAsBytes, &item)
	// 	return shim.Success(itemAsBytes)
	// case FarmInspectorStage:
	// 	item := FarmInspectorData{}
	// 	json.Unmarshal(itemAsBytes, &item)
	// 	return shim.Success(itemAsBytes)
	// case HarvesterStage:
	// 	item := HarvesterData{}
	// 	json.Unmarshal(itemAsBytes, &item)
	// 	return shim.Success(itemAsBytes)
	// case ExporterStage:
	// 	item := ExporterData{}
	// 	json.Unmarshal(itemAsBytes, &item)
	// 	return shim.Success(itemAsBytes)
	// case ImporterStage:
	// 	item := ImporterData{}
	// 	json.Unmarshal(itemAsBytes, &item)
	// 	return shim.Success(itemAsBytes)
	// case ProcessorStage:
	// 	item := ProcessorData{}
	// 	json.Unmarshal(itemAsBytes, &item)
	// 	return shim.Success(itemAsBytes)
	// }
}
