package main

import (
	"fmt"
	"strconv"

	"net"
	"net/rpc"
	"bytes"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spiral/goridge"
)

type Excel struct{}

func (s *Excel) Encode(data [][]interface{}, res *[]byte) error {
	xlsx := excelize.NewFile()

	for r, row := range data {
		for c, cell := range row {
			err := xlsx.SetCellValue("Sheet1", getColumnName(c)+strconv.Itoa(r+1), cell)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}

	buf, err := xlsx.WriteToBuffer()
	
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	*res = append(*res, buf.Bytes()...)

	return nil
}

func (s *Excel) Decode(file []byte, res *[][]string) error {
	buf := bytes.NewBuffer(file)
	xlsx, err := excelize.OpenReader(buf)
    if err != nil {
        fmt.Println(err)
        return err
    }

    // 获取 Sheet1 上所有单元格
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
        fmt.Println(err)
        return err
	}
	
	*res = append(*res, rows...)

	return nil
}

func getColumnName(c int) (name string) {
	for c >= 26 {
		temp := c%26
		name = string(temp + 65) + name
		c = c/26-1;
	}
	name = string(65 + c) + name
	
	return
}

func main() {

	ln, err := net.Listen("tcp", ":6001")
	if err != nil {
		panic(err)
	}

	rpc.Register(new(Excel))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeCodec(goridge.NewCodec(conn))
	}
}