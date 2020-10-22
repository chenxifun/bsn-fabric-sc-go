package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CrossChainInfo struct {
	ChainCodeName      string `json:"cn"`
	CrossChainCodeName string `json:"ccn"`
}

func getCrossChaincodeKey() string {
	return "cccn_key"
}

func NewCrossChainInfo(chainCodeName, crossChainCodeName string) []byte {

	info := &CrossChainInfo{
		ChainCodeName:      chainCodeName,
		CrossChainCodeName: crossChainCodeName,
	}

	b, _ := json.Marshal(info)
	return b
}

func InitCrossChain(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	if len(args) < 2 {
		return shim.Error("Init failed,args must be 2")
	}
	chaincode := args[0]
	crossChaincode := args[1]
	if err := stub.PutState(getCrossChaincodeKey(), NewCrossChainInfo(chaincode, crossChaincode)); err != nil {
		return shim.Error(fmt.Sprintf("put service info error；%s", err))
	}

	return shim.Success([]byte("ok"))

}

func CallService(stub shim.ChaincodeStubInterface, serviceName, input, callback string, timeout uint64) (string, error) {

	crossInfo, err := GetCrossInfo(stub)
	if err != nil {
		return "", err
	}
	req := &ServiceRequest{
		ServiceName: serviceName,
		Input:       input,
		Timeout:     timeout,
		CallBack: &CallBackInfo{
			ChainCode: crossInfo.ChainCodeName,
			FuncName:  callback,
		},
	}

	b, _ := json.Marshal(req)

	var args [][]byte
	args = append(args, []byte("callservice"))
	args = append(args, b)

	res := stub.InvokeChaincode(crossInfo.CrossChainCodeName, args, "")
	txId := string(res.Payload)
	//stub.SetEvent(fmt.Sprintf("CallService_%s", txId), []byte(txId))
	return txId, nil
}

func GetCrossInfo(stub shim.ChaincodeStubInterface) (*CrossChainInfo, error) {

	cr, err := stub.GetState(getCrossChaincodeKey())
	if err != nil || len(cr) == 0 {
		return nil, errors.New("the crosschain info invalid")
	}

	crinfo := &CrossChainInfo{}
	err = json.Unmarshal(cr, crinfo)
	if err != nil {
		return nil, errors.New("the crosschain info invalid")
	}
	return crinfo, nil
}
