package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	turnOn := true
	slot := 3
	minute := 59
	hour := 23
	weeks := []uint{0,2,4,5,6}
	day := 3

	thisIsWeekSchedule := true

	sch := 0
	if turnOn {
		sch |= 1 << 0 //1bit
	}

	sch |= (slot & 0x03) << 1 //2bit
	sch |= (minute) << 3      //6bit
	sch |= (hour % 24) << 9   //5bit

	if thisIsWeekSchedule {

		fmt.Printf("%02d:%02d slot%d=%v every %d week\n", hour, minute, slot, turnOn, weeks)
		var weekBit uint
		
		for _,week := range weeks {
		    weekBit |= (1<<(6-week))
		}
		sch |= int(weekBit) << 15 //3bit

	} else {
		fmt.Printf("%02d:%02d slot%d=%v every %d month\n", hour, minute, slot, turnOn, day)
		sch |= 1 << 14   //1bit
		sch |= day << 15 //5bit
	}
	fmt.Printf("output:%x\n", sch)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int32(sch))
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("little endian in int:%s\n", hex.EncodeToString(buf.Bytes()[0:3]))

	fmt.Printf("base64:%s\n", base64.StdEncoding.EncodeToString(buf.Bytes()[0:3]))
}

