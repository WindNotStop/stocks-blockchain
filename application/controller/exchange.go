package controller

import (
	"os/exec"
	. "stocks-blockchain/application/model"
	"strconv"
)

func Query(uid string)(error,string){
	cmd := exec.Command("docker","exec","cli","scripts/query.sh.template", uid)
	data, err := cmd.Output()
	if err != nil{
		return err, string(data)
	}
	return nil,string(data)
}

func Invoke(pattern string,stock *Stock)error{
	Type := strconv.FormatUint(uint64(stock.Type),10)
	amount := strconv.FormatUint(stock.Amount,10)
	price := strconv.FormatFloat(stock.Price,'f',11,64)
	cmd := exec.Command("docker","exec","cli","scripts/invoke.sh.template",pattern,stock.Uid,stock.Name,stock.Date,Type,amount,price)
	err := cmd.Run()
	if err != nil{
		return err
	}
	return nil
}
