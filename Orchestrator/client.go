package Orchestrator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Client struct
type client struct {
	httpAdrs string
}


//Init Client
func NewClient(httpAdrs string) *client {
	c := client{httpAdrs: httpAdrs}
	return &c
}

//	SendJson()
//	Sends Json file to target http server
//	returns server response as json
func ExchangeJson(c client, struc interface{}) interface{}{

	jsonStr, err := json.Marshal(struc)
	errorHandler(err)

	resp, err := http.Post(c.httpAdrs, "application/json", bytes.NewBuffer(jsonStr))
	errorHandler(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorHandler(err)

	jsonData := []byte(body)
	var result interface{}

	unmarshallerr := json.Unmarshal(jsonData, result)
	errorHandler(unmarshallerr)
	return result
}

func errorHandler(err error){
	if err != nil {
		panic(err)
	}
}