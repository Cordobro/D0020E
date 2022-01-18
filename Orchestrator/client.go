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
func NewClient(httpAdrs string) *client {
	c := client{httpAdrs: httpAdrs}
	return &c
}


func ExchangeJson(c *client, struc interface{}) interface{}{
	jsonStr, err := json.Marshal(struc)
	errorHandler(err)
    resp, err := http.Post(c.httpAdrs, "application/json", bytes.NewBuffer(jsonStr))
	errorHandler(err);
	defer resp.Body.Close()

	var jsonBody []byte
	jsonBody, err = ioutil.ReadAll(resp.Body)
	errorHandler(err)
	fmt.Println(string(jsonBody))
	
	var result interface{}
	err = json.Unmarshal(jsonBody, &result)
	errorHandler(err)
	return result
}



func errorHandler(err error){
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
