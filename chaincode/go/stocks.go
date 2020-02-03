package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

const (
	A = iota
	HK
	US
)

type Stock struct {
	Uid    string `json:"stock_uid"`
	Name   string `json:"stock_name"`
	Date   string `json:"ipo_date"`
	Type   uint8  `json:"stock_type"`
	Amount uint64  `json:"stock_amount"`
	Price  float64 `json:"stock_price"`
}

func NewStock(uid, name, date string, Type uint8,amount uint64,price float64)*Stock{
	return &Stock{
		Uid:    uid,
		Name:   name,
		Date:   date,
		Type:   Type,
		Amount: amount,
		Price:  price,
	}
}

type stocksCC struct{}

//IPO（首次公开募股）
func (c *stocksCC) IPO(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//检查参数的个数
	if len(args) != 6 {
		return shim.Error("Missing required args(uid, name, date, type, amount, price)")
	}

	//检查参数类型
	uid := args[0]
	name := args[1]
	date := args[2]
	Type, err := strconv.ParseUint(args[3], 10, 8)
	if err != nil{
		return shim.Error("Type err")
	}
	amount, err := strconv.ParseUint(args[4],10,64)
	if err != nil{
		return shim.Error("amount err")
	}
	price, err := strconv.ParseFloat(args[5],64)
	if err != nil{
		return shim.Error("price err")
	}

	//查看股票信息是否存在
	stockTemp , err := stub.GetState(uid)
	if err != nil {
		return shim.Error("query err:"+err.Error())
	}
	if stockTemp != nil{
		return shim.Error("stock exists")
	}

	//构造股票实例
	stock := NewStock(uid,name,date,uint8(Type),amount,price)
	stockBytes, err := json.Marshal(stock)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal stock error: %s", err))
	}

	//增加股票
	err = stub.PutState(uid, stockBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("save stock error: %s", err))
	}

	return shim.Success(nil)
}

//Update（修改股票信息）
func (c *stocksCC) Update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//检查参数的个数
	if len(args) != 6 {
		return shim.Error("Missing required args(uid, name, date, type, amount, price)")
	}

	//检查参数类型
	uid := args[0]
	name := args[1]
	date := args[2]
	Type, err := strconv.ParseUint(args[3], 10, 8)
	if err != nil{
		return shim.Error("Type err")
	}
	amount, err := strconv.ParseUint(args[4],10,64)
	if err != nil{
		return shim.Error("amount err")
	}
	price, err := strconv.ParseFloat(args[5],64)
	if err != nil{
		return shim.Error("price err")
	}

	//查看股票信息是否存在
	stockTemp , err := stub.GetState(uid)
	if err != nil {
		return shim.Error("query err"+err.Error())
	}
	if stockTemp == nil{
		return shim.Error("stock not exists")
	}

	//构造股票实例
	stock := NewStock(uid,name,date,uint8(Type),amount,price)
	stockBytes, err := json.Marshal(stock)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal stock error: %s", err))
	}

	//更新股票
	err = stub.PutState(uid, stockBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("save stock error: %s", err))
	}

	return shim.Success(nil)
}


//查询股票信息
func (c *stocksCC) Query(stub shim.ChaincodeStubInterface, uid string) pb.Response {
	stockBytes, err := stub.GetState(uid)
	if err != nil {
		return shim.Error("query err:"+err.Error())
	}
	if stockBytes == nil{
		return shim.Error("stock not exists")
	}
	return shim.Success(stockBytes)
}

//删除股票
func (c *stocksCC) Delete(stub shim.ChaincodeStubInterface, uid string) pb.Response {
	err := stub.DelState(uid)
	if err != nil {
		return shim.Error("delete err:"+err.Error())
	}
	return shim.Success(nil)
}

//
func (c *stocksCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}


func (c *stocksCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "IPO":
		return c.IPO(stub, args)
	case "Update":
		return c.Update(stub, args)
	case "Query":
		return c.Query(stub, args[0])
	case "Delete":
		return c.Delete(stub,args[0])
	default:
		return shim.Error(fmt.Sprintf("unsupported function: %s", funcName))
	}

}

func main() {
	err := shim.Start(new(stocksCC))
	if err != nil {
		fmt.Printf("Error starting AssertsExchange chaincode: %s", err)
	}
}