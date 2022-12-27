package csv_deserialization

import (
	"errors"
	"strconv"
	"strings"
)

const separation = ","
const mapSeparation = ":"

var P = &p{}

type p struct{}

func (p p) Int(val string) int {
	if val == "" {
		return 0
	}
	val = strings.TrimSpace(val)

	a, err := strconv.Atoi(val)
	if err != nil {
		e := errors.New("failed to parse string to int: " + val + ".  " + err.Error())
		panic(e)
	}
	return a
}

func (p p) Float(val string) float64 {
	if val == "" {
		return 0
	}
	val = strings.TrimSpace(val)
	a, err := strconv.ParseFloat(val, 10)
	if err != nil {
		e := errors.New("failed to parse string to float: " + val + ".  " + err.Error())
		panic(e)
	}
	return a
}

func (p p) Bool(val string) bool {
	if val == "" {
		return false
	}
	val = strings.TrimSpace(val)
	val = strings.ToLower(val)
	if val == "0" || val == "false" {
		return false
	}
	if val == "1" || val == "true" {
		return true
	}
	e := errors.New("failed to parse string to bool: " + val + ".  ")
	panic(e)
}

func (p p) String(val string) string {
	return strings.TrimSpace(val)
}

func (p p) ArrInt(val string) []int {
	list := strings.Split(val, separation)
	result := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = p.Int(list[i])
	}
	return result
}

func (p p) ArrFloat(val string) []float64 {
	list := strings.Split(val, separation)
	result := make([]float64, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = p.Float(list[i])
	}
	return result
}

func (p p) ArrBool(val string) []bool {
	list := strings.Split(val, separation)
	result := make([]bool, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = p.Bool(list[i])
	}
	return result
}

func (p p) ArrString(val string) []string {
	list := strings.Split(val, separation)
	return list
}

func (p p) MapInt2Int(val string) map[int]int {
	result := make(map[int]int)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}

		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Int(subList[0])
		v := p.Int(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapInt2Float(val string) map[int]float64 {
	result := make(map[int]float64)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}

		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Int(subList[0])
		v := p.Float(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapInt2Bool(val string) map[int]bool {
	result := make(map[int]bool)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Int(subList[0])
		v := p.Bool(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapInt2String(val string) map[int]string {
	result := make(map[int]string)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Int(subList[0])
		v := p.String(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapFloat2Int(val string) map[float64]int {
	result := make(map[float64]int)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Float(subList[0])
		v := p.Int(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapFloat2Float(val string) map[float64]float64 {
	result := make(map[float64]float64)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Float(subList[0])
		v := p.Float(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapFloat2Bool(val string) map[float64]bool {
	result := make(map[float64]bool)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Float(subList[0])
		v := p.Bool(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapFloat2String(val string) map[float64]string {
	result := make(map[float64]string)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.Float(subList[0])
		v := p.String(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapString2Int(val string) map[string]int {
	result := make(map[string]int)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.String(subList[0])
		v := p.Int(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapString2Float(val string) map[string]float64 {
	result := make(map[string]float64)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.String(subList[0])
		v := p.Float(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapString2Bool(val string) map[string]bool {
	result := make(map[string]bool)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.String(subList[0])
		v := p.Bool(subList[1])
		result[k] = v
	}
	return result
}

func (p p) MapString2String(val string) map[string]string {
	result := make(map[string]string)
	list := strings.Split(val, separation)
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		subList := strings.Split(list[i], mapSeparation)
		if len(subList) != 2 {
			e := errors.New("item is not map kv: " + list[i])
			panic(e)
		}
		k := p.String(subList[0])
		v := p.String(subList[1])
		result[k] = v
	}
	return result
}
