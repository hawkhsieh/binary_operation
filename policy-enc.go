package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	threshold := 30
	condition := "<="
	slots := []int{1,1,0,0}
	action := 1              //1: on , 2:off
	push := 1                //1: push , 0: no push
	thvtlv := "thv"

	policy := 0
	if thvtlv == "thv" {
		policy = 0x2 //2bit
	}else if thvtlv == "tlv" {
		policy = 0x1 //2bit
	}

	policy |= (push) << 2     //1bit
	policy |= (action) << 3   //1bit

	slotBit := 0
	for slot,en := range slots {
		if en == 1 {
			slotBit |= (1 << uint(slot))
		}
	}
	policy |= slotBit << 4   //4bit

	conditionBit:=0
	switch(condition){
	case ">":
		conditionBit=0
	case "=":
		conditionBit=1
	case "<":
		conditionBit=2
	case "<=":
		conditionBit=3
	case ">=":
		conditionBit=4

	}
	policy |= conditionBit << 8   //3bit
	policy |= threshold << 11     //8bit

	fmt.Printf("output:%x\n", policy)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int32(policy))
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("binary:%s\n", strconv.FormatInt(int64(policy),2))
	fmt.Printf("little endian in int:%s\n", hex.EncodeToString(buf.Bytes()[0:3]))

	fmt.Printf("base64:%s\n", base64.StdEncoding.EncodeToString(buf.Bytes()[0:3]))
}

