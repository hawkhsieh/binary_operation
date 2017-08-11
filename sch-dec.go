package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

func main() {

	schStr := "3+8P"
	schBin, err := base64.StdEncoding.DecodeString(schStr)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("%+v\n", schBin)

	var schUnt32 uint32
	sch4Bin := make([]byte, 4)
	for i, b := range schBin {
		sch4Bin[i] = b
	}
	buf := bytes.NewReader(sch4Bin)
	err = binary.Read(buf, binary.LittleEndian, &schUnt32)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Printf("%x\n", schUnt32)

	turnOn := schUnt32 & (1 << 0)        //1bit
	slot := schUnt32 >> 1 & (1<<2 - 1)   //2bit
	minute := schUnt32 >> 3 & (1<<6 - 1) //6bit
	hour := schUnt32 >> 9 & (1<<5 - 1)   //5bit

	fmt.Printf("turnOn:%d,slot:%d,minute:%d,hour:%d\n", turnOn, slot, minute, hour)

	thisIsWeekSchedule := false
	if (schUnt32 >> 14 & (1 << 0)) == 0 {
		thisIsWeekSchedule = true
	}

	if thisIsWeekSchedule {
		weekBin := schUnt32 >> 15 & (1<<7 - 1)
		week := make([]uint, 0)
		for i := uint(0); i < 7; i++ {
			if (weekBin & (1 << (6 - i))) != 0 {
				week = append(week, i)
			}
		}
		fmt.Printf("week:%+v", week)

	} else {
		day := schUnt32 >> 15 & (1<<6 - 1)
		fmt.Printf("day:%d", day)

	}

}

