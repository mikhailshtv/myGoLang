package chessdesk

import (
	"fmt"
	"strings"
)

func PrintChessDesk() {
	var space, hashtag string = " ", "#"
	var sideLength int

	fmt.Println("Введите желаемую длину стороны доски:")
	fmt.Scan(&sideLength)

	deskArr := make([][]string, sideLength)
	for i := range deskArr {
		deskArr[i] = make([]string, sideLength)
	}

	deskArr = getFilledDescArr(deskArr, space, hashtag)
	chessDeskString := getChessDeskString(deskArr)

	fmt.Println(chessDeskString)
}

func getFilledDescArr(deskArr [][]string, space string, hashtag string) [][]string {
	for i, value := range deskArr {
		for j := range value {
			if j%2 == 0 {
				if i > 0 && deskArr[i-1][j] == space {
					deskArr[i][j] = hashtag
				} else {
					deskArr[i][j] = space
				}
			} else {
				if i > 0 && deskArr[i-1][j] == hashtag {
					deskArr[i][j] = space
				} else {
					deskArr[i][j] = hashtag
				}
			}
		}
	}
	return deskArr
}

func getChessDeskString(deskArr [][]string) string {
	chessDeskString := ""
	for i, value := range deskArr {
		chessDeskString = chessDeskString + strings.Join(value, "")
		if i < len(deskArr) {
			chessDeskString += "\n"
		}
	}

	return chessDeskString
}
