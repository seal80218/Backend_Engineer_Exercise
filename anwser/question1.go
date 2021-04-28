package anwser

import (
	"strconv"
	"strings"
)

func GetAmountFormat(str string) string {
	strSplit := strings.Split(str, ".")

	if _, err := strconv.ParseFloat(str, 64); err != nil {
		return str
	}

	if len(strSplit) > 2 {
		return str
	}

	firstStrLen := len(strSplit[0])
	if firstStrLen <= 3 {
		return str
	}

	buf := make([]byte, 0)
	r := firstStrLen % 3
	data := []byte(strSplit[0])
	if r > 0 {
		buf = []byte(strSplit[0])[0:r]
		data = data[r:]
		buf = append(buf, ',')
	}

	for i := 0; i < len(data); i++ {
		if i > 0 && i%3 == 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, data[i])
	}

	if len(strSplit) == 2 {
		buf = append(buf, '.')
		buf = append(buf, strSplit[1]...)
	}

	return string(buf)
}
