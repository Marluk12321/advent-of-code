package main

import (
	"2022/13/part_1/checker"
	"2022/13/part_1/data"
	"bufio"
	"fmt"
	"os"
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
	var left interface{}
	var right interface{}
	var rightOrderSum int

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	oddLine := true
	var pairCounter int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var dataString string
		if oddLine {
			left = data.Parse(line)
			dataString = dataToString(&left)
		} else {
			right = data.Parse(line)
			dataString = dataToString(&right)
			pairCounter++
			if checker.InRightOrder(&left, &right) {
				rightOrderSum += pairCounter
			}
		}
		oddLine = !oddLine
		if dataString != line {
			fmt.Println(line)
			fmt.Println("->", dataString)
			fmt.Println("WRONG!")
		}
	}
	file.Close()

	fmt.Println("Sum:", rightOrderSum)
}
