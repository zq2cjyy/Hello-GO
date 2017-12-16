package jsonopt

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type (
	//定义结构的时候，字段首字母一定大写，否则无法序列化成功！！！！！！
	cityResult struct {
		City     string `json:"city"`
		CityName string `json:"cityname"`
	}

	returnResult struct {
		Status string       `json:"status"`
		Msg    string       `json:"msg"`
		Result []cityResult `json:"result"`
	}
)

func RunDecode() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://apis.baidu.com/netpopo/vehiclelimit/city", nil)

	if err != nil {
		log.Println("err1", err)
		return
	}

	request.Header.Add("apikey", "2728fdddd5a291237cf4a56c7b6b9d12")

	resp, err := client.Do(request)

	if err != nil {
		log.Println("err2", err)
		return
	}
	defer resp.Body.Close()

	//log.Println(resp.Body)
	//io.Copy(os.Stdout, resp.Body)

	var rst *returnResult
	//err = json.NewDecoder(resp.Body).Decode(&rst)
	var data = `
    {"status": "0",
    "msg": "ok",
    "result": [
        {
            "city": "beijing",
            "cityname": "北京"
        },
        {
            "city": "tianjin",
            "cityname": "天津"
        },
        {
            "city": "hangzhou",
            "cityname": "杭州"
        },
        {
            "city": "chengdu",
            "cityname": "成都"
        },
        {
            "city": "lanzhou",
            "cityname": "兰州"
        },
        {
            "city": "guiyang",
            "cityname": "贵阳"
        },
        {
            "city": "nanchang",
            "cityname": "南昌"
        },
        {
            "city": "changchun",
            "cityname": "长春"
        },
        {
            "city": "haerbin",
            "cityname": "哈尔滨"
        },
        {
            "city": "wuhan",
            "cityname": "武汉"
        },
        {
            "city": "shanghai",
            "cityname": "上海"
        },
        {
            "city": "shenzhen",
            "cityname": "深圳"
        }
    ]
}`
	//data, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data))
	err = json.Unmarshal([]byte(data), &rst)
	if err != nil {
		log.Println("err3", err)
		return
	}

	if rst.Status == "0" {
		for _, r := range rst.Result {
			fmt.Println("city->" + r.City + ".....cityname->" + r.CityName)
		}
		fmt.Println("fuck")
	} else {
		fmt.Println(rst.Status)
	}
}
