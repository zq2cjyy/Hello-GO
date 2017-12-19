package jsonopt

import (
	"encoding/json"
	"fmt"
)

func RunEncode() {
	var rt = returnResult{Status: "1", Msg: "success", Result: make([]cityResult, 0, 3)}
	//rt.Result[0] = cityResult{City: "1", CityName: "beijing"}
	//rt.Result[1] = cityResult{City: "2", CityName: "taiyuan"}
	rt.Result = append(rt.Result, cityResult{City: "1", CityName: "beijing"})
	rt.Result = append(rt.Result, cityResult{City: "2", CityName: "taiyuan"})

	data, err := json.Marshal(rt)
	if err != nil {
		fmt.Println("err3")
	}

	fmt.Println(string(data))
}
