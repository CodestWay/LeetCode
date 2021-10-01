package SwordToOffer

import (
	"strconv"
)

var res_46 = make(map[string]int)
var temp = make(map[string]string)

func translateNum(num int) int {
	temp["0"] = "a"
	temp["1"] = "b"
	temp["2"] = "c"
	temp["3"] = "d"
	temp["4"] = "e"
	temp["5"] = "f"
	temp["6"] = "g"
	temp["7"] = "h"
	temp["8"] = "i"
	temp["9"] = "j"
	temp["10"] = "k"
	temp["11"] = "l"
	temp["12"] = "m"
	temp["13"] = "n"
	temp["14"] = "o"
	temp["15"] = "p"
	temp["16"] = "q"
	temp["17"] = "r"
	temp["18"] = "s"
	temp["19"] = "t"
	temp["20"] = "u"
	temp["21"] = "v"
	temp["22"] = "w"
	temp["23"] = "x"
	temp["24"] = "y"
	temp["25"] = "z"
	bytes := []byte(strconv.Itoa(num))
	visit46(bytes, 0, 0, "")
	visit46(bytes, 0, 1, "")
	return len(res_46)
}

func visit46(arr []byte, start, end int, curr string) {
	if start == end {
		if start < len(arr) {
			val := temp[string(arr[start])]
			curr += val
			visit46(arr, end+1, end+1, curr)
			visit46(arr, end+1, end+2, curr)
		} else if start == len(arr) {
			res_46[curr] = 1
			return
		} else {
			return
		}
	} else {
		if end > len(arr) {
			return
		}
		if end == len(arr) {
			val, flag := temp[string(arr[start:end+1])]
			if !flag {
				return
			}
			curr += val
			res_46[curr] = 1
			return
		}
		if end < len(arr) {
			val, flag := temp[string(arr[start:end+1])]
			if !flag {
				return
			}
			curr += val
			visit46(arr, end+1, end+1, curr)
			visit46(arr, end+1, end+2, curr)
		}
	}
}
