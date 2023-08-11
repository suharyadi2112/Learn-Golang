package main

import (
	"fmt"//for print console
	"github.com/rdegges/go-ipify"//source
)

func main() {
	ip, err := ipify.GetIp()//get ip form ipify
	if err != nil {
		fmt.Println(err.Error())//output error
	}
	fmt.Println(ip)//output ip
	//36.80.XXX.XX
}
