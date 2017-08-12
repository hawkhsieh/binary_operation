package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {

	schStr := "PvMA"
	schBin, err := base64.StdEncoding.DecodeString(schStr)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("%+v\n", schBin)

	var policyUint32 uint32
	policyBytes := make([]byte, 4)
	for i, b := range schBin {
		policyBytes[i] = b
	}
	buf := bytes.NewReader(policyBytes)
	err = binary.Read(buf, binary.LittleEndian, &policyUint32)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Printf("%x\n", policyUint32)

type policy struct {
	thvtlv string
	push uint32
	action uint32
	slots []uint
	cond string
	threshold int
}
	var p policy
	thvtlvBin := policyUint32 & (1<<2 -1)       //2bit
	switch(thvtlvBin){
	case 2:
		p.thvtlv="thv"
	case 1:
		p.thvtlv="tlv"
	default:
		log.Panicf("0x%x is invalid",thvtlvBin)

	}

	p.push = policyUint32 >> 2 & (0x1)          //1bit
	p.action = policyUint32 >> 3 & (0x1)        //1bit
	slotsBin := policyUint32 >> 4 & (1<<4 - 1)  //4bit
	p.slots = make([]uint, 4)
	for i := 0; i < len(p.slots); i++ {
		if (slotsBin & (1 << uint(i))) != 0 {
			p.slots[i] = 1
		}
	}

	condBin := policyUint32 >> 8 & (1<<3 - 1)   //3bit
	switch(condBin){
	case 0x00:
		p.cond =">"
	case 0x01:
		p.cond ="="
	case 0x02:
		p.cond ="<"
	case 0x03:
		p.cond ="<="
	case 0x04:
		p.cond =">="
	default:
		log.Panicf("0x%x is invalid",condBin)
	}

	p.threshold = int(policyUint32) >> 11 & 0xff   //8bit

	fmt.Printf("%+v\n", p)


}

