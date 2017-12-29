package enterbj

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"strings"
)

var curtime="https://enterbj.zhongchebaolian.com/enterbj/platform/enterbj/curtime_03?userid=1aa12653655043019ff24e0fba2dd84b"
//var curtime = "http://www.baidu.com";

func TryCurTime() {
	resp, err := http.Get(curtime)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if !strings.Contains(resp.Status, "200") {
		fmt.Println(resp.Status)
		TryCurTime()
	} else {
		io.Copy(os.Stdout, resp.Body)
	}
}
