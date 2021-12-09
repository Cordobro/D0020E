package Orchestrator

import(
	"net/http"
	"bytes"
	"io/ioutil"
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

//	SendJson()
//	Sends Json file to target http server
//	returns server response as json
func SendJson(c client, jsonStr []byte) []byte{

	resp, err := http.Post(c.httpAdrs, "application/json", bytes.NewBuffer(jsonStr))
	
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//Failed to read response.
		panic(err)
	}

	result := []byte(body)

	return result
}

