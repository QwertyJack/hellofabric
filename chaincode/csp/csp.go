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
	"encoding/json"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/util"
)

// CSPChaincode implementation
type CSPChaincode struct {
}

// State 请求状态
type State int

const (
	UserInit State = 0 // 用户发起请求
	Chagened State = 1 // TPA 已生成挑战
	CSPRespd State = 2 // CSP 已响应挑战
	TChecked State = 3 // TPA 已检查结果
)

// UserRequest 完整性验证请求，只能被用户初始化
type UserRequest string

// Challange 挑战，只能被 TPA 生成
type Challenge string

// 证明，只能被 CSP 生成
type Proof string

// 结果，只能被 TPA 决定
type Result string

type Tx struct {
	UserID string      // 用户 id
	CSPID  string      // CSP id
	TPAID  string      // TPA id
	Status State       // 请求当前状态
	UR     UserRequest // 用户发出的申请
	Ch     Challenge   // TSP 生成的挑战
	Pf     Proof       // CSP 生成的证明
	Re     Result      // TSP 的检查结果
}

func newTxFromByte(value []byte) Tx {
	var ret Tx
	err := json.Unmarshal(value, &ret)
	if err != nil {
		log.Printf("Unmarshal err: %v", err)
	}
	return ret
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(CSPChaincode))
	if err != nil {
		log.Fatalf("Init error: %s", err)
	}
}

// Init initializes chaincode
func (t *CSPChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
func (t *CSPChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	// Handle different functions
	if function == "_set" { // 调试
		return t.set(stub, args)
	}
	if function == "_get" { // 打印值，调试
		return t.get(stub, args)
	}
	if function == "history" { // 查询历史
		return t.history(stub, args)
	}
	if function == "request" { // 用户发起请求
		return t.request(stub, args)
	}
	if function == "challange" { // TPA 生成挑战
		return t.challenge(stub, args)
	}
	if function == "proof" { // CSP 完成证明
		return t.proof(stub, args)
	}
	if function == "check" { // TPA 的检查结果
		return t.check(stub, args)
	}
	return shim.Error("Unknown function: " + function)
}

func (t *CSPChaincode) set(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// TODO: impl
	return shim.Success(nil)
}

// get - get value of a key
func (t *CSPChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Func get expecting 1")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	tx := newTxFromByte(value)

	log.Printf("- %v: %v", args[0], tx)
	return shim.Success(value)
}

func (t *CSPChaincode) history(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var id, jsonResp string
	id = args[0]
	valAsbytes, err := stub.GetState(id) //get the data from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Data does not exist, id: " + id + "\"}"
		return shim.Error(jsonResp)
	}
	//tx := newTxFromByte(valAsbytes)
	//currentUser, err := util.GetUser(stub)
	//if err != nil {
	//	return shim.Error(err.Error())
	//} else if !util.Contains(dataJSON.Owner, currentUser) && !util.IsAdmin(currentUser) {
	//	return shim.Error("Permission denied")
	//}

	log.Printf("- start history, id: %v", id)

	resultsIterator, err := stub.GetHistoryForKey(id)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	bResult, err := util.HistoryIter2json(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(bResult)
}

func (t *CSPChaincode) request(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// 0 	1 	2 	3
	// key	CSP TPA	UR

	// TODO: sanitize input
	// TODO: error handle
	user, _ := util.GetUser(stub)
	// TODO: check user perm
	data := Tx{
		UserID: user,
		CSPID:  args[1],
		TPAID:  args[2],
		Status: UserInit,
		UR:     UserRequest(args[3]),
		Ch:     Challenge(""),
		Pf:     Proof(""),
		Re:     Result(""),
	}
	dataJSONasBytes, _ := json.Marshal(data)
	// TODO: check whether args[0] exists before writing state
	_ = stub.PutState(args[0], dataJSONasBytes)

	return shim.Success(nil)
}

func (t *CSPChaincode) challenge(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// 0 	1
	// key 	challenge

	// TODO: sanitize input
	// TODO: error handle
	_, _ = util.GetUser(stub)
	// TODO: check user perm
	value, _ := stub.GetState(args[0])
	tx := newTxFromByte(value)
	// TODO: check status

	// update
	tx.Ch = Challenge(args[1])
	tx.Status = Chagened

	// write back
	dataJSONasBytes, _ := json.Marshal(tx)
	_ = stub.PutState(args[0], dataJSONasBytes)
	return shim.Success(nil)
}

func (t *CSPChaincode) proof(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// 0 	1
	// key 	proof

	// TODO: sanitize input
	// TODO: error handle
	_, _ = util.GetUser(stub)
	// TODO: check user perm
	value, _ := stub.GetState(args[0])
	tx := newTxFromByte(value)
	// TODO: check status

	// update
	tx.Pf = Proof(args[1])
	tx.Status = CSPRespd

	// write back
	dataJSONasBytes, _ := json.Marshal(tx)
	_ = stub.PutState(args[0], dataJSONasBytes)
	return shim.Success(nil)
}

func (t *CSPChaincode) check(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// 0 	1
	// key 	result

	// TODO: sanitize input
	// TODO: error handle
	_, _ = util.GetUser(stub)
	// TODO: check user perm
	value, _ := stub.GetState(args[0])
	tx := newTxFromByte(value)
	// TODO: check status

	// update
	tx.Ch = Challenge(args[1])
	tx.Status = TChecked

	// write back
	dataJSONasBytes, _ := json.Marshal(tx)
	_ = stub.PutState(args[0], dataJSONasBytes)
	return shim.Success(nil)
}
