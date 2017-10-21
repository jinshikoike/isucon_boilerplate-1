package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var prefix = "."

func main() {

	var fp *os.File

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	re := regexp.MustCompile(`(.*?): (\d+\.\d+|\d+)(.*):(.*)$`)
	data := make(map[string][]float64)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		matchList := re.FindAllStringSubmatch(scanner.Text(), -1)[0]
		//fmt.Println(matchList)
		var value float64
		switch matchList[3] {
		case "ns":
			value, _ = strconv.ParseFloat(matchList[2], 64)
		case "µs":
			value, _ = strconv.ParseFloat(matchList[2], 64)
			value = value * math.Pow10(3)
		case "ms":
			value, _ = strconv.ParseFloat(matchList[2], 64)
			value = value * math.Pow10(6)
		case "s":
			value, _ = strconv.ParseFloat(matchList[2], 64)
			value = value * math.Pow10(9)
		}
		data[matchList[1]] = append(data[matchList[1]], value)
	}

	dataStatics := make(map[string]map[string]float64)
	for key, value := range data {
		var sum float64
		sum = 0.0
		max := value[0]
		min := value[0]

		for _, v := range value {
			sum = sum + v
			if max < v {
				max = v
			}
			if min > v {
				min = v
			}
		}
		dataStatics[key] = map[string]float64{"avg": sum / float64(len(value))}
		dataStatics[key]["max"] = max
		dataStatics[key]["min"] = min
	}

	for key, value := range dataStatics {
		val, unitStr := getUnitAndValue(value["avg"])
		fmt.Println("平均: "+key+" : ", val, unitStr)
		val, unitStr = getUnitAndValue(value["max"])
		fmt.Println("MAX: "+key+" : ", val, unitStr)
		val, unitStr = getUnitAndValue(value["min"])
		fmt.Println("MIN: "+key+" : ", val, unitStr)
		fmt.Println("---------------------------------------------")
	}

}

func getUnitAndValue(val float64) (float64, string) {
	i := 0
	tmpSum := val
	for {
		if tmpSum < 1 {
			break
		}
		tmpSum = tmpSum / 10.0
		i++
	}

	unitStr := ""
	if i < 4 {
		unitStr = "ns"
	} else if i < 7 {
		val = val / math.Pow10(3)
		unitStr = "µs"
	} else if i < 10 {
		val = val / math.Pow10(6)
		unitStr = "ms"
	} else {
		val = val / math.Pow10(9)
		unitStr = "s"
	}
	return val, unitStr
}
