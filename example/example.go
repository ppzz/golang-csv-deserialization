package main

import (
	"fmt"
	csvdeserialization "github.com/ppzz/csv-deserialization"
	"os"
)

type Player struct {
	Id       int               `csv:"id"`
	Power    float64           `csv:"power"`
	IsNewbie bool              `csv:"is_newbie"`
	Desc     string            `csv:"desc"`
	Skill    []int             `csv:"skill"`
	Score    map[int]int       `csv:"score"`
	Subject  map[string]string `csv:"subject"`
}

func main() {
	f, _ := os.Open("./example/example.csv")
	defer f.Close()

	c := csvdeserialization.Csv{}
	c.Read(f)

	var list []Player
	c.Attach(&list)

	for i := 0; i < len(list); i++ {
		item := list[i]
		fmt.Println(i, item.Id, item.Power, item.IsNewbie, item.Desc, item.Skill, item.Score, item.Subject)
	}
}
