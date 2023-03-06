package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	// content, err := ioutil.ReadFile("./note.txt")
	file, err := os.Open("./note.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Out put to excel.
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Write content.
	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, " ")
		if len(l) == 2 {
			name := l[0]
			id := l[1]
			f.SetCellValue("sheet1", "A"+strconv.Itoa(index+1), name)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(index+1), id)
		}
		index++
	}
	// strs := strings.Split(string(content), "\n")
	// for index, str := range strs {
	// 	l := strings.Split(str, " ")
	// 	if len(l) == 2 {
	// 		name := l[0]
	// 		id := l[1]
	// 		f.SetCellValue("sheet1", "A"+strconv.Itoa(index+1), name)
	// 		f.SetCellValue("sheet1", "B"+strconv.Itoa(index+1), id)
	// 	}
	// }

	if err := f.SaveAs("test.xlsx"); err != nil {
		fmt.Println(err)
	}
}
