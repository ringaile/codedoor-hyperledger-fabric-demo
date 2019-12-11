package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Jewelry struct {
	Owner string `json:"owner"`
	Type string `json:"type"`
	Material string `json:"material"`
	Color string `json:"color"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryJewelry" {
		return s.queryJewelry(APIstub, args)
	} else if function == "createJewelry" {
		return s.createJewelry(APIstub, args)
	} else if function == "changeJewelryOwner" {
		return s.changeJewelryOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryJewelry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	docrecordAsBytes, _ := APIstub.GetState(args[0])

	if docrecordAsBytes == nil {
		return shim.Error("Document not found: " + args[0])
	}

	return shim.Success(docrecordAsBytes)
}

func (s *SmartContract) createJewelry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var jewelry = Jewelry{Owner: args[1], Type: args[2], Material: args[3], Color: args[4]}
	jewelryAsBytes, _ := json.Marshal(jewelry)
	APIstub.PutState(args[0], jewelryAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeJewelryOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	jewelryAsBytes, _ := APIstub.GetState(args[0])
	jewelry := Jewelry{}

	json.Unmarshal(jewelryAsBytes, &jewelry)
	jewelry.Owner = args[1]

	jewelryAsBytes, _ = json.Marshal(jewelry)
	APIstub.PutState(args[0], jewelryAsBytes)

	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

