package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// 테스트용 JSON 데이타
	jsonBytes, _ := json.Marshal(Member{"Reo", 31, true})

	// JSON 디코딩
	var mem Member
	err := json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(mem.Name, mem.Age, mem.Active)
}
