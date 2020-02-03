package controller

import (
	"fmt"
	. "stocks-blockchain/application/model"
	"testing"
)

func TestInvoke(t *testing.T) {
	stock := NewStock("A1305","华夏","2019-1-1",0,10000,1.0)
	err := Invoke("IPO",stock)
	if err != nil{
		t.Fatal(err.Error())
	}
}

func TestQuery(t *testing.T) {
	err , data:= Query("A1305")
	if err != nil{
		t.Fatal(err.Error())
	}
	fmt.Println(data)
}
