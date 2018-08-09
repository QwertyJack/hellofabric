/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		log.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	log.Println("invoke is running " + function)

	// Handle different functions
	if function == "set" {
		return t.set(stub, args)
	} else if function == "get" {
		return t.get(stub, args)
    }

	log.Println("Invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// set - create or update value of a key
// ============================================================
func (t *SimpleChaincode) set(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Func set expecting 2")
	}

    err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== data saved and indexed. Return success ====
    log.Println("- set " + args[0] + ", value: " + args[1])
	return shim.Success(nil)
}

// get - get value of a key
func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Func get expecting 1")
	}

	val, err := stub.GetState(args[0]) //get the data from chaincode state
	if err != nil {
		return shim.Error(err.Error())
	}

    log.Println("- get " + args[0] + ", value: " + string(val[:]))
	return shim.Success(val)
}

