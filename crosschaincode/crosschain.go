package crosschaincode

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

type ServiceRequest struct {
	RequestId   string        `json:"requestID,omitempty"`   //服务请求 ID  本合约中使用 合约交易ID
	ServiceName string        `json:"serviceName,omitempty"` //服务定义名称
	Input       string        `json:"input,omitempty"`       //服务请求输入；需符合服务的输入规范
	Timeout     uint64        `json:"timeout,omitempty"`     //请求超时时间；在目标链上等待的最大区块数
	CallBack    *CallBackInfo `json:"callback,omitempty"`    //回调的合约以及方法
}

type CallBackInfo struct {
	ChainCode string `json:"chainCode"`
	FuncName  string `json:"funcName"`
}
type ServiceResponse struct {
	RequestId   string `json:"requestID,omitempty"` //服务请求 ID  本合约中使用 合约交易ID
	ErrMsg      string `json:"errMsg,omitempty"`
	Output      string `json:"output,omitempty"`
	IcRequestId string `json:"icRequestID,omitempty"`
}
