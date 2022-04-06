package client

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"strconv"
)

type StoreData struct {
	owner      string
	identifier string
	rawData    RawDataBundle
	timestamp  uint64
}

type RawDataBundle struct {
	rawData  []byte
	mimeType string
	fileName string
}

func SubmitData(contract *client.Contract, data StoreData) {
	res, commit, err := contract.SubmitAsync("AddDataRecord", client.WithArguments(data.owner, data.identifier, strconv.FormatUint(data.timestamp, 10)))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
	}
	fmt.Printf("Successfully submitted transaction to transfer ownership from %s to Mark. \n", string(res))
	fmt.Println("Waiting for transaction commit.")

	if status, err := commit.Status(); err != nil {
		panic(fmt.Errorf("failed to get commit status: %w", err))
	} else if !status.Successful {
		panic(fmt.Errorf("transaction %s failed to commit with status: %d", status.TransactionID, int32(status.Code)))
	}

	if err := callOffchainSubmit(data); err != nil {
		panic(err)
	}
}

// QueryDataWithIdentifier TODO implementation
func QueryDataWithIdentifier(identifier string) error {
	return nil
}

//TODO implementation
func callOffchainSubmit(data StoreData) error {
	return nil
}
