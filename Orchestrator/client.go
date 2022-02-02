package Orchestrator

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	//forms "arrowhead/Orchestrator/forms"
)

//Client struct
type client struct {
	httpAdrs string
}

//Init Client
func newClient(httpAdrs string) *client {
	c := client{httpAdrs: httpAdrs}
	return &c
}

func exchangeJson(c *client, struc interface{}) []byte {
	jsonStr, err := json.Marshal(struc)
	errorHandler(err)
	resp, err := http.Post(c.httpAdrs, "application/json", bytes.NewBuffer(jsonStr))
	errorHandler(err)
	defer resp.Body.Close()

	jsonBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(jsonBody))

	errorHandler(err)

	return jsonBody
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

//OLD
/*func ExchangeJson(c client, struc interface{}) interface{}{

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
}*/
