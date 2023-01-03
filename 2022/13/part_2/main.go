package main

import (
	"2022/13/part_2/comparison"
	"2022/13/part_2/data"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func dataToString(data *interface{}) string {
	switch (*data).(type) {
	case []interface{}:
		dataElems := (*data).([]interface{})
		strElems := make([]string, len(dataElems))
		for i := range dataElems {
			strElems[i] = dataToString(&dataElems[i])
		}
		return fmt.Sprintf("[%v]", strings.Join(strElems, ","))
	case int:
		return fmt.Sprint((*data).(int))
	default:
		panic((*data))
	}
}

func main() {
	dividerPacket1 := data.Parse("[[2]]")
	dividerPacket2 := data.Parse("[[6]]")
	packets := []interface{}{dividerPacket1, dividerPacket2}

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		packet := data.Parse(line)
		packets = append(packets, packet)
	}
	file.Close()

	sort.Slice(packets, func(i, j int) bool {
		packet1 := &packets[i]
		packet2 := &packets[j]
		return comparison.Compare(packet1, packet2) < 0
	})
	dividerPosProduct := 1
	for i := range packets {
		packet := &packets[i]
		fmt.Println(dataToString(&packets[i]))
		if comparison.Compare(packet, &dividerPacket1) == 0 || comparison.Compare(packet, &dividerPacket2) == 0 {
			dividerPosProduct *= i + 1
		}
	}
	fmt.Println("Product:", dividerPosProduct)
}
