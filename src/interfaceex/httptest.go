package interfaceex

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

func Run() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	err = r.Body.Close()
	if err != nil {
		fmt.Print(err)
	}
}
