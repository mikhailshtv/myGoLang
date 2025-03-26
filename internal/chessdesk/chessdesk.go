package chessdesk

import (
	"fmt"
	"log"
	"strings"
)

const space, hashtag string = " ", "#"

func PrintChessDesk() {
	var sideLength int

	fmt.Println("Введите желаемую длину стороны доски:")
	_, err := fmt.Scan(&sideLength)

	if err != nil {
		fmt.Println("Введено некорректное значение, введите целое число (int)")
		return
	}

	deskArr := getFilledDescArr(sideLength)
	chessDeskString := getChessDeskString(deskArr)

	fmt.Println(chessDeskString)
}

func getFilledDescArr(sideLength int) [][]string {
	deskArr := make([][]string, sideLength)
	for i := range deskArr {
		deskArr[i] = make([]string, sideLength)
	}
	for i, value := range deskArr {
		for j := range value {
			if (i+j)%2 == 0 {
				deskArr[i][j] = space
			} else {
				deskArr[i][j] = hashtag
			}
		}
	}
	return deskArr
}

func getChessDeskString(deskArr [][]string) string {
	var builder strings.Builder
	for i, value := range deskArr {
		_, err := builder.WriteString(strings.Join(value, ""))
		if err != nil {
			log.Fatal(err)
		}
		if i < len(deskArr) {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
