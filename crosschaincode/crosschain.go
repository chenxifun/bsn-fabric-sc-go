package crosschaincode

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
)

const (
	CrossChaincode = "cc_cross"
)

func CallService(stub shim.ChaincodeStubInterface, serviceName, input, callbackCC, callbackFcn string, timeout uint64) (string, error) {

	req := &ServiceRequest{
		ServiceName: serviceName,
		Input:       input,
		Timeout:     timeout,
	}

	if strings.TrimSpace(callbackCC) != "" && strings.TrimSpace(callbackFcn) != "" {
		req.callBack = &CallBackInfo{
			ChainCode: callbackCC,
			FuncName:  callbackFcn,
		}
	}

	b, _ := json.Marshal(req)

	var args [][]byte
	args = append(args, []byte("callservice"))
	args = append(args, b)

	res := stub.InvokeChaincode(CrossChaincode, args, "")
	txId := string(res.Payload)
	//stub.SetEvent(fmt.Sprintf("CallService_%s", txId), []byte(txId))
	return txId, nil
}

func GetCallBackInfo(output string) (*ServiceResponse, error) {
	res := &ServiceResponse{}
	err := json.Unmarshal([]byte(output), res)
	if err != nil {
		return nil, errors.New("error")
	}
	return res, nil
}

type ServiceRequest struct {
	RequestId   string        `json:"requestID,omitempty"`   //服务请求 ID  本合约中使用 合约交易ID
	ServiceName string        `json:"serviceName,omitempty"` //服务定义名称
	Input       string        `json:"input,omitempty"`       //服务请求输入；需符合服务的输入规范
	Timeout     uint64        `json:"timeout,omitempty"`     //请求超时时间；在目标链上等待的最大区块数
	callBack    *CallBackInfo `json:"callback,omitempty"`    //回调的合约以及方法
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
