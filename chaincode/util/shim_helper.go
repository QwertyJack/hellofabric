//
// history.go
// Copyright (C) 2018 jack <jack@HP-WorkStation>
//
// Distributed under terms of the MIT license.
//

package util

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// HistoryIter2json parses the history iterator and constuct history lists
func HistoryIter2json(resultsIterator shim.HistoryQueryIteratorInterface) ([]byte, error) {

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"tx_id\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"is_delete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return buffer.Bytes(), nil
}

// GetUser gets the user name of the caller
func GetUser(stub shim.ChaincodeStubInterface) (string, error) {
	creatorByte, _ := stub.GetCreator()
	certStart := bytes.IndexAny(creatorByte, "-----BEGIN")
	if certStart == -1 {
		return "", errors.New("No certificate found")
	}
	certText := creatorByte[certStart:]
	bl, _ := pem.Decode(certText)
	if bl == nil {
		return "", errors.New("Could not decode the PEM structure")
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return "", errors.New("ParseCertificate failed")
	}
	uname := cert.Subject.CommonName
	return uname, nil
}
