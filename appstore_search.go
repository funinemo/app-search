package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func assert(data interface{}, tabnum int) {
	var tab string
	for i := 0; i < tabnum; i++ {
		tab = tab + "\t"
	}
	//	fmt.Printf("%s[%d]\r\n", tab, tabnum)
	tabnum++
	switch data.(type) {
	case string:
		//fmt.Printf("%s%v\r\n", tab, data.(string))
	case float64:
		//fmt.Printf("%s%v\r\n", tab, data.(float64))
	case bool:
		//fmt.Printf("%s%v\r\n", tab, data.(bool))
	case nil:
		//		fmt.Printf("null\r\n")
	case []interface{}:
		//		fmt.Printf("=>\r\n")
		for _, v := range data.([]interface{}) {
			//			fmt.Printf("%s{\r\n", tab)
			fmt.Printf("%s%v\r\n", tab, v)
			assert(v, tabnum)
			//			fmt.Printf("%s}\r\n", tab)
		}
		//		fmt.Printf("<=\r\n")
	case map[string]interface{}:
		//		fmt.Printf("==>\r\n")
		for k, v := range data.(map[string]interface{}) {
			//			fmt.Printf("%s{\r\n", tab)
			//			fmt.Printf("%s[%v]:%v\r\n", tab, k, v)
			fmt.Printf("%s[%v]:%v\r\n", tab, k, v)
			assert(v, tabnum)
			//			fmt.Printf("%s}\r\n", tab)
		}
		//		fmt.Printf("<===\r\n")
	default:
	}
	tabnum--
	//	fmt.Printf("=====END=============\r\n")
}
func main() {
	url := "https://itunes.apple.com/search?term=puzzle&country=JP&entity=software&limit=100"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data interface{}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.Decode(&data)
	assert(data, 0)
	fmt.Println()
	//	fmt.Printf("%v", string(body))

}
