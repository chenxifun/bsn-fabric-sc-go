package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/chenxifun/bsn-fabric-sc-go/crosschaincode"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var successMsg = []byte("success")
var err_NoFunc = shim.Error("function not found")

type CrossData struct {
	Id     string `json:"id"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

func crossKey(id string) string {
	return "css_" + id
}

type SCChaincode struct {
}

func (c *SCChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Init")
	return shim.Success(successMsg)
}

func (c *SCChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Invoke")

	function, args := stub.GetFunctionAndParameters()
	if strings.ToLower(function) == "callnft" {
		return c.callNFT(stub, args)
	}

	if strings.ToLower(function) == "callback" {
		return c.callback(stub, args)
	}

	if strings.ToLower(function) == "query" {
		return c.query(stub, args)
	}

	return err_NoFunc
}

func (c *SCChaincode) callNFT(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 3 {
		return shim.Error("the args has error")
	}

	input := args[0]
	callBackChaincode := args[1]
	callBackFcn := args[2]

	reqId, err := crosschaincode.CallService(stub, "nft", input, callBackChaincode, callBackFcn, 100)
	if err != nil {
		return shim.Error("callNFT has failed ," + err.Error())
	}
	fmt.Println(reqId)

	cd := &CrossData{
		Id:    reqId,
		Input: input,
	}

	cdb, _ := json.Marshal(cd)
	if err := stub.PutState(crossKey(reqId), cdb); err != nil {
		return shim.Error(fmt.Sprintf("put data info error；%s", err))
	}

	return shim.Success([]byte(reqId))
}

func (c *SCChaincode) callback(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	output := args[0]
	res, err := crosschaincode.GetCallBackInfo(output)
	if err != nil {
		return shim.Error("get callback info has error")
	}

	ser, err := stub.GetState(crossKey(res.RequestId))
	if err != nil || len(ser) == 0 {
		return shim.Error("the requestID invalid")
	}

	cd := &CrossData{}
	err = json.Unmarshal(ser, cd)
	if err != nil {
		return shim.Error("error")
	}
	cd.Output = res.Output
	cdb, _ := json.Marshal(cd)
	if err := stub.PutState(crossKey(res.RequestId), cdb); err != nil {
		return shim.Error(fmt.Sprintf("put data info error；%s", err))
	}
	return shim.Success(successMsg)
}

func (c *SCChaincode) query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	id := args[0]
	ser, err := stub.GetState(crossKey(id))
	if err != nil || len(ser) == 0 {
		return shim.Error("the requestID invalid")
	}
	return shim.Success(ser)
}
